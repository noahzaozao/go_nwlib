package nsq

type NsqConfig struct {
	Type string `yaml:"type"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	DB int `yaml:"db"`
	Password string `yaml:"password"`
}
