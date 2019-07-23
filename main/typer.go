package main

import (
	"encoding/json"
	"fmt"
	database "github.com/curtischong/lizzie_server/database"
	network "github.com/curtischong/lizzie_server/network"
	"log"
	"net/http"
)

type TyperObj = network.TyperObj
type MessengerObj = network.MessengerObj

func (s *server) typerSentFieldCall(config ConfigObj) http.HandlerFunc {
	return func(w http.ResponseWriter, response *http.Request) {
		body := getResponseBody(w, response)
		// NOTE: do not remove this line. Very good to debug incorrect encoding types from client
		// I thought the ajax was encoded in JSON but I didn't JSON.stringify it so it was urlencoding instead
		// fmt.Println(string(body))
		parsedResonse := TyperObj{}
		// NOTE: json.Unmarshal only works for json-encoded data
		jsonErr := json.Unmarshal(body, &parsedResonse)
		if jsonErr != nil {
			log.Println(body)
			log.Println("couldn't parse body")
			log.Println(jsonErr)
			w.WriteHeader(500)
			return
		}
		fmt.Println(parsedResonse)
		if database.InsertTyperObj(parsedResonse, config) {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(500)
		}
	}
}

func (s *server) messengerSentFieldCall(config ConfigObj) http.HandlerFunc {
	return func(w http.ResponseWriter, response *http.Request) {
		body := getResponseBody(w, response)
		// NOTE: do not remove this line. Very good to debug incorrect encoding types from client
		// I thought the ajax was encoded in JSON but I didn't JSON.stringify it so it was urlencoding instead
		// fmt.Println(string(body))
		parsedResonse := MessengerObj{}
		// NOTE: json.Unmarshal only works for json-encoded data
		jsonErr := json.Unmarshal(body, &parsedResonse)
		if jsonErr != nil {
			log.Println(body)
			log.Println("couldn't parse body")
			log.Println(jsonErr)
			w.WriteHeader(500)
			return
		}
		fmt.Println(parsedResonse)

		if database.InsertMessengerObj(parsedResonse, config) {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(500)
		}
	}
}
