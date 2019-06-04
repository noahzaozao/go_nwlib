package database

type DBConfig struct {
	Type string `yaml:"type"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Database string `yaml:"database"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
}
