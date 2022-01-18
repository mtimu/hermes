package emq

type Config struct {
	ClientID string `koanf:"client_id"`
	URL      string `koanf:"url"`
}
