package main

//TODO: use "github.com/pkg/errors"
import (
	"fmt"
	//bioworker "github.com/curtischong/lizzie_server/bioworker"
	config "github.com/curtischong/lizzie_server/config"
	database "github.com/curtischong/lizzie_server/database"
	"io/ioutil"
	"log"
	"net/http"
)

type server struct {
	router *http.ServeMux
}

// General

type DBConfigObj = database.DBConfigObj
type ConfigObj = config.ConfigObj
type DBObj = database.DBObj

func getResponseBody(w http.ResponseWriter, response *http.Request) []byte {
	if response.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return nil
	}
	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Println("failed here")
		log.Fatal(readErr)
	}
	return body
}

func enableCors(w http.ResponseWriter, response *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=ascii")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	if response.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	}
	return
}

/*
func handleSuccess(w http.ResponseWriter, success bool) {
	if success {
		log.Println("yes")
		w.Write([]byte("hi"))
		w.WriteHeader(200)
	} else {
		log.Println("nope")
		w.WriteHeader(500)
	}
}*/

func main() {
	config := config.LoadConfiguration("../configSecrets/server_config.json")
	fmt.Printf("IsDev: %t\n", config.ServerConfig.IsDev)
	s := server{
		router: http.NewServeMux(),
	}

	database.SetupDBConfig(&config)

	s.routes(config)
	fmt.Printf("serving on port: %s\n", config.ServerConfig.Port)
	log.Fatal(http.ListenAndServe(":"+config.ServerConfig.Port, s.router))
}
