package main

//TODO: use "github.com/pkg/errors"
import (
	"encoding/json"
	"fmt"
	bioworker "github.com/curtischong/lizzie_server/bioworker"
	typerworker "github.com/curtischong/lizzie_server/typerworker"
	"io/ioutil"
	"log"
	"net/http"
)

type server struct {
	router *http.ServeMux
}

// Watch
func (s server) routes() {
	s.router.HandleFunc("/watch_bio_snapshot", s.watchBioSnapshotCall())
	s.router.HandleFunc("/typer_sent_field", s.typerSentFieldCall())
	s.router.HandleFunc("/upload_emotion_evaluation", s.uploadEmotionEvaluationCall())
	//s.router.HandleFunc("/admin", s.adminOnly(s.handleAdminIndex()))
}
func (s *server) uploadEmotionEvaluationCall() http.HandlerFunc {
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

		parsedResonse := bioworker.EmotionEvaluation{}
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
	}
}
func (s *server) watchBioSnapshotCall() http.HandlerFunc {
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

		parsedResonse := bioworker.BioSnapshot{}
		jsonErr := json.Unmarshal(body, &parsedResonse)
		if jsonErr != nil {
			log.Println(body)
			log.Printf("error decoding watch response: %v", jsonErr)
			if e, ok := jsonErr.(*json.SyntaxError); ok {
				log.Printf("syntax error at byte offset %d", e.Offset)
			}
			log.Printf("watch response: %q", body)
		}

		fmt.Println(parsedResonse.Heartrate)
		fmt.Fprintf(w, "bio snapshot")
	}
}

// Typer
/*func (s *server) typerSentFieldCall() http.HandlerFunc {
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
}*/

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
	s := server{
		router: http.NewServeMux(),
	}
	s.routes()
	log.Fatal(http.ListenAndServe(":9000", s.router))

}
