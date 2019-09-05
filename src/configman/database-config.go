package configman

type DatabaseConfig struct {
	Type     string `yaml:"type"`
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"db"`
}
