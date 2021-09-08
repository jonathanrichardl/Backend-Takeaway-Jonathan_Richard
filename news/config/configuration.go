package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type DatabaseConfig struct {
	DatabaseManagementSystem string `json:"databaseManagementSystem"`
	Username                 string `json:"username"`
	Password                 string `json:"password"`
	Address                  string `json:"address"`
	DatabaseName             string `json:"databaseName"`
}

func LoadDatabaseConfiguration() (*DatabaseConfig, error) {
	jsonFile, err := os.Open("database.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, err
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result DatabaseConfig
	json.Unmarshal([]byte(byteValue), &result)
	return &result, nil

}
