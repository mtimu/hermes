package log

type Config struct {
	Production bool   `koanf:"production_env"`
	Encoding   string `koanf:"encoding"`
	Level      string `koanf:"level"`
}
