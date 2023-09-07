package initApp

type PostrgesqlConfig struct {
	Url             string `yaml:"url"`
	MaxConnIdleTime int    `yaml:"maxConnIdleTime"`
	MaxConns        int    `yaml:"maxConns"`
	MinConns        int    `yaml:"minConns"`
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
	Name             string           `yaml:"name"`
	PostrgesqlConfig PostrgesqlConfig `yaml:"postrgesql"`
	HttpConfig       HttpConfig       `yaml:"http"`
}
