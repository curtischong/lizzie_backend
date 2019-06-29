package main

//TODO: use "github.com/pkg/errors"
import (
	"encoding/json"
	"fmt"
	//bioworker "github.com/curtischong/lizzie_server/bioworker"
	config "github.com/curtischong/lizzie_server/config"
	database "github.com/curtischong/lizzie_server/database"
	network "github.com/curtischong/lizzie_server/network"
	"io/ioutil"
	"log"
	"net/http"
)

type server struct {
	router *http.ServeMux
}

type DatabaseConfigObj = database.DatabaseConfigObj
type EmotionEvaluationObj = network.EmotionEvaluationObj
type EmotionEvaluationNetworkObj = network.EmotionEvaluationNetworkObj
type BioSamplesObj = network.BioSamplesObj
type MarkEventObj = network.MarkEventObj
type SkillObj = network.SkillObj
type ReviewObj = network.ReviewObj
type ScheduledReviewObj = network.ScheduledReviewObj
type TyperObj = network.TyperObj

// Watch
func (s server) routes(config DatabaseConfigObj) {
	s.router.HandleFunc("/typer_sent_field", s.typerSentFieldCall())
	s.router.HandleFunc("/messenger_sent_field", s.messengerSentFieldCall())
	s.router.HandleFunc("/upload_bio_samples", s.uploadBioSamplesCall(config))
	s.router.HandleFunc("/upload_emotion_evaluation", s.uploadEmotionEvaluationCall(config))
	s.router.HandleFunc("/upload_mark_event", s.uploadMarkEventCall(config))
	s.router.HandleFunc("/upload_skill", s.uploadSkillCall(config))
	s.router.HandleFunc("/upload_review", s.uploadReviewCall(config))
	s.router.HandleFunc("/upload_scheduled_review", s.uploadScheduledReviewCall(config))
	//s.router.HandleFunc("/admin", s.adminOnly(s.handleAdminIndex()))
}

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

func uploadBioSamplesCall(w http.ResponseWriter, response *http.Request, resCall string) http.HandlerFunc {

	return func(w http.ResponseWriter, response *http.Request) {
		body := getResponseBody(w, response)
		parsedResonse := BioSamplesObj{}
		jsonErr := json.Unmarshal(body, &parsedResonse)
		if jsonErr != nil {
			log.Println(body)
			log.Printf("error decoding emotion evaluation response: %v", jsonErr)
			if e, ok := jsonErr.(*json.SyntaxError); ok {
				log.Printf("syntax error at byte offset %d", e.Offset)
			}
			log.Printf("watch response: %q", body)
		}
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
		body := getResponseBody(w, response)
		parsedResonse := EmotionEvaluationNetworkObj{}
		jsonErr := json.Unmarshal(body, &parsedResonse)
		if jsonErr != nil {
			log.Println(body)
			log.Printf("error decoding emotion evaluation response: %v", jsonErr)
			if e, ok := jsonErr.(*json.SyntaxError); ok {
				log.Printf("syntax error at byte offset %d", e.Offset)
			}
			log.Printf("watch response: %q", body)
		}

		sliders := parsedResonse.EvalSliders
		for i := 0; i < len(sliders); i++ {
			println(i, sliders[i])
		}

		resObj := EmotionEvaluationObj{}

		//fmt.Println(parsedResonse.TiredEval)

		database.InsertEmotionEvaluationObj(resObj, config)
	}
}

func (s *server) uploadMarkEventCall(config DatabaseConfigObj) http.HandlerFunc {
	return func(w http.ResponseWriter, response *http.Request) {
		body := getResponseBody(w, response)
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

		database.InsertMarkEventObj(parsedResonse, config)
	}
}
func (s *server) uploadBioSamplesCall(config DatabaseConfigObj) http.HandlerFunc {
	return func(w http.ResponseWriter, response *http.Request) {
		body := getResponseBody(w, response)
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

func (s *server) uploadSkillCall(config DatabaseConfigObj) http.HandlerFunc {
	return func(w http.ResponseWriter, response *http.Request) {
		body := getResponseBody(w, response)
		parsedResonse := SkillObj{}
		jsonErr := json.Unmarshal(body, &parsedResonse)
		if jsonErr != nil {
			log.Println(body)
			log.Println("died here")
			log.Fatal(jsonErr)
		}

		fmt.Println(parsedResonse)
		database.InsertSkillObj(parsedResonse, config)
	}
}

func (s *server) uploadReviewCall(config DatabaseConfigObj) http.HandlerFunc {
	return func(w http.ResponseWriter, response *http.Request) {
		body := getResponseBody(w, response)
		parsedResonse := ReviewObj{}
		jsonErr := json.Unmarshal(body, &parsedResonse)
		if jsonErr != nil {
			log.Println(body)
			log.Println("died here")
			log.Fatal(jsonErr)
		}

		fmt.Println(parsedResonse)
		database.InsertReviewObj(parsedResonse, config)
	}
}

func (s *server) uploadScheduledReviewCall(config DatabaseConfigObj) http.HandlerFunc {
	return func(w http.ResponseWriter, response *http.Request) {
		body := getResponseBody(w, response)
		parsedResonse := ScheduledReviewObj{}
		jsonErr := json.Unmarshal(body, &parsedResonse)
		if jsonErr != nil {
			log.Println(body)
			log.Println("died here")
			log.Fatal(jsonErr)
		}

		fmt.Println(parsedResonse)
		database.InsertScheduledReviewObj(parsedResonse, config)
	}
}

func (s *server) typerSentFieldCall() http.HandlerFunc {
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
			log.Println("died here")
			log.Fatal(jsonErr)
		}
		fmt.Println(parsedResonse)
		database.InsertTyperObj(parsedResonse, config)
	}
}

func (s *server) messengerSentFieldCall() http.HandlerFunc {
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
			log.Println("died here")
			log.Fatal(jsonErr)
		}
		fmt.Println(parsedResonse)
		database.InsertTyperObj(parsedResonse, config)
	}
}

func main() {
	config := config.LoadConfiguration("../configSecrets/server_config.json")
	s := server{
		router: http.NewServeMux(),
	}
	s.routes(config.DatabaseConfigObj)
	log.Println("serving on port: " + config.ServerConfigObj.Port)
	log.Fatal(http.ListenAndServe(":"+config.ServerConfigObj.Port, s.router))

}
