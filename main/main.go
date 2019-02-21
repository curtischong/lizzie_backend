package main

//TODO: use "github.com/pkg/errors"
import (
	"encoding/json"
	"fmt"
	//bioworker "github.com/curtischong/lizzie_server/bioworker"
	config "github.com/curtischong/lizzie_server/config"
	serverutuls "github.com/curtischong/lizzie_server/serverutils"
	typerworker "github.com/curtischong/lizzie_server/typerworker"
	"github.com/influxdata/influxdb/client/v2"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type server struct {
	router *http.ServeMux
}

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

		parsedResonse := EmotionEvaluation{}
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

		insertEmotionEvaluation(parsedResonse, config)
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

		parsedResonse := MarkEvent{}
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

		insertMarkEvent(parsedResonse, config)
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

		parsedResonse := BioSamples{}
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

		insertBioSamples(parsedResonse, config)
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
	s.routes(config.DatabaseConfig)
	log.Println("asd")
	log.Fatal(http.ListenAndServe(":9000", s.router))

}
