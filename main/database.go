package main

type DatabaseConfigObj struct {
	Dbname   string `json:"dbname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func insertEmotionEvaluation(sample EmotionEvaluation, config database.DatabaseConfigObj) {
	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://10.8.0.2:8086",
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
	log.Printf("added emotionEvaluation!")

	// Close client resources
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}

//TODO: udpate the type to MarkEvent
func insertMarkEvent(sample MarkEvent, config database.DatabaseConfigObj) {
	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://10.8.0.2:8086",
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

	pt, err := client.NewPoint("markEvents", nil, fields, time.Unix(0, int64(parsedTimeOfMark*1000000)))
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	// Write the batch
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}
	log.Printf("added MarkEvent!")

	// Close client resources
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}

func insertBioSamples(sample BioSamples, config database.DatabaseConfigObj) {
	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://10.8.0.2:8086",
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

	var dataPointNames []string
	err2 := json.Unmarshal([]byte(sample.DataPointNames), &dataPointNames)
	if err != nil {
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

		pt, err := client.NewPoint("bioSamples", nil, fields, time.Unix(0, int64(parsedEndTime*1000000)))
		if err != nil {
			log.Fatal(err)
		}
		bp.AddPoint(pt)
	}

	// Write the batch
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}
	log.Printf("added BioSamples!")

	// Close client resources
	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
}
