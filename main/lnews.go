// useful links
//https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831
//https://blog.questionable.services/article/testing-http-handlers-go/

package main

import (
	"encoding/json"
	database "github.com/curtischong/lizzie_server/database"
	network "github.com/curtischong/lizzie_server/network"
	utils "github.com/curtischong/lizzie_server/util"
	"log"
	"net/http"
)

type GetCardsAndPanelsObj = network.GetCardsAndPanelsObj
type DismissPanelObj = network.DismissPanelObj

func (s *server) getCardsAndPanelsCall(config ConfigObj) http.HandlerFunc {
	return func(w http.ResponseWriter, response *http.Request) {
		enableCors(w, response)

		q := response.URL.Query()
		parsedResonse := GetCardsAndPanelsObj{
			CardAmount:  utils.BetterAtoi(q.Get("cardAmount")),
			CardOffset:  utils.BetterAtoi(q.Get("cardOffset")),
			PanelAmount: utils.BetterAtoi(q.Get("panelAmount")),
			PanelOffset: utils.BetterAtoi(q.Get("panelOffset")),
		}

		cards, cardsSucc := database.GetCards(parsedResonse, config)
		panels, panelsSucc := database.GetPanels(parsedResonse, config)
		//log.Println(panelsSucc)

		if cardsSucc && panelsSucc {
			cardsAndPanelsObj := map[string][]map[string]string{"cards": cards, "panels": panels}
			cardsAndPanelsJsonStr, _ := json.Marshal(cardsAndPanelsObj)

			w.Write([]byte(cardsAndPanelsJsonStr))
		} else {
			w.WriteHeader(500)
		}
	}
}

func (s *server) dismissPanel(config ConfigObj) http.HandlerFunc {
	return func(w http.ResponseWriter, response *http.Request) {
		enableCors(w, response)

		body := getResponseBody(w, response)
		parsedResonse := DismissPanelObj{}
		jsonErr := json.Unmarshal(body, &parsedResonse)
		if jsonErr != nil {
			log.Println(body)
			log.Println("couldn't parse body")
			log.Println(jsonErr)
			w.WriteHeader(500)
			return
		}

		succ := database.DismissPanel(parsedResonse.Unixt, config)

		if !succ {
			w.WriteHeader(500)
		}
	}
}

func (s *server) getPeaksSkills(config ConfigObj) http.HandlerFunc {
	return func(w http.ResponseWriter, response *http.Request) {
		enableCors(w, response)

		// TODO: finish this
		q := response.URL.Query()
		parsedResonse := GetCardsAndPanelsObj{
			CardAmount:  utils.BetterAtoi(q.Get("cardAmount")),
			CardOffset:  utils.BetterAtoi(q.Get("cardOffset")),
			PanelAmount: utils.BetterAtoi(q.Get("panelAmount")),
			PanelOffset: utils.BetterAtoi(q.Get("panelOffset")),
		}

		cards, cardsSucc := database.GetCards(parsedResonse, config)
		panels, panelsSucc := database.GetPanels(parsedResonse, config)
		//log.Println(panelsSucc)

		if cardsSucc && panelsSucc {
			cardsAndPanelsObj := map[string][]map[string]string{"cards": cards, "panels": panels}
			cardsAndPanelsJsonStr, _ := json.Marshal(cardsAndPanelsObj)

			w.Write([]byte(cardsAndPanelsJsonStr))
		} else {
			w.WriteHeader(500)
		}
	}
}
