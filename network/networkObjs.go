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
	EvalDatetime     string `json:"evalDatetime"`
	AccomplishedEval string `json:"accomplishedEval"`
	SocialEval       string `json:"socialEval"`
	ExhaustedEval    string `json:"exhaustedEval"`
	TiredEval        string `json:"tiredEval"`
	HappyEval        string `json:"happyEval"`
	Comments         string `json:"comments"`
	EvalLocation     string `json:"evalLocation"`
}

// EmotionEvaluationNetworkObj describes a received EmotionEvaluation from the network
type EmotionEvaluationNetworkObj struct {
	EvalDatetime string `json:"evalDatetime"`
	EvalSliders  string `json:"accomplishedEval"`
	Comments     string `json:"comments"`
	EvalLocation string `json:"evalLocation"`
}

// EmotionEvaluationSliderObj describes a received EmotionEvaluation from the network
type EmotionEvaluationSliderObj struct {
	EvalType string `json:"evalType"`
	EvalVal  string `json:"evalVal"`
}

// MarkEventObj describes a received MarkEvent
type MarkEventObj struct {
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
	Concept           string `json:"concept"`
	NewLearnings      string `json:"newLearnings"`
	OldSkills         string `json:"oldSkills"`
	PercentNew        string `json:"percentNew"`
	TimeLearned       string `json:"timeLearned"`
	TimeSpentLearning string `json:"timeSpentLearning"`
}

// ReviewObj describes a received Review
type ReviewObj struct {
	Concept        string `json:"concept"`
	DateReviewed   string `json:"dateReviewed"`
	NewLearnings   string `json:"newLearnings"`
	ReviewDuration string `json:"reviewDuration"`
	TimeLearned    string `json:"timeLearned"`
}

// ScheduledReviewObj describes a received scheduledReview
type ScheduledReviewObj struct {
	Concept           string `json:"concept"`
	TimeLearned       string `json:"timeLearned"`
	ScheduledDate     string `json:"scheduledDate"`
	ScheduledDuration string `json:"scheduledDuration"`
}
