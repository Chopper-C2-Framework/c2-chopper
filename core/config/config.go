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
}

func CreateDefaultConfig() *Config {
	return &Config{
		PluginsDir:    "~/.c2chopper/plugins",
		ClientPort:    9001,
		ServerPort:    9002,
		Host:          "localhost",
		ServerCertKey: "cert/server-key.pem",
		ServerCert:    "cert/server-cert.pem",
	}
}
