package ngw

import (
	"net/http"
)

type Action struct {
	R *http.Request
	W http.ResponseWriter
}

func (c Action) GetValue(key string) string {
	return c.R.FormValue(key)
}

func (c Action) GetHeader(key string) string {
	return c.R.Header.Get(key)
}

func (c Action) Render(r []byte) {
	c.W.Write(r)
}

func (c Action) Error(err string) {
	c.W.Write([]byte(err))
}
