package routes

import (
	"net/http"
	"fmt"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w,"Welcome! %s\n", r.RemoteAddr)
}
