package ngw

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type route struct {
	Addr string
	*mux.Router
}

func Start() {
	http.Handle("/", R)
	println("listen " + R.Addr)
	http.ListenAndServe(R.Addr, nil)
}

var R *route

func init() {
	initMgo()
	r := route{":3000", mux.NewRouter()}
	R = &r
}

func (r route) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.URL.String() != "/favicon.ico" {
		fmt.Println(req.URL)
	}
	R.Router.ServeHTTP(w, req)
}

func RestRoute(pat, methods string, f func(a A)) {
	R.HandleFunc(pat, func(w http.ResponseWriter, r *http.Request) {
		var a A
		a.R = r
		a.W = w
		a.Var = mux.Vars(r)
		f(a)
	}).Methods(methods)
}

func Post(pat string, f func(a A)) {
	RestRoute(pat, "POST", f)
}

func Get(pat string, f func(a A)) {
	RestRoute(pat, "GET", f)
}

func Put(pat string, f func(a A)) {
	RestRoute(pat, "PUT", f)
}

func Delete(pat string, f func(a A)) {
	RestRoute(pat, "DELETE", f)
}
