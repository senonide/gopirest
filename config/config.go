package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Port   int    `json:"port"`
	DbName string `json:"dbName"`
	DbUsr  string `json:"dbUser"`
	DbPw   string `json:"dbPw"`
	DbHost string `json:"dbHost"`
}

const ConfigFilePath = "config/config.json"

func ReadConfig() (*Config, error) {

	jsonFile, err := os.Open(ConfigFilePath)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var config Config
	return &config, json.Unmarshal(byteValue, &config)
}
