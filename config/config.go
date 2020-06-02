package config

// AppConfig 应用配置
type (
	AppConfig struct{
		DBConfig struct{
			Host string `yaml:"host"`
			Port     string `yaml:"port"`
			UserName string `yaml:"username"`
			Password string `yaml:"password"`
			DataBase string `yaml:"database"`
		}
	}
)