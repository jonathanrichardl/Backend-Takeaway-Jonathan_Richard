package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type DatabaseConfig struct {
	Username     string
	Password     string
	Address      string
	DatabaseName string
}
type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	Redis struct {
		Host       string `yaml:"host"`
		Port       int    `yaml:"port"`
		Password   string `yaml:"password"`
		Expiration int    `yaml:"expiration"`
		DatabaseNo int    `yaml:"databaseno"`
	} `yaml:"redis"`
}

func LoadDatabaseConfiguration() (DatabaseConfig, error) {
	var result DatabaseConfig
	result.Username = os.Getenv("USER_NAME")
	result.Password = os.Getenv("PASSWORD")
	result.Address = os.Getenv("ADRESS")
	result.Address = os.Getenv("DATABASE_NAME")
	return result, nil
}

func LoadServerConfiguration() (*Config, error) {
	config := &Config{}

	// Open config file
	file, err := os.Open("../config/config.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil

}
