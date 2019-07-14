/*
Handles all DB communications

Great resource to setup psql: https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/
Great resource for bulk inserts: https://stackoverflow.com/questions/21108084/how-to-insert-multiple-data-at-once
Using Go's sql library: https://www.calhoun.io/updating-and-deleting-postgresql-records-using-gos-sql-package/
Better Bulk Inserts: github.com/jmoiron/sqlx, https://github.com/jmoiron/sqlx/pull/285
seltup: https://www.calhoun.io/inserting-records-into-a-postgresql-database-with-gos-database-sql-package/
guide: https://jbrandhorst.com/post/postgres/

*/

package database

import (
	"database/sql"
	//"encoding/json"
	"fmt"
	config "github.com/curtischong/lizzie_server/config"
	network "github.com/curtischong/lizzie_server/network"
	//utils "github.com/curtischong/lizzie_server/serverUtils"
	//influx "github.com/influxdata/influxdb/client/v2"
	//"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	//"strconv"
	"time"
)

type EmotionEvaluationObj = network.EmotionEvaluationObj
type BioSamplesObj = network.BioSamplesObj
type MarkEventObj = network.MarkEventObj
type SkillObj = network.SkillObj
type ReviewObj = network.ReviewObj
type ScheduledReviewObj = network.ScheduledReviewObj
type TyperObj = network.TyperObj
type MessengerObj = network.MessengerObj

type DBConfigObj = config.DBConfigObj
type ConfigObj = config.ConfigObj
type Time = time.Time

//type Client = influx.Client

func SetupDB(config ConfigObj) DBObj {
	db := DBObj{
		DBClient: nil,
	}
	return db
}

func SetupDBConfig(config *ConfigObj) {
	var dbip string
	var dbport string
	var dbConfig = config.DBConfig
	if config.ServerConfig.IsDev {
		dbip = dbConfig.DevDBIP
		dbport = dbConfig.DevDBPPort
	} else {
		dbport = dbConfig.ProdDBPort
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbip, dbport, dbConfig.Username, dbConfig.Password, dbConfig.DBName)
	config.DBConfig.DBConfigStr = psqlInfo
}

func connectDB(config ConfigObj) *sql.DB {
	//TODO: have multiple retries if failed
	db, err := sql.Open("postgres", config.DBConfig.DBConfigStr)
	if err != nil {
		panic(err)
	}

	// Ping forces connection to be made
	err = db.Ping()
	if err != nil {
		log.Println(err)
		panic(err)
	}

	fmt.Printf("Successfully connected to DB with configuration %v\n", config.DBConfig.DBConfigStr)
	return db
}

func GetCards(config ConfigObj) ([]map[string]string, bool) {
	db := connectDB(config)
	defer db.Close()

	var cards = make([]map[string]string, 0)

	rows, err := db.Query(`select unixt, card from lnews.card ORDER BY unixt DESC LIMIT 3`)
	if err != nil {
		log.Fatal(err)
		return cards, false
	}
	defer rows.Close()
	for rows.Next() {
		var unixt string
		var card string
		err = rows.Scan(&unixt, &card)
		if err != nil {
			log.Fatal(err)
			return cards, false
		}
		//cardJson := map[string]string{"unixt": strconv.Itoa(unixt), "card": card}
		cardJson := map[string]string{"unixt": unixt, "card": card}
		cards = append(cards, cardJson)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return cards, false
	}
	//fmt.Println(string(cardsJson))

	return cards, true
}

func GetPanels(config ConfigObj) ([]map[string]string, bool) {
	db := connectDB(config)
	defer db.Close()

	var panels = make([]map[string]string, 0)

	rows, err := db.Query(`select unixt, dismissed, panel from lnews.panel ORDER BY unixt DESC LIMIT 3`)
	if err != nil {
		log.Fatal(err)
		return panels, false
	}
	defer rows.Close()
	for rows.Next() {
		var unixt string
		var dismissed string
		var panel string
		err = rows.Scan(&unixt, &dismissed, &panel)
		if err != nil {
			log.Fatal(err)
			return panels, false
		}
		//cardJson := map[string]string{"unixt": strconv.Itoa(unixt), "card": card}
		panelJson := map[string]string{"unixt": unixt, "dismissed": dismissed, "card": panel}
		panels = append(panels, panelJson)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return panels, false
	}
	//fmt.Println(string(cardsJson))

	return panels, true
}

func InsertEmotionEvaluationObj(sample EmotionEvaluationObj, config ConfigObj) bool {
	db := connectDB(config)
	defer db.Close()

	/*fields := map[string]interface{}{
		"accomplished_eval": sample.AccomplishedEval,
		"social_eval":       sample.SocialEval,
		"exhausted_eval":    sample.ExhaustedEval,
		"tired_eval":        sample.TiredEval,
		"happy_eval":        sample.HappyEval,
		"comments":          sample.Comments,
	}

	/*

		pt, err := influx.NewPoint("emotion_evaluations", nil, fields, time.Unix(0, int64(sample.EvalDatetime*1000000)))
		if err != nil {
			log.Fatal(err)
			return false
		}
		bp.AddPoint(pt)

		// write the batch
		if err := db.DBClient.Write(bp); err != nil {
			log.Fatal(err)
			return false
		}
		log.Printf("added emotionevaluation!")*/
	return true
}

func InsertMarkEventObj(sample MarkEventObj, config ConfigObj) bool {
	db := connectDB(config)
	defer db.Close()

	/*var emotionRatings []int
	err2 := json.Unmarshal([]byte(sample.EmotionsFelt), &emotionRatings)
	if err2 != nil {
		log.Fatal(err2)
		return false
	}

	var typeBiometricsViewed []int
	err3 := json.Unmarshal([]byte(sample.TypeBiometricsViewed), &typeBiometricsViewed)
	if err3 != nil {
		log.Fatal(err3)
		return false
	}

	parsedTimeOfMark, err := strconv.ParseFloat(sample.TimeOfMark, 64)
	if err != nil {
		log.Fatal(err)
		return false
	}

	parsedIsReaction, err := strconv.Atoi(sample.IsReaction)
	if err != nil {
		log.Fatal(err)
		return false
	}

	parsedAnticipationStart, err := strconv.ParseFloat(sample.AnticipationStart, 64)
	if err != nil {
		log.Fatal(err)
		return false
	}
	parsedTimeOfEvent, err := strconv.ParseFloat(sample.TimeOfEvent, 64)
	if err != nil {
		log.Fatal(err)
		return false
	}
	parsedReactionEnd, err := strconv.ParseFloat(sample.ReactionEnd, 64)
	if err != nil {
		log.Fatal(err)
		return false
	}

	fields := map[string]interface{}{
		"is_reaction":            parsedIsReaction,
		"time_of_event":          parsedTimeOfEvent,
		"reaction_end":           parsedReactionEnd,
		"emotions_felt_fear":     emotionRatings[0],
		"emotions_felt_joy":      emotionRatings[1],
		"emotions_felt_anger":    emotionRatings[2],
		"emotions_felt_sad":      emotionRatings[3],
		"emotions_felt_disgust":  emotionRatings[4],
		"emotions_felt_suprise":  emotionRatings[5],
		"emotions_felt_contempt": emotionRatings[6],
		"emotions_felt_interest": emotionRatings[7],
		"comments":               sample.Comments,
		"biometrics_viewed_hr":   typeBiometricsViewed[0],
	}

	if parsedIsReaction == 0 {
		fields["anticipationStart"] = parsedAnticipationStart
	}*/

	/*
		pt, err := influx.NewPoint("mark_events", nil, fields, time.Unix(0, int64(parsedTimeOfMark*1000000)))
		if err != nil {
			log.Fatal(err)
			return false
		}
		bp.AddPoint(pt)

		// Write the batch
		if err := db.DBClient.Write(bp); err != nil {
			log.Fatal(err)
			return false
		}
		log.Printf("added MarkEventObj!")*/
	return true
}

func InsertBioSamplesObj(sample BioSamplesObj, config ConfigObj) bool {
	db := connectDB(config)
	defer db.Close()

	/*var dataPointNames []string
	err2 := json.Unmarshal([]byte(sample.DataPointNames), &dataPointNames)
	if err2 != nil {
		log.Fatal(err2)
		return false
	}

	var startTimes []string
	_ = json.Unmarshal([]byte(sample.StartTimes), &startTimes)

	var endTimes []string
	_ = json.Unmarshal([]byte(sample.EndTimes), &endTimes)

	//TODO: please unmarshal directly into a float
	var measurements_string []string
	_ = json.Unmarshal([]byte(sample.Measurements), &measurements_string)*/

	/*
		for i := 0; i < len(dataPointNames); i++ {

			measurement, err := strconv.ParseFloat(measurements_string[i], 64)

			parsedStartTime, err := strconv.ParseFloat(startTimes[i], 64)
			parsedEndTime, err := strconv.ParseFloat(endTimes[i], 64)

			fields := map[string]interface{}{
				"data_point_name": dataPointNames[i],
				"start_time":      parsedStartTime,
				"measurement":     measurement,
			}

			pt, err := influx.NewPoint("bio_samples", nil, fields, time.Unix(0, int64(parsedEndTime*1000000)))
			if err != nil {
				log.Fatal(err)
			}
			bp.AddPoint(pt)
		}

		// Write the batch
		if err := db.DBClient.Write(bp); err != nil {
			log.Fatal(err)
			return false
		}
		log.Printf("added BioSamplesObj!")*/
	return true
}

func InsertSkillObj(sample SkillObj, config ConfigObj) bool {
	db := connectDB(config)
	defer db.Close()

	/*parsedPercentNew, err := strconv.ParseInt(sample.PercentNew, 10, 64)
	parsedTimeSpentLearning, err := strconv.ParseInt(sample.TimeSpentLearning, 10, 64)
	parsedTimeLearned, err := strconv.ParseFloat(sample.TimeLearned, 64)

	fields := map[string]interface{}{
		"concept":             sample.Concept,
		"new_earnings":        sample.NewLearnings,
		"old_skills":          sample.OldSkills,
		"percent_new":         parsedPercentNew,
		"time_spent_learning": parsedTimeSpentLearning,
	}
	/*

		pt, err := influx.NewPoint("learned_skills", nil, fields, time.Unix(0, int64(parsedTimeLearned*1000000)))
		if err != nil {
			log.Fatal(err)
			return false
		}
		bp.AddPoint(pt)

		// write the batch
		if err := db.DBClient.Write(bp); err != nil {
			log.Fatal(err)
			return false
		}
		log.Printf("added learnedSkill!")*/
	return true
}

func InsertReviewObj(sample ReviewObj, config ConfigObj) bool {
	/*
		db := connectDB(config)
		defer db.Close()
		bp := setupBP(config)

		parsedDateReviewed, err := strconv.ParseFloat(sample.DateReviewed, 64)
		parsedReviewDuration, err := strconv.ParseInt(sample.ReviewDuration, 10, 64)
		parsedTimeLearned, err := strconv.ParseFloat(sample.TimeLearned, 64)

		fields := map[string]interface{}{
			"concept":         sample.Concept,
			"date_reviewed":   parsedDateReviewed,
			"new_learnings":   sample.NewLearnings,
			"review_duration": parsedReviewDuration,
		}

		pt, err := influx.NewPoint("skill_reviews", nil, fields, time.Unix(0, int64(parsedTimeLearned*1000000)))
		if err != nil {
			log.Fatal(err)
			return false
		}
		bp.AddPoint(pt)

		// write the batch
		if err := db.DBClient.Write(bp); err != nil {
			log.Fatal(err)
			return false
		}
		log.Printf("added SkillReview!")*/
	return true
}

func InsertScheduledReviewObj(sample ScheduledReviewObj, config ConfigObj) bool {
	/*
		db := connectDB(config)
		defer db.Close()
		bp := setupBP(config)

		parsedScheduledDate, err := strconv.ParseFloat(sample.ScheduledDate, 64)
		parsedScheduledDuration, err := strconv.ParseInt(sample.ScheduledDuration, 10, 64)
		parsedTimeLearned, err := strconv.ParseFloat(sample.TimeLearned, 64)

		fields := map[string]interface{}{
			"concept":            sample.Concept,
			"time_learned":       parsedTimeLearned,
			"scheduled_date":     parsedScheduledDate,
			"scheduled_duration": parsedScheduledDuration,
		}

		pt, err := influx.NewPoint("scheduled_reviews", nil, fields, time.Unix(0, int64(parsedTimeLearned*1000000)))
		if err != nil {
			log.Fatal(err)
			return false
		}
		bp.AddPoint(pt)

		// write the batch
		if err := db.DBClient.Write(bp); err != nil {
			log.Fatal(err)
			return false
		}
		log.Printf("added ScheduledSkillReview!")*/
	return true
}

func InsertTyperObj(sample TyperObj, config ConfigObj) bool {
	db := connectDB(config)
	defer db.Close()
	// NOTE: deleted_text should probably be a tag but tags cannot be bools

	/*fields := map[string]interface{}{
		"url":          sample.Url,
		"text":         sample.Text,
		"deleted_text": sample.DeletedText,
	}

	/*
		pt, err := influx.NewPoint("typer_text", nil, fields, time.Unix(0, int64(sample.TimeSent*1000000)))
		if err != nil {
			log.Fatal(err)
			return false
		}
		bp.AddPoint(pt)

		// write the batch
		if err := db.DBClient.Write(bp); err != nil {
			log.Fatal(err)
			return false
		}
		log.Printf("added TyperObj!")*/
	return true
}

//TODO: maybe have a tag that says if I sent somebody a photo or a file bc those messages don't have text
func InsertMessengerObj(sample MessengerObj, config ConfigObj) bool {
	db := connectDB(config)
	defer db.Close()

	/*fields := map[string]interface{}{
		"fbid":         sample.FBID,
		"message":      sample.Message,
		"deleted_text": sample.DeletedText,
	}

	/*
		pt, err := influx.NewPoint("messenger_text", nil, fields, time.Unix(0, int64(sample.TimeSent*1000000)))
		if err != nil {
			log.Fatal(err)
			return false
		}
		bp.AddPoint(pt)

		// write the batch
		if err := db.DBClient.Write(bp); err != nil {
			log.Fatal(err)
			return false
		}
		log.Printf("added TyperObj!")*/
	return true
}
