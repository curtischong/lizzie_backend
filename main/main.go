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

const DEV = true

type server struct {
	router *http.ServeMux
}

type DBConfigObj = database.DBConfigObj
type ConfigObj = config.ConfigObj
type DBObj = database.DBObj
type EmotionEvaluationObj = network.EmotionEvaluationObj
type EmotionEvaluationNetworkObj = network.EmotionEvaluationNetworkObj
type BioSamplesObj = network.BioSamplesObj
type MarkEventObj = network.MarkEventObj
type SkillObj = network.SkillObj
type ReviewObj = network.ReviewObj
type ScheduledReviewObj = network.ScheduledReviewObj
type TyperObj = network.TyperObj
type MessengerObj = network.MessengerObj

// Watch
func (s server) routes(config ConfigObj, db DBObj) {
	s.router.HandleFunc("/typer_sent_field", s.typerSentFieldCall(config, db))
	s.router.HandleFunc("/messenger_sent_text", s.messengerSentFieldCall(config, db))
	s.router.HandleFunc("/upload_bio_samples", s.uploadBioSamplesCall(config, db))
	s.router.HandleFunc("/upload_emotion_evaluation", s.uploadEmotionEvaluationCall(config, db))
	s.router.HandleFunc("/upload_mark_event", s.uploadMarkEventCall(config, db))
	s.router.HandleFunc("/upload_skill", s.uploadSkillCall(config, db))
	s.router.HandleFunc("/upload_review", s.uploadReviewCall(config, db))
	s.router.HandleFunc("/upload_scheduled_review", s.uploadScheduledReviewCall(config, db))
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

func (s *server) uploadEmotionEvaluationCall(config ConfigObj, db DBObj) http.HandlerFunc {
	return func(w http.ResponseWriter, response *http.Request) {
		body := getResponseBody(w, response)
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
		/*
			sliders := parsedResonse.EvalSliders
			for i := 0; i < len(sliders); i++ {
				println(i, sliders[i])
			}

			resObj := EmotionEvaluationObj{}

			//fmt.Println(parsedResonse.TiredEval)

			if database.InsertEmotionEvaluationObj(parsedResonse, config, db) {
				w.WriteHeader(200)
			} else {
				w.WriteHeader(500)
			}*/
	}
}

func (s *server) uploadMarkEventCall(config ConfigObj, db DBObj) http.HandlerFunc {
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

		if database.InsertMarkEventObj(parsedResonse, config, db) {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(500)
		}
	}
}
func (s *server) uploadBioSamplesCall(config ConfigObj, db DBObj) http.HandlerFunc {
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

		if database.InsertBioSamplesObj(parsedResonse, config, db) {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(500)
		}
	}
}

func (s *server) uploadSkillCall(config ConfigObj, db DBObj) http.HandlerFunc {
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
		if database.InsertSkillObj(parsedResonse, config, db) {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(500)
		}
	}
}

func (s *server) uploadReviewCall(config ConfigObj, db DBObj) http.HandlerFunc {
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
		if database.InsertReviewObj(parsedResonse, config, db) {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(500)
		}
	}
}

func (s *server) uploadScheduledReviewCall(config ConfigObj, db DBObj) http.HandlerFunc {
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
		if database.InsertScheduledReviewObj(parsedResonse, config, db) {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(500)
		}
	}
}

func (s *server) typerSentFieldCall(config ConfigObj, db DBObj) http.HandlerFunc {
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
		if database.InsertTyperObj(parsedResonse, config, db) {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(500)
		}
	}
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

func (s *server) messengerSentFieldCall(config ConfigObj, db DBObj) http.HandlerFunc {
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
		//handleSuccess(w, database.InsertMessengerObj(parsedResonse, config, db))

		//w.Header().Set("Content-Type", "application/json")
		if database.InsertMessengerObj(parsedResonse, config, db) {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(500)
		}
	}
}

func main() {
	config := config.LoadConfiguration("../configSecrets/server_config.json")
	fmt.Printf("IsDev: %t \n", config.ServerConfig.IsDev)
	s := server{
		router: http.NewServeMux(),
	}

	db := database.SetupDB(config)

	s.routes(config, db)
	fmt.Printf("serving on port: %s", config.ServerConfig.Port)
	log.Fatal(http.ListenAndServe(":"+config.ServerConfig.Port, s.router))

	// write a case to close this connection
	//database.closeDBConnection(&DBObj.dbClient)

}
