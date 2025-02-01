package router

import (
	"an_odd_website/router/common"
	"fmt"
	"net/http"
	"os"
)

func GetWebPage(w http.ResponseWriter, r *http.Request) {
	url := r.URL.RequestURI()
	var file []byte = []byte{}
	var e error = nil

	file, e = os.ReadFile(fmt.Sprintf("static/%s.html", url))
	if url == "/" {
		file, e = os.ReadFile("static/index.html")
	}
	errorHandler(w, e)
	common.SendContent(w, common.HTML_RESPONSE, file)
}

func GetStyles(w http.ResponseWriter, r *http.Request) {
	s := r.PathValue("style")
	f, e := os.ReadFile(fmt.Sprintf("static/styles/%s.css", s))
	errorHandler(w, e)
	common.SendContent(w, common.CSS_RESPONSE, f)
}

func GetScripts(w http.ResponseWriter, r *http.Request) {
	s := r.PathValue("script")
	f, e := os.ReadFile(fmt.Sprintf("static/scripts/%s.js", s))
	errorHandler(w, e)
	common.SendContent(w, common.JS_RESPONSE, f)
}

func GetIcon(w http.ResponseWriter, r *http.Request) {
	f, e := os.ReadFile("static/favicon.ico")
	errorHandler(w, e)
	common.SendContent(w, common.ICON_RESPONSE, f)
}
