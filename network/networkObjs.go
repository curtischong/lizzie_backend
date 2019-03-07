package network

// BioSamplesObj describes a received BioSample
type BioSamplesObj struct {
	DataPointNames string `json:"dataPointNames"`
	StartTimes     string `json:"startTimes"`
	EndTimes       string `json:"endTimes"`
	Measurements   string `json:"measurements"`
}

// EmotionEvaluationObj describes a received EmotionEvaluation
type EmotionEvaluationObj struct {
	TimeStartFillingForm string `json:"timeStartFillingForm"`
	TimeEndFillingForm   string `json:"timeEndFillingForm"`
	NormalEval           string `json:"normalEval"`
	SocialEval           string `json:"socialEval"`
	ExhaustedEval        string `json:"exhaustedEval"`
	TiredEval            string `json:"tiredEval"`
	HappyEval            string `json:"happyEval"`
	Comments             string `json:"comments"`
}

// MarkEventObj describes a received MarkEvent
type MarkEventObj struct {
	TimeStartFillingForm string `json:"timeStartFillingForm"`
	TimeEndFillingForm   string `json:"timeEndFillingForm"`
	TimeOfMark           string `json:"timeOfMark"`
	IsReaction           string `json:"isReaction"`
	AnticipationStart    string `json:"anticipationStart"`
	TimeOfEvent          string `json:"timeOfEvent"`
	ReactionEnd          string `json:"reactionEnd"`
	EmotionsFelt         string `json:"emotionsFelt"`
	Comments             string `json:"comments"`
	TypeBiometricsViewed string `json:"typeBiometricsViewed"`
}

// SkillObj describes a received Skill
type SkillObj struct {
	Concept                  string `json:"concept"`
	NewLearnings             string `json:"newLearnings"`
	OldSkills                string `json:"oldSkills"`
	PercentNew               string `json:"percentNew"`
	TimeLearned              string `json:"timeLearned"`
	TimeSpentLearning        string `json:"timeSpentLearning"`
	ScheduledReviews         string `json:"scheduledReviews"`
	ScheduledReviewDurations string `json:"scheduledReviewDurations"`
	Reviews                  string `json:"reviews"`
	ReviewDurations          string `json:"reviewDurations"`
}

// ReviewObj describes a received Review
type ReviewObj struct {
	Concept        string `json:"concept"`
	DateReviewed   string `json:"dateReviewed"`
	NewLearnings   string `json:"newLearnings"`
	ReviewDuration string `json:"reviewDuration"`
	TimeLearned    string `json:"timeLearned"`
}
