package main

import (
	"encoding/json"
	"fmt"
	database "github.com/curtischong/lizzie_server/database"
	network "github.com/curtischong/lizzie_server/network"
	"log"
	"net/http"
)

type SkillObj = network.SkillObj
type ReviewObj = network.ReviewObj
type ScheduledReviewObj = network.ScheduledReviewObj

func (s *server) getPeaksSkills(config ConfigObj) http.HandlerFunc {
	return func(w http.ResponseWriter, response *http.Request) {
		enableCors(w, response)

		skills, skillsSucc := database.GetPeaksSkills(config)

		if skillsSucc {
			skillsJsonStr, _ := json.Marshal(skills)

			w.Write([]byte(skillsJsonStr))
		} else {
			w.WriteHeader(500)
		}
	}
}

func (s *server) uploadSkillCall(config ConfigObj) http.HandlerFunc {
	return func(w http.ResponseWriter, response *http.Request) {
		body := getResponseBody(w, response)
		parsedResonse := SkillObj{}
		jsonErr := json.Unmarshal(body, &parsedResonse)
		if jsonErr != nil {
			log.Println(body)
			log.Println("couldn't parse body")
			log.Println(jsonErr)
			w.WriteHeader(500)
			return
		}

		fmt.Println(parsedResonse)
		if database.InsertSkillObj(parsedResonse, config) {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(500)
		}
	}
}

func (s *server) uploadReviewCall(config ConfigObj) http.HandlerFunc {
	return func(w http.ResponseWriter, response *http.Request) {
		body := getResponseBody(w, response)
		parsedResonse := ReviewObj{}
		jsonErr := json.Unmarshal(body, &parsedResonse)
		if jsonErr != nil {
			log.Println(body)
			log.Println("couldn't parse body")
			log.Println(jsonErr)
			w.WriteHeader(500)
			return
		}

		fmt.Println(parsedResonse)
		if database.InsertReviewObj(parsedResonse, config) {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(500)
		}
	}
}

func (s *server) uploadScheduledReviewCall(config ConfigObj) http.HandlerFunc {
	return func(w http.ResponseWriter, response *http.Request) {
		body := getResponseBody(w, response)
		parsedResonse := ScheduledReviewObj{}
		jsonErr := json.Unmarshal(body, &parsedResonse)
		if jsonErr != nil {
			log.Println(body)
			log.Println("couldn't parse body")
			log.Println(jsonErr)
			w.WriteHeader(500)
			return
		}

		fmt.Println(parsedResonse)
		if database.InsertScheduledReviewObj(parsedResonse, config) {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(500)
		}
	}
}
