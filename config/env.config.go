package config

import "github.com/spf13/viper"

type EnvConfig struct {
	Environment string `mapstructure:"ENVIRONMENT" default:"development"`
	Port        string `mapstructure:"PORT" default:"8080"`
	LoggerPath  string `mapstructure:"LOGGER_PATH" default:"./logs"`
	DbSource    string `mapstructure:"DB_SOURCE"`
}

func LoadEnvConfig(path string) (config *EnvConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
