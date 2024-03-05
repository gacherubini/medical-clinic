package api

import (
	"net/http"

	_ "github.com/lib/pq"
)

func HandlerHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("oi"))
}
