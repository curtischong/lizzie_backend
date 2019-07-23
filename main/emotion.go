package main

import (
	"encoding/json"
	"fmt"
	database "github.com/curtischong/lizzie_server/database"
	network "github.com/curtischong/lizzie_server/network"
	"log"
	"net/http"
)

type EmotionEvaluationObj = network.EmotionEvaluationObj
type EmotionEvaluationNetworkObj = network.EmotionEvaluationNetworkObj

func (s *server) uploadEmotionEvaluationCall(config ConfigObj) http.HandlerFunc {
	return func(w http.ResponseWriter, response *http.Request) {
		body := getResponseBody(w, response)
		parsedResonse := EmotionEvaluationObj{}
		jsonErr := json.Unmarshal(body, &parsedResonse)
		if jsonErr != nil {
			log.Println(body)
			log.Println("couldn't parse body")
			log.Println(jsonErr)
			w.WriteHeader(500)
			return
		}
		fmt.Println(parsedResonse)

		if database.InsertEmotionEvaluationObj(parsedResonse, config) {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(500)
		}
	}
}
