package setting

import (
	"an_odd_website/src/router"
	"encoding/json"
	"log"
	"os"
)

func writeSettingList(sl []Setting) {
	f, e := os.Create("data/settings.json")
	router.CheckAndLogError(e)
	defer f.Close()

	j, e := json.MarshalIndent(sl, "", "\t")
	if e != nil {
		log.Fatal(e)
	}

	f.Write(j)
}
