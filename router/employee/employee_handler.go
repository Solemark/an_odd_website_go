package employee

import (
	"an_odd_website/router"
	"an_odd_website/router/common"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetEmployeeData(w http.ResponseWriter, r *http.Request) {
	employees := []employee{}
	for _, item := range getEmployeeList() {
		if item.Visible {
			employees = append(employees, item)
		}
	}
	msg, e := json.Marshal(employees)
	router.CheckAndLogError(e)
	common.SendContent(w, common.JSON_RESPONSE, msg)
}

func NewEmployee(w http.ResponseWriter, r *http.Request) {
	employees := getEmployeeList()
	_, fn, ln, ea, rl := parseForm(r)
	employees = append(employees, employee{
		Id:           len(employees),
		FirstName:    fn,
		LastName:     ln,
		EmailAddress: ea,
		Role:         rl,
		Visible:      true,
	})
	writeEmployeeList(employees)
	common.SendRedirect(w, r, "/employees")
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	employees := getEmployeeList()
	id, fn, ln, ea, rl := parseForm(r)
	employees[id] = employee{Id: id, FirstName: fn, LastName: ln, EmailAddress: ea, Role: rl, Visible: true}
	writeEmployeeList(employees)
	common.SendRedirect(w, r, "/employees")
}

func RemoveEmployee(w http.ResponseWriter, r *http.Request) {
	employees := getEmployeeList()
	id, _, _, _, _ := parseForm(r)
	employees[id].Visible = false
	writeEmployeeList(employees)
	common.SendRedirect(w, r, "/employees")
}

func parseForm(r *http.Request) (int, string, string, string, string) {
	r.ParseForm()
	buffer := [5]string{"0"}
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
		case "role":
			buffer[4] = value[0]
		}
	}
	id, e := strconv.Atoi(buffer[0])
	router.CheckAndLogError(e)
	return id, buffer[1], buffer[2], buffer[3], buffer[4]
}
