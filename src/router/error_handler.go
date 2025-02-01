package router

import (
	"an_odd_website/src/router/common"
	"encoding/json"
	"log"
	"net/http"
)

func errorHandler(w http.ResponseWriter, e error) {
	if e != nil {
		o, e := json.Marshal(e)
		CheckAndLogError(e)
		common.SendContent(w, common.JSON_RESPONSE, o)
	}
}

func CheckAndLogError(e error) {
	if e != nil {
		log.Println(e)
	}
}
