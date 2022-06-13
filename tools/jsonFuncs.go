package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errorObject struct {
	Message string `json:"message"`
}

func WriteJsonBadRequest(rw http.ResponseWriter, err error) {
	writeJson(rw, http.StatusBadRequest, &errorObject{err.Error()})
}

func WriteJsonInternalError(rw http.ResponseWriter, err error) {
	writeJson(rw, http.StatusInternalServerError, &errorObject{err.Error()})
}

func WriteJsonOk(rw http.ResponseWriter, res interface{}) {
	writeJson(rw, http.StatusOK, res)
}

func writeJson(rw http.ResponseWriter, status int, res interface{}) {
	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(status)
	err := json.NewEncoder(rw).Encode(res)
	if err != nil {
		fmt.Printf("Error writing response: %s", err)
	}
}
