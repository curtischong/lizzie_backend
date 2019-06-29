package config

import (
	"encoding/json"
	"fmt"
	//database "github.com/curtischong/lizzie_server/database"
	"os"
)

//type DatabaseConfigObj = database.DatabaseConfigObj

type ConfigObj struct {
	DBConfig     DBConfigObj     `json:"databaseConfig"`
	ServerConfig ServerConfigObj `json:"serverConfig"`
}

func LoadConfiguration(file string) ConfigObj {
	var config ConfigObj
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
