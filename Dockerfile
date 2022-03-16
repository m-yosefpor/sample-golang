ARG GO_VERSION=1.17

# 1. Build Container
# We are not using Alpine for build stage as it does not include race-detection libraries
# which is required when running tests.
FROM golang:${GO_VERSION} AS builder

ENV GOOS=linux \
    GOARCH=amd64

WORKDIR /src

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
COPY go.sum go.mod /src/
RUN go mod download

COPY . /src
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /app/main


# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /app/main /
USER 65532:65532

CMD ["/main"]
