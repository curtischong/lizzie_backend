package main

//TODO: use "github.com/pkg/errors"
import (
	"encoding/json"
	"fmt"
	bioworker "github.com/curtischong/lizzie_server/bioworker"
	typerworker "github.com/curtischong/lizzie_server/typerworker"
	"github.com/influxdata/influxdb/client/v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

func stringToDate(unparsed string) time.Time {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, unparsed)

	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(t.Format("20060102150405"))
	return t
}

// Watch
func (s server) routes(config DatabaseConfigObj) {
	s.router.HandleFunc("/watch_bio_snapshot", s.watchBioSnapshotCall())
	s.router.HandleFunc("/typer_sent_field", s.typerSentFieldCall())
	s.router.HandleFunc("/upload_emotion_evaluation", s.uploadEmotionEvaluationCall(config))
	s.router.HandleFunc("/upload_mark_event", s.uploadMarkEvent(config))
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

	pt, err := client.NewPoint("emotionEvaluation", nil, fields, stringToDate(sample.TimeEndFillingForm))
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
func uploadMarkEvent(sample bioworker.EmotionEvaluation, config DatabaseConfigObj) {
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

	pt, err := client.NewPoint("emotionEvaluation", nil, fields, stringToDate(sample.TimeEndFillingForm))
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
