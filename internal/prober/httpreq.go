package prober

import (
	"context"
	"net"
	"net/http"
	"net/http/httptrace"
	"net/url"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/klog/v2"
)

var (
	httpDurations = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "health_http_duration_seconds",
			Help:    "The response time of http requests",
			Buckets: []float64{0.001, 0.005, 0.01, 0.02, 0.03, 0.05, 0.075, 0.1, 0.2, 0.5, 0.75, 1, 1.5, 2, 2.5, 3, 4, 5},
		},
		[]string{"name", "status_code", "result", "url"},
	)
	dnsLookupDurations = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "health_http_dns_lookup_time_seconds",
			Help:    "The response time of dns lookup",
			Buckets: []float64{0.0005, 0.001, 0.002, 0.003, 0.004, 0.005, 0.006, 0.008, 0.01, 0.015, 0.02, 0.03, 0.05, 0.075, 0.1, 0.2, 0.5, 1},
		},
		[]string{"name", "status_code", "dns_error", "result", "url"},
	)
)

// TODO: it could be a bad practice!
func init() {
	prometheus.MustRegister(httpDurations)
	prometheus.MustRegister(dnsLookupDurations)

}

func (p *Prober) sendRequest(ctx context.Context, url string) error {
	var startTime time.Time
	var dnsStartTime time.Time
	var dnsLookupTime float64
	var dnsError string
	var dnsHost string
	httpTrace := &httptrace.ClientTrace{
		DNSStart: func(info httptrace.DNSStartInfo) {
			dnsHost = info.Host
			dnsStartTime = time.Now()
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			dnsLookupTime = time.Since(dnsStartTime).Seconds()
			if info.Err != nil {
				klog.Infof("DNS Error for %v in %v seconds: %v ", dnsHost, dnsLookupTime, info.Err)
				dnsError = info.Err.Error()
			}
		},
	}

	clientTraceCtx := httptrace.WithClientTrace(ctx, httpTrace)
	req, _ := http.NewRequestWithContext(clientTraceCtx, http.MethodGet, url, nil)

	startTime = time.Now()
	res, err := http.DefaultClient.Do(req)
	responseTime := time.Since(startTime).Seconds()
	defer res.Body.Close()
	var result string
	if err != nil {
		result = errorTypeMap(err)
	} else {
		result = statusCodeMap(res.StatusCode)
	}

	httpDurations.With(prometheus.Labels{
		"url":         url,
		"status_code": strconv.Itoa(res.StatusCode),
		"result":      result,
	}).Observe(responseTime)
	dnsLookupDurations.With(prometheus.Labels{
		"url":         url,
		"status_code": strconv.Itoa(res.StatusCode),
		"result":      result,
		"dns_error":   dnsError,
	}).Observe(dnsLookupTime)
	return err
}

func errorTypeMap(err error) string {
	if timeoutError, ok := err.(net.Error); ok && timeoutError.Timeout() {
		return "timeout"
	}
	urlErr, isUrlErr := err.(*url.Error)
	if !isUrlErr {
		return "connection_failed"
	}

	opErr, isNetErr := (urlErr.Err).(*net.OpError)
	if isNetErr {
		switch (opErr.Err).(type) {
		case *net.DNSError:
			return "dns_error"
		case *net.ParseError:
			return "address_error"
		}
	}

	return "connection_failed"
}

func statusCodeMap(sc int) string {
	switch {
	case sc >= 200 && sc < 400:
		return "http_success"
	case sc >= 400 && sc < 500:
		return "http_client_error"
	case sc >= 500:
		return "http_server_error"
	}
	return "http_other_error"
}
