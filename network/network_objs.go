package network

// Typer

// TyperObj describes text I sent
type TyperObj struct {
	Url         string `json:"url"`
	Text        string `json:"text"`
	DeletedText bool   `json:deletedText`
	TimeSent    uint64 `json:"timeSent"` // int32 should fit but just in case
}

// MessengerObj describes the messages I sent
type MessengerObj struct {
	FBID        string `json:"FBID"`
	Message     string `json:"message"`
	DeletedText bool   `json:deletedText`
	TimeSent    uint64 `json:"timeSent"` // int32 should fit but just in case
}

// LNews

type GetCardsAndPanelsObj struct {
	CardAmount  int `json:"cardAmount"`
	CardOffset  int `json:"cardOffset"`
	PanelAmount int `json:"panelAmount"`
	PanelOffset int `json:"panelOffset"`
}

// BioSamples

// BioSamplesObj describes a received BioSample
type BioSamplesObj struct {
	DataPointNames string `json:"dataPointNames"`
	StartTimes     string `json:"startTimes"`
	EndTimes       string `json:"endTimes"`
	Measurements   string `json:"measurements"`
}

// Emotions

// EmotionEvaluationObj describes a received EmotionEvaluation
type EmotionEvaluationObj struct {
	EvalDatetime     uint64 `json:"evalDatetime"`
	AccomplishedEval int    `json:"accomplishedEval"`
	SocialEval       int    `json:"socialEval"`
	ExhaustedEval    int    `json:"exhaustedEval"`
	TiredEval        int    `json:"tiredEval"`
	HappyEval        int    `json:"happyEval"`
	Comments         string `json:"comments"`
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
	EvalTime string `json:evalTime` // TODO: update API on clients to send this
}

// Events

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

// Lizzie Peaks

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
