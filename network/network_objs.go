package network

import (
	"time"
)

// Typer

// TyperObj describes text I sent
type TyperObj struct {
	Unixt       int64     `json:unixt`
	Ts          time.Time `json:ts`
	DeletedText bool      `json:deletedText`
	Url         string    `json:"url"`
	SentText    string    `json:"sentText"`
}

// MessengerObj describes the messages I sent
type MessengerObj struct {
	Unixt       int64     `json:unixt`
	Ts          time.Time `json:ts`
	DeletedText bool      `json:deletedText`
	FBID        string    `json:"FBID"`
	Message     string    `json:"message"`
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
// Time learned is when you press new learning on the app
// it's not time started learning
// remember that it can take multiple days to learn so
// that's why we have timeSpentLearning not timeStartLearning
type SkillObj struct {
	TimeLearned       string `json:"timeLearned"`       // represent to 2 decimal places (smallint) (0-10,000)
	TimeSpentLearning string `json:"timeSpentLearning"` // represent in seconds (int)
	Concept           string `json:"concept"`
	NewLearnings      string `json:"newLearnings"`
	OldSkills         string `json:"oldSkills"`
	PercentNew        string `json:"percentNew"`
}

// ReviewObj describes a received Review
type ReviewObj struct {
	TimeLearned        string `json:"timeLearned"`
	TimeReviewed       string `json:"timeReviewed"`
	Concept            string `json:"concept"`
	NewLearnings       string `json:"newLearnings"`
	TimeSpentReviewing string `json:"timeSpentReviewing"` // represent in seconds (int)
}

// ScheduledReviewObj describes a received scheduledReview
type ScheduledReviewObj struct {
	TimeLearned       string `json:"timeLearned"`
	TimeScheduled     string `json:"timeScheduled"`
	Concept           string `json:"concept"`
	ScheduledDuration string `json:"scheduleDuration"`
}
