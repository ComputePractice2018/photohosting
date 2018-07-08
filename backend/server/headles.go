package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ComputePractice2018/photohosting/backend/data"
	"github.com/gorilla/mux"
)

//GetProfiles обрабатывает запросы на получение списка профилей
func GetProfiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(data.GetProfiles())
	if err != nil {
		message := fmt.Sprintf("Unable to encode data: %v", err)
		http.Error(w, message, http.StatusInternalServerError)
		log.Println(message)
		return
	}
}

//AddProfile обрабатывает POST запрос для профиля
func AddProfile(w http.ResponseWriter, r *http.Request) {
	var profile data.Profile
	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		message := fmt.Sprintf("Unable to decode POST data: %v", err)
		http.Error(w, message, http.StatusUnsupportedMediaType)
		log.Println(message)
		return
	}
	id := data.AddProfile(profile)
	w.Header().Add("Location", r.URL.String()+"/"+strconv.Itoa(id))
	w.WriteHeader(http.StatusCreated)
}

//EditProfile обрабатывает PUT запрос для профиля
func EditProfile(w http.ResponseWriter, r *http.Request) {
	var profile data.Profile
	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		message := fmt.Sprintf("Unable to decode PUT data: %v", err)
		http.Error(w, message, http.StatusUnsupportedMediaType)
		log.Println(message)
		return
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		message := fmt.Sprintf("Incorrect ID: %v", err)
		http.Error(w, message, http.StatusBadRequest)
		log.Println(message)
		return
	}

	err = data.EditProfile(profile, id)
	if err != nil {
		message := fmt.Sprintf("Incorrect ID: %v", err)
		http.Error(w, message, http.StatusBadRequest)
		log.Println(message)
		return
	}
	w.Header().Add("Location", r.URL.String())
	w.WriteHeader(http.StatusAccepted)

}

// DeleteProfile обрабатывает DELETE запрос для профиля
func DeleteProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		message := fmt.Sprintf("Incorrect ID: %v", err)
		http.Error(w, message, http.StatusBadRequest)
		log.Println(message)
		return
	}

	err = data.RemoveProfile(id)
	if err != nil {
		message := fmt.Sprintf("Incorrect ID: %v", err)
		http.Error(w, message, http.StatusBadRequest)
		log.Println(message)
		return
	}
	w.Header().Add("Location", r.URL.String())
	w.WriteHeader(http.StatusNoContent)
}

//GetPhotos обрабатывает запросы на получение списка фотографий
func GetPhotos(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(data.GetPhotos())
	if err != nil {
		message := fmt.Sprintf("Unable to encode data: %v", err)
		http.Error(w, message, http.StatusInternalServerError)
		log.Println(message)
		return
	}
}

//AddPhoto обрабатывает POST запрос для фотографий
func AddPhoto(w http.ResponseWriter, r *http.Request) {
	var photo data.Photo
	err := json.NewDecoder(r.Body).Decode(&photo)
	if err != nil {
		message := fmt.Sprintf("Unable to decode POST data: %v", err)
		http.Error(w, message, http.StatusUnsupportedMediaType)
		log.Println(message)
		return
	}
	id := data.AddPhoto(photo)
	w.Header().Add("Location", r.URL.String()+"/"+strconv.Itoa(id))
	w.WriteHeader(http.StatusCreated)
}

//EditPhoto обрабатывает PUT запрос для фотографий
func EditPhoto(w http.ResponseWriter, r *http.Request) {
	var photo data.Photo
	err := json.NewDecoder(r.Body).Decode(&photo)
	if err != nil {
		message := fmt.Sprintf("Unable to decode PUT data: %v", err)
		http.Error(w, message, http.StatusUnsupportedMediaType)
		log.Println(message)
		return
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		message := fmt.Sprintf("Incorrect ID: %v", err)
		http.Error(w, message, http.StatusBadRequest)
		log.Println(message)
		return
	}

	err = data.EditPhoto(photo, id)
	if err != nil {
		message := fmt.Sprintf("Incorrect ID: %v", err)
		http.Error(w, message, http.StatusBadRequest)
		log.Println(message)
		return
	}
	w.Header().Add("Location", r.URL.String())
	w.WriteHeader(http.StatusAccepted)

}

// DeletePhoto обрабатывает DELETE запрос для фотографий
func DeletePhoto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		message := fmt.Sprintf("Incorrect ID: %v", err)
		http.Error(w, message, http.StatusBadRequest)
		log.Println(message)
		return
	}

	err = data.RemovePhoto(id)
	if err != nil {
		message := fmt.Sprintf("Incorrect ID: %v", err)
		http.Error(w, message, http.StatusBadRequest)
		log.Println(message)
		return
	}
	w.Header().Add("Location", r.URL.String())
	w.WriteHeader(http.StatusNoContent)
}
