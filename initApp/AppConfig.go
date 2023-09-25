package initApp

type MongoConfig struct {
	URL      string `yaml:"url"`
	Database string `yaml:"database"`
}

type HttpTimeout struct {
	Read  int `yaml:"read"`
	Write int `yaml:"write"`
	Idle  int `yaml:"idle"`
}

type HttpConfig struct {
	Port        int         `yaml:"port"`
	HttpTimeout HttpTimeout `yaml:"timeout"`
}

type AppConfig struct {
	Name        string      `yaml:"name"`
	MongoConfig MongoConfig `yaml:"mongo"`
	HttpConfig  HttpConfig  `yaml:"http"`
}
