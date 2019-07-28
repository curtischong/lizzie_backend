package main

//package todoist_worker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Configuration : config
type Configuration struct {
	Token string
	Url   string
}

func getTasks(config Configuration) {
	url := config.Url + "tasks?token=" + config.Token
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func readConfig() Configuration {
	file, err1 := os.Open("../configSecrets/todoist.json")
	if err1 != nil {
		log.Fatalf("Unable to read todoist config file: %v", err1)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err2 := decoder.Decode(&configuration)
	if err2 != nil {
		fmt.Println("error:", err2)
	}
	return configuration
}

func startTodoist() {
	config := readConfig()
	fmt.Println("asd")
	getTasks(config)

}

func main() {
	startTodoist()
}
