package profiler

type Config struct {
	Enabled bool   `koanf:"enabled"`
	Address string `koanf:"address"`
}
