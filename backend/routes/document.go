package routes

import (
	. "github.com/jigar3/docBabu/utils"
	"net/http"
)

var server = Server{}

func Abc(w http.ResponseWriter, r *http.Request) {
	server.RespondWithJson(w, http.StatusOK, map[string]string{"result": "failure"})
}
