package configman

import "fmt"

type DatabaseConfig struct {
	Type     string `yaml:"type"`
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"db"`
}

func (cfg DatabaseConfig) GetDBUri() (string) {
	if cfg.Type == "sqlite3" {
		return cfg.Address
	} else {
		return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
			cfg.Address, cfg.Port, cfg.Username, cfg.Database, cfg.Password)
	}
}
