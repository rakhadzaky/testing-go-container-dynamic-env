package main

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func main() {
	filePath := "config.dev.yaml"
	if os.Getenv("SERVICE_ENV") != "" {
		filePath = os.Getenv("SERVICE_ENV")
	}
	cfg := viper.New()
	cfg.SetConfigFile(filePath)

	if err := cfg.ReadInConfig(); err != nil {
		log.Fatal(err, "Error while reading config file")
	}

	log.Printf("Check filepath: %s and its value: %s\n", filePath, cfg.Get("serviceEnv"))
}
