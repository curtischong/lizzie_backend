package database

import (
	"encoding/json"
	config "github.com/curtischong/lizzie_server/config"
	network "github.com/curtischong/lizzie_server/network"
	serverutils "github.com/curtischong/lizzie_server/serverUtils"
	influx "github.com/influxdata/influxdb/client/v2"
	"log"
	"strconv"
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

type DatabaseConfigObj = config.DatabaseConfigObj

//type Client = influx.Client

// setupDB returns influxDB client
func setupDB(config DatabaseConfigObj) influx.Client {
	c, err := influx.NewHTTPClient(influx.HTTPConfig{
		Addr:     config.DBIP,
		Username: config.Username,
		Password: config.Password,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	return c
}

func setupBP(c influx.Client, config DatabaseConfigObj) influx.BatchPoints {
	bp, err := influx.NewBatchPoints(influx.BatchPointsConfig{
		Database:  config.DBName,
		Precision: "ms",
	})
	if err != nil {
		log.Fatal(err)
	}
	return bp
}

func InsertEmotionEvaluationObj(sample EmotionEvaluationObj, config DatabaseConfigObj) {
	c := setupDB(config)
	bp := setupBP(c, config)

	fields := map[string]interface{}{
		"accomplished_eval": sample.AccomplishedEval,
		"social_eval":       sample.SocialEval,
		"exhausted_eval":    sample.ExhaustedEval,
		"tired_eval":        sample.TiredEval,
		"happy_eval":        sample.HappyEval,
		"comments":          sample.Comments,
	}

	pt, err := influx.NewPoint("emotion_evaluations", nil, fields, serverutils.StringToDate(sample.EvalDatetime))
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	// write the batch
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}
	log.Printf("added emotionevaluation!")

	// close client resources
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}

func InsertMarkEventObj(sample MarkEventObj, config DatabaseConfigObj) {
	c := setupDB(config)
	bp := setupBP(c, config)

	var emotionRatings []int
	err2 := json.Unmarshal([]byte(sample.EmotionsFelt), &emotionRatings)
	if err2 != nil {
		log.Fatal(err2)
	}

	var typeBiometricsViewed []int
	err3 := json.Unmarshal([]byte(sample.TypeBiometricsViewed), &typeBiometricsViewed)
	if err3 != nil {
		log.Fatal(err3)
	}

	parsedTimeOfMark, err := strconv.ParseFloat(sample.TimeOfMark, 64)
	if err != nil {
		log.Fatal(err)
	}

	parsedIsReaction, err := strconv.Atoi(sample.IsReaction)
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
	}

	pt, err := influx.NewPoint("mark_events", nil, fields, time.Unix(0, int64(parsedTimeOfMark*1000000)))
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	// Write the batch
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}
	log.Printf("added MarkEventObj!")

	// Close client resources
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}

func InsertBioSamplesObj(sample BioSamplesObj, config DatabaseConfigObj) {
	c := setupDB(config)
	bp := setupBP(c, config)

	var dataPointNames []string
	err2 := json.Unmarshal([]byte(sample.DataPointNames), &dataPointNames)
	if err2 != nil {
		log.Fatal(err2)
	}

	var startTimes []string
	_ = json.Unmarshal([]byte(sample.StartTimes), &startTimes)

	var endTimes []string
	_ = json.Unmarshal([]byte(sample.EndTimes), &endTimes)

	//TODO: please unmarshal directly into a float
	var measurements_string []string
	_ = json.Unmarshal([]byte(sample.Measurements), &measurements_string)

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
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}
	log.Printf("added BioSamplesObj!")

	// Close client resources
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}

func InsertSkillObj(sample SkillObj, config DatabaseConfigObj) {
	c := setupDB(config)
	bp := setupBP(c, config)

	parsedPercentNew, err := strconv.ParseInt(sample.PercentNew, 10, 64)
	parsedTimeSpentLearning, err := strconv.ParseInt(sample.TimeSpentLearning, 10, 64)
	parsedTimeLearned, err := strconv.ParseFloat(sample.TimeLearned, 64)

	fields := map[string]interface{}{
		"concept":             sample.Concept,
		"new_earnings":        sample.NewLearnings,
		"old_skills":          sample.OldSkills,
		"percent_new":         parsedPercentNew,
		"time_spent_learning": parsedTimeSpentLearning,
	}

	pt, err := influx.NewPoint("learned_skills", nil, fields, time.Unix(0, int64(parsedTimeLearned*1000000)))
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	// write the batch
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}
	log.Printf("added learnedSkill!")

	// close client resources
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}

func InsertReviewObj(sample ReviewObj, config DatabaseConfigObj) {
	c := setupDB(config)
	bp := setupBP(c, config)

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
	}
	bp.AddPoint(pt)

	// write the batch
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}
	log.Printf("added SkillReview!")

	// close client resources
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}

func InsertScheduledReviewObj(sample ScheduledReviewObj, config DatabaseConfigObj) {
	c := setupDB(config)
	bp := setupBP(c, config)

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
	}
	bp.AddPoint(pt)

	// write the batch
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}
	log.Printf("added ScheduledSkillReview!")

	// close client resources
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}

func InsertTyperObj(sample TyperObj, config DatabaseConfigObj) {
	c := setupDB(config)
	bp := setupBP(c, config)

	fields := map[string]interface{}{
		"url":          sample.Url,
		"text":         sample.Text,
		"deleted_text": sample.DeletedText,
	}

	pt, err := influx.NewPoint("typer_text", nil, fields, time.Unix(0, int64(sample.TimeSent*1000000)))
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	// write the batch
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}
	log.Printf("added TyperObj!")

	// close client resources
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}

func InsertMessengerObj(sample MessengerObj, config DatabaseConfigObj) {
	c := setupDB(config)
	bp := setupBP(c, config)

	fields := map[string]interface{}{
		"fbid":         sample.FBID,
		"message":      sample.Message,
		"deleted_text": sample.DeletedText,
	}

	pt, err := influx.NewPoint("typer_text", nil, fields, time.Unix(0, int64(sample.TimeSent*1000000)))
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	// write the batch
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}
	log.Printf("added TyperObj!")

	// close client resources
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}
