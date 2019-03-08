package config

import (
	"encoding/json"
	"fmt"
	//database "github.com/curtischong/lizzie_server/database"
	"os"
)

//type DatabaseConfigObj = database.DatabaseConfigObj

type Config struct {
	DatabaseConfigObj DatabaseConfigObj `json:"databaseconfig"`
}

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
