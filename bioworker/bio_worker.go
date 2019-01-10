package typerworker

type BioSamples struct {
	DataPointNames string `json:"dataPointNames"`
	StartTimes     string `json:"startTimes"`
	EndTimes       string `json:"endTimes"`
	Measurements   string `json:"measurements"`
}

type EmotionEvaluation struct {
	TimeStartFillingForm string `json:"timeStartFillingForm"`
	TimeEndFillingForm   string `json:"timeEndFillingForm"`
	NormalEval           string `json:"normalEval"`
	SocialEval           string `json:"socialEval"`
	ExhaustedEval        string `json:"exhaustedEval"`
	TiredEval            string `json:"tiredEval"`
	HappyEval            string `json:"happyEval"`
	Comments             string `json:"comments"`
}

type MarkEvent struct {
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
