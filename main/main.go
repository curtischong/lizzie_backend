package main

//TODO: use "github.com/pkg/errors"
import (
	"encoding/json"
	"fmt"
	//bioworker "github.com/curtischong/lizzie_server/bioworker"
	config "github.com/curtischong/lizzie_server/config"
	database "github.com/curtischong/lizzie_server/database"
	network "github.com/curtischong/lizzie_server/network"
	typerworker "github.com/curtischong/lizzie_server/typerworker"
	"io/ioutil"
	"log"
	"net/http"
)

type server struct {
	router *http.ServeMux
}

type DatabaseConfigObj = database.DatabaseConfigObj
type EmotionEvaluationObj = network.EmotionEvaluationObj
type BioSamplesObj = network.BioSamplesObj
type MarkEventObj = network.MarkEventObj

// Watch
func (s server) routes(config DatabaseConfigObj) {
	s.router.HandleFunc("/typer_sent_field", s.typerSentFieldCall())
	s.router.HandleFunc("/upload_bio_samples", s.uploadBioSamplesCall(config))
	s.router.HandleFunc("/upload_emotion_evaluation", s.uploadEmotionEvaluationCall(config))
	s.router.HandleFunc("/upload_mark_event", s.uploadMarkEventCall(config))
	s.router.HandleFunc("/upload_skill", s.uploadSkillCall(config))
	s.router.HandleFunc("/upload_review", s.uploadReviewCall(config))
	//s.router.HandleFunc("/admin", s.adminOnly(s.handleAdminIndex()))
}

func uploadBioSamplesCall(w http.ResponseWriter, response *http.Request, resCall string) http.HandlerFunc {

	return func(w http.ResponseWriter, response *http.Request) {
		if response.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		body, readErr := ioutil.ReadAll(response.Body)
		if readErr != nil {
			log.Println("failed here")
			log.Fatal(readErr)
		}

		parsedResonse := EmotionEvaluationObj{}
		jsonErr := json.Unmarshal(body, &parsedResonse)
		if jsonErr != nil {
			log.Println(body)
			log.Printf("error decoding emotion evaluation response: %v", jsonErr)
			if e, ok := jsonErr.(*json.SyntaxError); ok {
				log.Printf("syntax error at byte offset %d", e.Offset)
			}
			log.Printf("watch response: %q", body)
		}
		return networkObj
	}

	/*
		if(resCall == "bioSample"){
			parsedResonss := NetworkObj{}
		}else if(resCall == "emotionEval"){

		}else if(resCall == "markEvent"){

		}else if(resCall == "skill"){

		}else if(resCall == "review"){

		}*/

}

func (s *server) uploadEmotionEvaluationCall(config DatabaseConfigObj) http.HandlerFunc {
	return func(w http.ResponseWriter, response *http.Request) {

		if response.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		body, readErr := ioutil.ReadAll(response.Body)
		if readErr != nil {
			log.Println("failed here")
			log.Fatal(readErr)
		}

		parsedResonse := EmotionEvaluationObj{}
		jsonErr := json.Unmarshal(body, &parsedResonse)
		if jsonErr != nil {
			log.Println(body)
			log.Printf("error decoding emotion evaluation response: %v", jsonErr)
			if e, ok := jsonErr.(*json.SyntaxError); ok {
				log.Printf("syntax error at byte offset %d", e.Offset)
			}
			log.Printf("watch response: %q", body)
		}

		fmt.Println(parsedResonse.TiredEval)
		fmt.Fprintf(w, "bio snapshot")

		database.InsertEmotionEvaluationObj(parsedResonse, config)
	}
}
func (s *server) uploadMarkEventCall(config DatabaseConfigObj) http.HandlerFunc {
	return func(w http.ResponseWriter, response *http.Request) {
		if response.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		body, readErr := ioutil.ReadAll(response.Body)
		if readErr != nil {
			log.Println("failed here")
			log.Fatal(readErr)
		}

		parsedResonse := MarkEventObj{}
		jsonErr := json.Unmarshal(body, &parsedResonse)
		if jsonErr != nil {
			log.Println(body)
			log.Printf("error decoding watch response: %v", jsonErr)
			if e, ok := jsonErr.(*json.SyntaxError); ok {
				log.Printf("syntax error at byte offset %d", e.Offset)
			}
			log.Printf("watch response: %q", body)
		}

		fmt.Println(parsedResonse.IsReaction)
		fmt.Fprintf(w, "bio snapshot")

		database.InsertMarkEventObj(parsedResonse, config)
	}
}
func (s *server) uploadBioSamplesCall(config DatabaseConfigObj) http.HandlerFunc {
	return func(w http.ResponseWriter, response *http.Request) {
		if response.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		body, readErr := ioutil.ReadAll(response.Body)
		if readErr != nil {
			log.Println("failed here")
			log.Fatal(readErr)
		}

		parsedResonse := BioSamplesObj{}
		jsonErr := json.Unmarshal(body, &parsedResonse)
		if jsonErr != nil {
			log.Println(body)
			log.Printf("error decoding watch response: %v", jsonErr)
			if e, ok := jsonErr.(*json.SyntaxError); ok {
				log.Printf("syntax error at byte offset %d", e.Offset)
			}
			log.Printf("watch response: %q", body)
		}

		//fmt.Println(parsedResonse)
		fmt.Fprintf(w, "bio snapshot")

		database.InsertBioSamplesObj(parsedResonse, config)
	}
}

func (s *server) uploadSkillCall() http.HandlerFunc {
	return func(w http.ResponseWriter, response *http.Request) {
		if response.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		body, readErr := ioutil.ReadAll(response.Body)
		if readErr != nil {
			log.Println("failed here")
			log.Fatal(readErr)
		}

		parsedResonse := typerworker.SentField{}
		jsonErr := json.Unmarshal(body, &parsedResonse)
		if jsonErr != nil {
			log.Println(body)
			log.Println("died here")
			log.Fatal(jsonErr)
		}

		fmt.Println(parsedResonse.Url)
		database.InsertSkillObj(parsedResonse, config)
	}
}

func (s *server) uploadReviewCall() http.HandlerFunc {
	return func(w http.ResponseWriter, response *http.Request) {
		if response.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		body, readErr := ioutil.ReadAll(response.Body)
		if readErr != nil {
			log.Println("failed here")
			log.Fatal(readErr)
		}

		parsedResonse := typerworker.SentField{}
		jsonErr := json.Unmarshal(body, &parsedResonse)
		if jsonErr != nil {
			log.Println(body)
			log.Println("died here")
			log.Fatal(jsonErr)
		}

		fmt.Println(parsedResonse.Url)
		database.InsertReviewObj(parsedResonse, config)
	}
}

func (s *server) typerSentFieldCall() http.HandlerFunc {
	return func(w http.ResponseWriter, response *http.Request) {
		if response.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		body, readErr := ioutil.ReadAll(response.Body)
		if readErr != nil {
			log.Println("failed here")
			log.Fatal(readErr)
		}

		parsedResonse := typerworker.SentField{}
		jsonErr := json.Unmarshal(body, &parsedResonse)
		if jsonErr != nil {
			log.Println(body)
			log.Println("died here")
			log.Fatal(jsonErr)
		}

		fmt.Println(parsedResonse.Url)
	}
}

func main() {
	config := config.LoadConfiguration("../config/server_config.json")
	s := server{
		router: http.NewServeMux(),
	}
	s.routes(config.DatabaseConfigObj)
	log.Println("asd")
	log.Fatal(http.ListenAndServe(":9000", s.router))

}
