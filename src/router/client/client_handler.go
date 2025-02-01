package client

import (
	"an_odd_website/src/router"
	"an_odd_website/src/router/common"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetClientData(w http.ResponseWriter, r *http.Request) {
	clients := []client{}
	for _, item := range getClientList() {
		if item.Visible {
			clients = append(clients, item)
		}
	}
	msg, e := json.Marshal(clients)
	router.CheckAndLogError(e)
	common.SendContent(w, common.JSON_RESPONSE, msg)
}

func NewClient(w http.ResponseWriter, r *http.Request) {
	clients := getClientList()
	_, fn, ln, ea := parseForm(r)
	clients = append(clients, client{
		Id:           len(clients),
		FirstName:    fn,
		LastName:     ln,
		EmailAddress: ea,
		Visible:      true,
	})
	writeClientList(clients)
	common.SendRedirect(w, r, "/clients")
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	clients := getClientList()
	id, fn, ln, ea := parseForm(r)
	clients[id] = client{Id: id, FirstName: fn, LastName: ln, EmailAddress: ea, Visible: true}
	writeClientList(clients)
	common.SendRedirect(w, r, "/clients")
}

func RemoveClient(w http.ResponseWriter, r *http.Request) {
	var clients []client = getClientList()
	id, _, _, _ := parseForm(r)
	clients[id].Visible = false
	writeClientList(clients)
	common.SendRedirect(w, r, "/clients")
}

func parseForm(r *http.Request) (int, string, string, string) {
	r.ParseForm()
	buffer := [4]string{"0"}
	for key, value := range r.Form {
		switch key {
		case "Id":
			buffer[0] = value[0]
		case "first_name":
			buffer[1] = value[0]
		case "last_name":
			buffer[2] = value[0]
		case "email_address":
			buffer[3] = value[0]
		}
	}
	id, e := strconv.Atoi(buffer[0])
	router.CheckAndLogError(e)
	return id, buffer[1], buffer[2], buffer[3]
}
