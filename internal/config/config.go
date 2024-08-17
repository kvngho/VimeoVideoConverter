package config

import "github.com/spf13/viper"

type Config struct {
	DatabaseURL      string
	DatabaseUser     string
	DatabasePassword string
	DatabasePort     string
	DatabaseName     string
	WorkerNumber     int
	VimeoToken       string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{
		DatabaseURL:      viper.GetString("DATABASE_URL"),
		DatabaseUser:     viper.GetString("DATABASE_USER"),
		DatabaseName:     viper.GetString("DATABASE_NAME"),
		DatabasePort:     viper.GetString("DATABASE_PORT"),
		DatabasePassword: viper.GetString("DATABASE_PASSWORD"),
		WorkerNumber:     viper.GetInt("WORKER_NUMBER"),
		VimeoToken:       viper.GetString("VIMEO_TOKEN"),
	}, nil
}
