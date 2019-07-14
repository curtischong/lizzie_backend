package main

//TODO: use "github.com/pkg/errors"
import (
	"encoding/json"
	"fmt"
	//bioworker "github.com/curtischong/lizzie_server/bioworker"
	config "github.com/curtischong/lizzie_server/config"
	database "github.com/curtischong/lizzie_server/database"
	network "github.com/curtischong/lizzie_server/network"
	utils "github.com/curtischong/lizzie_server/serverUtils"
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

// Typer

type TyperObj = network.TyperObj
type MessengerObj = network.MessengerObj

// LNews

type GetCardsAndPanelsObj = network.GetCardsAndPanelsObj

// BioSamplesObj

type BioSamplesObj = network.BioSamplesObj

// Emotions

type EmotionEvaluationObj = network.EmotionEvaluationObj
type EmotionEvaluationNetworkObj = network.EmotionEvaluationNetworkObj

// Events

type MarkEventObj = network.MarkEventObj

// Lizzie Peaks

type SkillObj = network.SkillObj
type ReviewObj = network.ReviewObj
type ScheduledReviewObj = network.ScheduledReviewObj

// TODO: think about standardizing these endpoint names
func (s server) routes(config ConfigObj) {
	s.router.HandleFunc("/get_cards_and_panels", s.getCardsAndPanelsCall(config))
	s.router.HandleFunc("/typer_sent_field", s.typerSentFieldCall(config))
	s.router.HandleFunc("/messenger_sent_text", s.messengerSentFieldCall(config))
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

func enableCors(w http.ResponseWriter, response *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=ascii")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	if response.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	}
	return
}

func (s *server) getCardsAndPanelsCall(config ConfigObj) http.HandlerFunc {
	return func(w http.ResponseWriter, response *http.Request) {
		enableCors(w, response)

		q := response.URL.Query()
		parsedResonse := GetCardsAndPanelsObj{
			CardAmount:  utils.BetterAtoi(q.Get("cardAmount")),
			CardOffset:  utils.BetterAtoi(q.Get("cardOffset")),
			PanelAmount: utils.BetterAtoi(q.Get("panelAmount")),
			PanelOffset: utils.BetterAtoi(q.Get("panelOffset")),
		}

		cards, cardsSucc := database.GetCards(parsedResonse, config)
		panels, panelsSucc := database.GetPanels(parsedResonse, config)
		//log.Println(panelsSucc)

		if cardsSucc && panelsSucc {
			cardsAndPanelsObj := map[string][]map[string]string{"cards": cards, "panels": panels}
			cardsAndPanelsJsonStr, _ := json.Marshal(cardsAndPanelsObj)

			w.Write([]byte(cardsAndPanelsJsonStr))
		} else {
			w.WriteHeader(500)
		}
	}
}

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

func (s *server) uploadMarkEventCall(config ConfigObj) http.HandlerFunc {
	return func(w http.ResponseWriter, response *http.Request) {
		body := getResponseBody(w, response)
		parsedResonse := MarkEventObj{}
		jsonErr := json.Unmarshal(body, &parsedResonse)
		if jsonErr != nil {
			log.Println(body)
			log.Println("couldn't parse body")
			log.Println(jsonErr)
			w.WriteHeader(500)
			return
		}

		fmt.Println(parsedResonse.IsReaction)

		if database.InsertMarkEventObj(parsedResonse, config) {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(500)
		}
	}
}
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
		//handleSuccess(w, database.InsertMessengerObj(parsedResonse, config, db))

		//w.Header().Set("Content-Type", "application/json")
		if database.InsertMessengerObj(parsedResonse, config) {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(500)
		}
	}
}

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
