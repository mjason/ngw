package ngw

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type A struct {
	W   http.ResponseWriter
	R   *http.Request
	Var map[string]string
}

func (a A) Header(key string) string {
	return a.R.Header.Get(key)
}

func (a A) UrlValue(key string) string {
	return a.Var[key]
}

func (a A) From() url.Values {
	a.R.FormValue("")
	return a.R.Form
}

func (a A) OK(body []byte) {
	a.W.Write(body)
}

func (a A) Error(body []byte, code int) {
	a.W.WriteHeader(code)
	a.W.Write(body)
}

func (a A) OKJSON(data interface{}) {
	json_, _ := json.Marshal(data)
	a.OKJSON(json_)
}
