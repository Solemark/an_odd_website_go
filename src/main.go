package main

import (
	"an_odd_website/src/router"
	"an_odd_website/src/router/client"
	"an_odd_website/src/router/employee"
	"an_odd_website/src/router/setting"
	"log"
	"net/http"
)

func main() {
	mapRoutes()
	log.Println("Starting server at localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func mapRoutes() {
	mapPageRoutes()
	mapClientRoutes()
	mapEmployeeRoutes()
	//mapAccountingRoutes()
	mapSettingRoutes()
}

func mapPageRoutes() {
	http.HandleFunc("/", router.GetWebPage)
	http.HandleFunc("/styles/{style}", router.GetStyles)
	http.HandleFunc("/scripts/{script}", router.GetScripts)
	http.HandleFunc("/favicon.ico", router.GetIcon)
}

func mapClientRoutes() {
	http.HandleFunc("/data/clients", client.GetClientData)
	http.HandleFunc("/data/clients/new", client.NewClient)
	http.HandleFunc("/data/clients/update", client.UpdateClient)
	http.HandleFunc("/data/clients/remove", client.RemoveClient)
}

func mapEmployeeRoutes() {
	http.HandleFunc("/data/employees", employee.GetEmployeeData)
	http.HandleFunc("/data/employees/new", employee.NewEmployee)
	http.HandleFunc("/data/employees/update", employee.UpdateEmployee)
	http.HandleFunc("/data/employees/remove", employee.RemoveEmployee)
}

//func mapAccountingRoutes() {}

func mapSettingRoutes() {
	http.HandleFunc("/data/settings", setting.GetSettingData)
	http.HandleFunc("/data/settings/update", setting.UpdateSetting)
}
