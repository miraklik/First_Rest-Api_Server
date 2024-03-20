package apiserver

type Config struct {
	BinAddr    string `toml:"bind_addr"`
	LogLevel   string `toml:"log_level"`
	DataBase   string `toml:"database_url"`
	sessionKey string `toml:"session_key"`
}

func NewConfig() *Config {
	return &Config{
		BinAddr:  ":8080",
		LogLevel: "debug",
	}
}
