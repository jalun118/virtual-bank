package app

import "github.com/spf13/viper"

func NewViper() *viper.Viper {
	config := viper.New()

	config.SetConfigFile(".env")
	config.AddConfigPath("./../")
	config.AddConfigPath("./")
	err := config.ReadInConfig()

	if err != nil {
		panic(err)
	}

	return config
}
