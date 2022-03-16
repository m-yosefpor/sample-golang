package metrics

type Metrics struct{}

func New(cfg Config) (*Metrics, error) {
	return &Metrics{}, nil
}
