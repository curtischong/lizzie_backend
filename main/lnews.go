// useful links
//https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831
//https://blog.questionable.services/article/testing-http-handlers-go/

package main

import (
	"encoding/json"
	database "github.com/curtischong/lizzie_server/database"
	network "github.com/curtischong/lizzie_server/network"
	utils "github.com/curtischong/lizzie_server/util"
	"net/http"
)

type GetCardsAndPanelsObj = network.GetCardsAndPanelsObj

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
