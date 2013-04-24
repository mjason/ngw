package ngw

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var R *mux.Router
var Listen string

func initRoute() {
	R = mux.NewRouter()
}

func Start() {
	http.Handle("/", R)
	if Listen != "" {
		fmt.Println("listen for " + Listen)
		http.ListenAndServe(Listen, nil)
	} else {
		fmt.Println("listen for 127.0.0.1:3000")
		http.ListenAndServe(":3000", nil)
	}
}

func AddFunc(path string, f func(http.ResponseWriter, *http.Request)) {
	R.HandleFunc(path, f)
}

func Route(path string, f func(c Action)) {
	f1 := func(w http.ResponseWriter, r *http.Request) {
		var c Action
		c.R = r
		c.W = w
		f(c)
	}
	R.HandleFunc(path, f1)
}
