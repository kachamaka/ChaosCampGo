package database

import "github.com/spf13/viper"

type Config struct {
	ServerAddress   string `mapstructure:"HTTP_ADDRESS"`
	DatabaseAddress string `mapstructure:"DB_ADDRESS"`
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	config := &Config{}

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
