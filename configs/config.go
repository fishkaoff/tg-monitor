package config


import (
	"log"
	"github.com/spf13/viper"

)

type Config struct {
	TGTOKEN string `mapstructure:"TGTOKEN"`
	DBURL string `mapstructure:"DBURL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("preferences")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		log.Fatal(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}
	return
}