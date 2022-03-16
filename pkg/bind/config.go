package bind

type Config struct {
	Host string `koanf:"host"`
	Port int    `koanf:"port"`
}
