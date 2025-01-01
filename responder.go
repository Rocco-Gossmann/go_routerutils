package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func RespondWithCode(w http.ResponseWriter, status int, msg string, args ...any) {
	w.WriteHeader(status)
	w.Write([]byte(fmt.Sprintf(msg, args...)))
}

func RespondWithError(w http.ResponseWriter, status int, err error) bool {
	if err != nil {

		if status >= 500 {
			log.Panicf("ERROR: %d: %s\n", status, err.Error())
			defer recover()
		}

		w.WriteHeader(status)
		w.Write([]byte(fmt.Sprintf("Error: \n%+v", err.Error())))
		return true
	} else {
		return false
	}
}

func RespondWithJSON(w http.ResponseWriter, status int, data any) {
	if out, err := json.Marshal(data); err == nil {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(status)
		w.Write(out)
	} else {
		RespondWithError(w, 500, err)
	}
}
