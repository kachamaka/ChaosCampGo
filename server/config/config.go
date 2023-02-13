package config

import "github.com/spf13/viper"

// Config is a struct model that represents the config properties being read from the app.env file
type Config struct {
	ServerAddress   string `mapstructure:"HTTP_ADDRESS"`
	DatabaseAddress string `mapstructure:"DB_ADDRESS"`
	DatabaseName    string `mapstructure:"DB_NAME"`
	TokenSecret     string `mapstructure:"TOKEN_SECRET"`
	SendgridAPiKey  string `mapstructure:"SENDGRID_API_KEY"`
}

// LoadConfig is a function that loads the config properties and returns a struct entailing them
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
