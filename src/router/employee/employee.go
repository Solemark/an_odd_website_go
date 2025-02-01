package employee

import (
	"an_odd_website/src/router"
	"encoding/csv"
	"os"
	"strconv"
)

type employee struct {
	Id           int
	FirstName    string
	LastName     string
	EmailAddress string
	Role         string
	Visible      bool
}

func (e employee) toCSV() []string {
	return []string{strconv.Itoa(e.Id), e.FirstName, e.LastName, e.EmailAddress, e.Role, strconv.FormatBool(e.Visible)}
}

func parseEmployee(s []string) employee {
	id, _ := strconv.Atoi(s[0])
	visible, _ := strconv.ParseBool(s[5])
	return employee{id, s[1], s[2], s[3], s[4], visible}
}

func getEmployeeList() []employee {
	f, e := os.Open("data/employees.csv")
	router.CheckAndLogError(e)

	r := csv.NewReader(f)
	res, e := r.ReadAll()
	router.CheckAndLogError(e)

	list := []employee{}
	for _, i := range res {
		list = append(list, parseEmployee(i))
	}

	return list
}

func writeEmployeeList(list []employee) {
	f, e := os.Create("data/employees.csv")
	router.CheckAndLogError(e)
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	str := [][]string{}
	for _, e := range list {
		str = append(str, e.toCSV())
	}
	w.WriteAll(str)
}
