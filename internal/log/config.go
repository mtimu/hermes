package log

type Config struct {
	Production bool   `koanf:"production"`
	Encoding   string `koanf:"encoding"`
	Level      string `koanf:"level"`
}
