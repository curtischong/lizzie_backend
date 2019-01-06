package typerworker

type BioSnapshot struct {
	TimeStart string `json:"timestart"`
	TimeEnd   string `json:"timeend"`
	Heartrate string `json:"heartrate"`
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
	LeftCrop             string `json:"leftCrop"`
	RightCrop            string `json:"rightCrop"`
	TimeOfEvent          string `json:"TimeOfEvent"`
	EmotionsFelt         string `json:"happyEval"`
	Comments             string `json:"comments"`
	TypeBiometricsViewed string `json:"typeBiometricsViewed"`
}
