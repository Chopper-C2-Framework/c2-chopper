package config

const (
	CONFIG_DEFAULT_DIR  = ".c2-chopper"
	CONFIG_DEFAULT_FILE = "config.yaml"
	CONFIG_DEFAULT_PATH = ".c2-chopper/config.yaml"
)

type Config struct {
	PluginsDir    string `yaml:"plugins_path"`
	ClientPort    int    `yaml:"client_port"`
	ServerPort    int    `yaml:"server_port"`
	Host          string `yaml:"host"`
	ServerCert    string `yaml:"server_cert_path"`
	ServerCertKey string `yaml:"sever_cert_key_path"`
	UseTLS        bool   `yaml:"use_tls"`
	ServerDb      string `yaml:"server_db_path"`
}

var DefaultConfig = Config{
	PluginsDir:    "~/.c2chopper/plugins",
	ClientPort:    9001,
	ServerPort:    9002,
	Host:          "localhost",
	ServerCert:    "./cert/server-cert.pem",
	ServerCertKey: "./cert/server-key.pem",
	UseTLS:        false,
	ServerDb:      "server.db",
}

func CreateDefaultConfig() *Config {
	cfg := DefaultConfig
	return &cfg
}
