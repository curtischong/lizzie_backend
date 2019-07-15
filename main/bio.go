package main

import (
	"encoding/json"
	"fmt"
	database "github.com/curtischong/lizzie_server/database"
	network "github.com/curtischong/lizzie_server/network"
	"log"
	"net/http"
)

type BioSamplesObj = network.BioSamplesObj

func (s *server) uploadBioSamplesCall(config ConfigObj) http.HandlerFunc {
	return func(w http.ResponseWriter, response *http.Request) {
		body := getResponseBody(w, response)
		parsedResonse := BioSamplesObj{}
		jsonErr := json.Unmarshal(body, &parsedResonse)
		if jsonErr != nil {
			/*log.Println(body)
			log.Printf("error decoding watch response: %v", jsonErr)
			if e, ok := jsonErr.(*json.SyntaxError); ok {
				log.Printf("syntax error at byte offset %d", e.Offset)
			}
			log.Printf("watch response: %q", body)*/
			log.Println(body)
			log.Println("couldn't parse body")
			log.Println(jsonErr)
			w.WriteHeader(500)
			return
		}

		//fmt.Println(parsedResonse)
		fmt.Fprintf(w, "bio snapshot")

		if database.InsertBioSamplesObj(parsedResonse, config) {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(500)
		}
	}
}
