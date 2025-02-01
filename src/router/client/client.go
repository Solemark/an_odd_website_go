package client

import (
	"an_odd_website/src/router"
	"encoding/csv"
	"os"
	"strconv"
)

type client struct {
	Id           int
	FirstName    string
	LastName     string
	EmailAddress string
	Visible      bool
}

func (c client) toCSV() []string {
	return []string{strconv.Itoa(c.Id), c.FirstName, c.LastName, c.EmailAddress, strconv.FormatBool(c.Visible)}
}

func parseClient(s []string) client {
	id, _ := strconv.Atoi(s[0])
	visible, _ := strconv.ParseBool(s[4])
	return client{id, s[1], s[2], s[3], visible}
}

func getClientList() []client {
	f, e := os.Open("data/clients.csv")
	router.CheckAndLogError(e)

	r := csv.NewReader(f)
	res, e := r.ReadAll()
	router.CheckAndLogError(e)

	list := []client{}
	for _, i := range res {
		list = append(list, parseClient(i))
	}

	return list
}

func writeClientList(list []client) {
	f, e := os.Create("data/clients.csv")
	router.CheckAndLogError(e)
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	str := [][]string{}
	for _, c := range list {
		str = append(str, c.toCSV())
	}
	w.WriteAll(str)
}
