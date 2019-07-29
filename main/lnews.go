// useful links
//https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831
//https://blog.questionable.services/article/testing-http-handlers-go/

package main

import (
	"encoding/json"
	database "github.com/curtischong/lizzie_server/database"
	network "github.com/curtischong/lizzie_server/network"
	util "github.com/curtischong/lizzie_server/util"
	"log"
	"net/http"
)

type GetCardsObj = network.GetCardsObj
type GetPanelsObj = network.GetPanelsObj
type DismissPanelObj = network.DismissPanelObj

func (s *server) getNewsCards(config ConfigObj) http.HandlerFunc {
	return func(w http.ResponseWriter, response *http.Request) {
		enableCors(w, response)

		q := response.URL.Query()
		parsedResonse := GetCardsObj{
			CardAmount: util.BetterAtoi(q.Get("cardAmount")),
			CardOffset: util.BetterAtoi(q.Get("cardOffset")),
		}

		cards, cardsSucc := database.GetCards(parsedResonse, config)
		if cardsSucc {
			cardsJsonStr, _ := json.Marshal(cards)
			w.Write([]byte(cardsJsonStr))
		} else {
			w.WriteHeader(500)
		}
	}
}

func (s *server) getNewsPanels(config ConfigObj) http.HandlerFunc {
	return func(w http.ResponseWriter, response *http.Request) {
		enableCors(w, response)

		q := response.URL.Query()
		parsedResonse := GetPanelsObj{
			PanelAmount: util.BetterAtoi(q.Get("panelAmount")),
			PanelOffset: util.BetterAtoi(q.Get("panelOffset")),
		}

		panels, panelsSucc := database.GetPanels(parsedResonse, config)
		if panelsSucc {
			panelsJsonStr, _ := json.Marshal(panels)
			w.Write([]byte(panelsJsonStr))
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
