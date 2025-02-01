package common

import "net/http"

const CONTENT_TYPE = "Content-Type"
const JSON_RESPONSE = "application/json"
const HTML_RESPONSE = "text/html"
const CSS_RESPONSE = "text/css"
const JS_RESPONSE = "text/javascript"
const ICON_RESPONSE = "iamge/x-icon"

func SendContent(w http.ResponseWriter, msgType string, msg []byte) {
	w.Header().Set(CONTENT_TYPE, msgType)
	w.WriteHeader(http.StatusOK)
	w.Write(msg)
}

func SendRedirect(w http.ResponseWriter, r *http.Request, route string) {
	http.Redirect(w, r, route, http.StatusPermanentRedirect)
}
