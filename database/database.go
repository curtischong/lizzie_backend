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
		"timeStartFillingForm": sample.TimeStartFillingForm,
		"normalEval":           sample.NormalEval,
		"socialEval":           sample.SocialEval,
		"exhaustedEval":        sample.ExhaustedEval,
		"tiredEval":            sample.TiredEval,
		"happyEval":            sample.HappyEval,
		"comments":             sample.Comments,
	}

	pt, err := influx.NewPoint("emotionEvaluations", nil, fields, serverutils.StringToDate(sample.TimeEndFillingForm))
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

	parsedTimeStartFillingForm, err := strconv.ParseFloat(sample.TimeStartFillingForm, 64)
	if err != nil {
		log.Fatal(err)
	}

	parsedTimeEndFillingForm, err := strconv.ParseFloat(sample.TimeEndFillingForm, 64)
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
		"timeStartFillingForm": parsedTimeStartFillingForm,
		"timeEndFillingForm":   parsedTimeEndFillingForm,
		"isReaction":           parsedIsReaction,
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

	if parsedIsReaction == 0 {
		fields["anticipationStart"] = parsedAnticipationStart
	}

	pt, err := influx.NewPoint("MarkEventObjs", nil, fields, time.Unix(0, int64(parsedTimeOfMark*1000000)))
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
			"dataPointName": dataPointNames[i],
			"startTime":     parsedStartTime,
			"measurement":   measurement,
		}

		pt, err := influx.NewPoint("BioSamplesObj", nil, fields, time.Unix(0, int64(parsedEndTime*1000000)))
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
		"concept":           sample.Concept,
		"newLearnings":      sample.NewLearnings,
		"oldSkills":         sample.OldSkills,
		"percentNew":        parsedPercentNew,
		"timeSpentLearning": parsedTimeSpentLearning,
	}

	pt, err := influx.NewPoint("learnedSkills", nil, fields, time.Unix(0, int64(parsedTimeLearned*1000000)))
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
		"concept":        sample.Concept,
		"dateReviewed":   parsedDateReviewed,
		"newLearnings":   sample.NewLearnings,
		"reviewDuration": parsedReviewDuration,
	}

	pt, err := influx.NewPoint("skillReviews", nil, fields, time.Unix(0, int64(parsedTimeLearned*1000000)))
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
		"concept":           sample.Concept,
		"timeLearned":       parsedTimeLearned,
		"scheduledDate":     parsedScheduledDate,
		"scheduledDuration": parsedScheduledDuration,
	}

	pt, err := influx.NewPoint("scheduledReviews", nil, fields, time.Unix(0, int64(parsedTimeLearned*1000000)))
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
