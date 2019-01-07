package main

//TODO: use "github.com/pkg/errors"
import (
	"encoding/json"
	"fmt"
	bioworker "github.com/curtischong/lizzie_server/bioworker"
	serverutuls "github.com/curtischong/lizzie_server/serverutils"
	typerworker "github.com/curtischong/lizzie_server/typerworker"
	"github.com/influxdata/influxdb/client/v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Config struct {
	DatabaseConfig DatabaseConfigObj `json:"databaseconfig"`
}
type DatabaseConfigObj struct {
	Dbname   string `json:"dbname"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type server struct {
	router *http.ServeMux
}

// Watch
func (s server) routes(config DatabaseConfigObj) {
	s.router.HandleFunc("/watch_bio_snapshot", s.watchBioSnapshotCall())
	s.router.HandleFunc("/typer_sent_field", s.typerSentFieldCall())
	s.router.HandleFunc("/upload_emotion_evaluation", s.uploadEmotionEvaluationCall(config))
	s.router.HandleFunc("/upload_mark_event", s.uploadMarkEventCall(config))
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

		parsedResonse := bioworker.MarkEvent{}
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

func insertEmotionEvaluation(sample bioworker.EmotionEvaluation, config DatabaseConfigObj) {
	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://10.8.0.1:8086",
		Username: config.Username,
		Password: config.Password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  config.Dbname,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a point and add to batch
	//tags := map[string]string{"": "cpu-total"}
	fields := map[string]interface{}{
		"timeStartFillingForm": sample.TimeStartFillingForm,
		"normalEval":           sample.NormalEval,
		"socialEval":           sample.SocialEval,
		"exhaustedEval":        sample.ExhaustedEval,
		"tiredEval":            sample.TiredEval,
		"happyEval":            sample.HappyEval,
		"comments":             sample.Comments,
	}

	pt, err := client.NewPoint("emotionEvaluations", nil, fields, serverutuls.StringToDate(sample.TimeEndFillingForm))
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	// Write the batch
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}
	log.Printf("added point!")

	// Close client resources
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}

//TODO: udpate the type to MarkEvent
func insertMarkEvent(sample bioworker.MarkEvent, config DatabaseConfigObj) {
	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://10.8.0.1:8086",
		Username: config.Username,
		Password: config.Password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  config.Dbname,
		Precision: "ms",
	})
	if err != nil {
		log.Fatal(err)
	}

	var emotionRatings []int
	err2 := json.Unmarshal([]byte(sample.EmotionsFelt), &emotionRatings)
	if err != nil {
		log.Fatal(err2)
	}

	var typeBiometricsViewed []int
	err3 := json.Unmarshal([]byte(sample.TypeBiometricsViewed), &typeBiometricsViewed)
	if err != nil {
		log.Fatal(err3)
	}

	parsedTimeOfMark, err := strconv.ParseFloat(sample.TimeOfMark, 64)
	if err != nil {
		log.Fatal(err)
	}

	parsedTimeStartFillingForm, err := strconv.ParseFloat(sample.TimeStartFillingForm, 64)
	if err != nil {
		log.Fatal(err)
	}

	parsedTimeEndFillingForm, err := strconv.ParseFloat(sample.TimeEndFillingForm, 64)
	if err != nil {
		log.Fatal(err)
	}
	parsedAnticipationStart, err := strconv.ParseFloat(sample.AnticipationStart, 64)
	if err != nil {
		log.Fatal(err)
	}
	parsedTimeOfEvent, err := strconv.ParseFloat(sample.TimeOfEvent, 64)
	if err != nil {
		log.Fatal(err)
	}
	parsedReactionEnd, err := strconv.ParseFloat(sample.ReactionEnd, 64)
	if err != nil {
		log.Fatal(err)
	}

	// Create a point and add to batch
	//tags := map[string]string{"": "cpu-total"}
	fields := map[string]interface{}{
		"timeStartFillingForm": parsedTimeStartFillingForm,
		"timeEndFillingForm":   parsedTimeEndFillingForm,
		"isReaction":           sample.IsReaction,
		"anticipationStart":    parsedAnticipationStart,
		"timeOfEvent":          parsedTimeOfEvent,
		"reactionEnd":          parsedReactionEnd,
		"emotionsFeltFear":     emotionRatings[0],
		"emotionsFeltJoy":      emotionRatings[1],
		"emotionsFeltAnger":    emotionRatings[2],
		"emotionsFeltSad":      emotionRatings[3],
		"emotionsFeltDisgust":  emotionRatings[4],
		"emotionsFeltSuprise":  emotionRatings[5],
		"emotionsFeltContempt": emotionRatings[6],
		"emotionsFeltInterest": emotionRatings[7],
		"comments":             sample.Comments,
		"biometricsViewedHR":   typeBiometricsViewed[0],
	}

	pt, err := client.NewPoint("markEvents", nil, fields, time.Unix(0, int64(parsedTimeOfMark*1000000)))
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	// Write the batch
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}
	log.Printf("added point!")

	// Close client resources
	if err := c.Close(); err != nil {
		log.Fatal(err)
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

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

func main() {
	config := LoadConfiguration("../config/server_config.json")
	s := server{
		router: http.NewServeMux(),
	}
	s.routes(config.DatabaseConfig)
	log.Fatal(http.ListenAndServe(":9000", s.router))

}
