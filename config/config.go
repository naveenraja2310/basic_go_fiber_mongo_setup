/*
config.go
Author: Naveenraj O M
Description: This file handles the configuration settings of the application,
including loading configurations from a YAML file and mapping them to structured Go struct.
*/
package config

import (
	"log"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

var Config Configuration

type Configuration struct {
	DBUri       string `mapstructure:"DB_URI"`
	DBName      string `mapstructure:"DB_NAME"`
	Host        string `mapstructure:"HOST"`
	Port        int    `mapstructure:"PORT"`
	JwtSecret   string `mapstructure:"JWTSECRET"`
	Environment string `mapstructure:"ENVIRONMENT"`
}

func LoadConfig() {
	// Get the project root directory
	_, b, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(b), "../")

	viper.SetConfigName(".env")      // Name of config file
	viper.SetConfigType("env")       // Type of config file
	viper.AddConfigPath(projectRoot) // Path to look for the config file in
	viper.AddConfigPath(".")         // Optionally look for config in the working directory

	// Enable viper to read Environment Variables
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatalf("error reading config file: %v", err)
		}
		log.Fatalf("error reading config file: %v", err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalf("unable to decode config into struct: %v", err)
	}
}
