package db

type Config struct {
	Name string `koanf:"name"`
	URI  string `koanf:"uri"`
}
