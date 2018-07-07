package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ComputePractice2018/photohosting/backend/data"
)

//ProfileHandler обрабатывает обрабатывает все запросы к /api/photohosting/profile
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		GetProfile(w, r)
		return
	}
	if r.Method == "POST" {
		AddProfile(w, r)
		return
	}

	message := fmt.Sprintf("Method %s is not allowed", r.Method)
	http.Error(w, message, http.StatusMethodNotAllowed)
	log.Println(message)

}

//GetProfile обрабатывает запросы на получение списка профилей
func GetProfile(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request method: %s", r.Method)
	binaryData, err := json.Marshal(data.ProfileList)
	if err != nil {
		message := fmt.Sprintf("JSON marshaling error: %v", err)
		http.Error(w, message, http.StatusInternalServerError)
		log.Println(message)
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(binaryData)
	if err != nil {
		message := fmt.Sprintf("JSON marshaling error: %v", err)
		http.Error(w, message, http.StatusInternalServerError)
		log.Println(message)
	}
}

//AddProfile обрабатывает POST запрос
func AddProfile(w http.ResponseWriter, r *http.Request) {
	var profile data.Profile
	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		message := fmt.Sprintf("Unable to decode POST data: %v", err)
		http.Error(w, message, http.StatusUnsupportedMediaType)
		log.Printf(message)
		return
	}
	log.Printf("%+v", profile)
	w.WriteHeader(http.StatusCreated)
}
