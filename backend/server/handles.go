package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ComputePractice/photohosting/backend/data"
)

//GetProfile обрабатывает запросы на получение списка профилей
func GetProfile(w http.ResponseWriter, r *http.Request) {
	binaryData, err := json.Marshal(data.ProfileList)
	if err != nil {
		w.Header().Add("Content-Type", "plain/text; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "JSON marshaling error: %v", err)
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(binaryData)
	if err != nil {
		log.Printf("Hadler Write Error: %v", err)
	}
}
