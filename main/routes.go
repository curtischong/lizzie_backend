package main

// TODO: think about standardizing these endpoint names
func (s server) routes(config ConfigObj) {
	s.router.HandleFunc("/get_news_cards", s.getNewsCards(config))
	s.router.HandleFunc("/get_news_panels", s.getNewsPanels(config))
	s.router.HandleFunc("/get_peaks_skills", s.getPeaksSkills(config))
	s.router.HandleFunc("/delete_peaks_skill", s.deletePeaksSkills(config))
	s.router.HandleFunc("/dismiss_panel", s.dismissPanel(config))
	s.router.HandleFunc("/typer_sent_field", s.typerSentFieldCall(config))
	s.router.HandleFunc("/messenger_sent_text", s.messengerSentFieldCall(config))
	s.router.HandleFunc("/upload_bio_samples", s.uploadBioSamplesCall(config))
	s.router.HandleFunc("/upload_emotion_evaluation", s.uploadEmotionEvaluationCall(config))
	s.router.HandleFunc("/upload_mark_event", s.uploadMarkEventCall(config))
	s.router.HandleFunc("/upload_skill", s.uploadSkillCall(config))
	s.router.HandleFunc("/upload_review", s.uploadReviewCall(config))
	s.router.HandleFunc("/upload_scheduled_review", s.uploadScheduledReviewCall(config))
	//s.router.HandleFunc("/admin", s.adminOnly(s.handleAdminIndex()))
}
