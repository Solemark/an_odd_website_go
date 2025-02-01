package main

import (
	"an_odd_website/router"
	"an_odd_website/router/client"
	"an_odd_website/router/employee"
	"an_odd_website/router/setting"
	"log"
	"net/http"
)

func main() {
	setRoutes()
	log.Println("Starting server at localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func setRoutes() {
	http.HandleFunc("/", router.GetWebPage)
	http.HandleFunc("/styles/{style}", router.GetStyles)
	http.HandleFunc("/scripts/{script}", router.GetScripts)
	http.HandleFunc("/favicon.ico", router.GetIcon)
	setClientRoutes()
	setEmployeeRoutes()
	//setAccountingRoutes()
	setSettingRoutes()
}

func setClientRoutes() {
	http.HandleFunc("/data/clients", client.GetClientData)
	http.HandleFunc("/data/clients/new", client.NewClient)
	http.HandleFunc("/data/clients/update", client.UpdateClient)
	http.HandleFunc("/data/clients/remove", client.RemoveClient)
}

func setEmployeeRoutes() {
	http.HandleFunc("/data/employees", employee.GetEmployeeData)
	http.HandleFunc("/data/employees/new", employee.NewEmployee)
	http.HandleFunc("/data/employees/update", employee.UpdateEmployee)
	http.HandleFunc("/data/employees/remove", employee.RemoveEmployee)
}

//func setAccountingRoutes() {}

func setSettingRoutes() {
	http.HandleFunc("/data/settings", setting.GetSettingData)
	http.HandleFunc("/data/settings/update", setting.UpdateSetting)
}
