package routes

import (
	. "github.com/jigar3/docBabu/utils"
	"net/http"
)

func Abc(w http.ResponseWriter, r *http.Request) {
	RespondWithJson(w, http.StatusOK, map[string]string{"result": "failure"})
}
