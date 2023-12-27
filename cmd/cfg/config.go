package cfg

type Config struct {
	ServerEndpoint string
	DataBaseDSN    string
	NATSEndpoint   string
	ClusterID      string
	ClientID       string
	Channel        string
}

type Option func(*Config)

func NewConfig(option ...Option) *Config {
	cfg := &Config{
		ServerEndpoint: "localhost:8000",
		DataBaseDSN:    "postgres://postgres:qwerty@localhost:5430/postgres?sslmode=disable",
		NATSEndpoint:   "nats://localhost:4223",
		ClusterID:      "test-cluster",
		ClientID:       "client_1",
		Channel:        "channel",
	}

	for _, opt := range option {
		opt(cfg)
	}

	return cfg
}

func WithServerEndpoint(endpoint string) Option {
	return func(c *Config) {
		c.ServerEndpoint = endpoint
	}
}

func WithDataBaseDSN(dsn string) Option {
	return func(c *Config) {
		c.DataBaseDSN = dsn
	}
}

func WithNATSEndpoint(NATSEndpoint string) Option {
	return func(c *Config) {
		c.NATSEndpoint = NATSEndpoint
	}
}

func WithClusterID(clusterID string) Option {
	return func(c *Config) {
		c.ClusterID = clusterID
	}
}

func WithClientID(clientID string) Option {
	return func(c *Config) {
		c.ClusterID = clientID
	}
}

func WithChannel(channel string) Option {
	return func(c *Config) {
		c.Channel = channel
	}
}
