package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig `mapstructure:"server"`
	Log    LogConfig    `mapstructure:"log"`
}

type ServerConfig struct {
	Addr    string `mapstructure:"addr"`
	MaxConn int    `mapstructure:"max_conn"`
}

type LogConfig struct {
	// 日志文件路径，空则只输出到控制台
	Filename string `mapstructure:"filename"`
	// 日志级别：debug / info / warn / error，默认 info
	Level string `mapstructure:"level"`
	// 日志格式：json / text，默认 json
	Format string `mapstructure:"format"`
	// 单文件最大体积（MB），默认 100
	MaxSize int `mapstructure:"max_size"`
	// 最多保留文件数，默认 7
	MaxBackups int `mapstructure:"max_backups"`
	// 最多保留天数，默认 30
	MaxAge int `mapstructure:"max_age"`
	// 是否压缩归档，默认 true
	Compress bool `mapstructure:"compress"`
}

func Load(path string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType(configType(path))

	v.SetDefault("server.addr", ":8080")
	v.SetDefault("server.max_conn", 0)
	v.SetDefault("log.level", "info")
	v.SetDefault("log.format", "json")
	v.SetDefault("log.max_size", 100)
	v.SetDefault("log.max_backups", 7)
	v.SetDefault("log.max_age", 30)
	v.SetDefault("log.compress", true)

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func configType(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '.' {
			ext := path[i+1:]
			if ext == "yml" {
				return "yaml"
			}
			return ext
		}
	}
	return "yaml"
}
