package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database DatabaseConfiguration
	Logging  LoggingConfiguration
}

type DatabaseConfiguration struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"dbname"`
}

type LoggingConfiguration struct {
	Database bool
	Filename string
}

var configuration *Config

// Function that initializes the environment configuration
func InitializeConfiguration() *Config {

	if configuration == nil {
		yamlFile, err := ioutil.ReadFile("config.yaml")

		if err != nil {
			log.Printf("yamlFile.Get err   #%v ", err)
		}

		err = yaml.Unmarshal(yamlFile, &configuration)

		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
	}
	return configuration
}

// Returns the database connection string
func GetConnectionString() string {

	configuration := InitializeConfiguration()

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configuration.Database.User, configuration.Database.Password, configuration.Database.Host, configuration.Database.Port, configuration.Database.Name)

}
