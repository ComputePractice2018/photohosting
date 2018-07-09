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
func GetProfiles(cl data.Editable) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(cl.GetProfiles())
		if err != nil {
			message := fmt.Sprintf("Unable to encode data: %v", err)
			http.Error(w, message, http.StatusInternalServerError)
			log.Println(message)
			return
		}
	}
}

//AddProfile обрабатывает POST запрос для профиля
func AddProfile(cl data.Editable) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var profile data.Profile
		err := json.NewDecoder(r.Body).Decode(&profile)
		if err != nil {
			message := fmt.Sprintf("Unable to decode POST data: %v", err)
			http.Error(w, message, http.StatusUnsupportedMediaType)
			log.Println(message)
			return
		}
		id := cl.AddProfile(profile)
		w.Header().Add("Location", r.URL.String()+"/"+strconv.Itoa(id))
		w.WriteHeader(http.StatusCreated)
	}
}

//EditProfile обрабатывает PUT запрос для профиля
func EditProfile(cl data.Editable) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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

		err = cl.EditProfile(profile, id)
		if err != nil {
			message := fmt.Sprintf("Incorrect ID: %v", err)
			http.Error(w, message, http.StatusBadRequest)
			log.Println(message)
			return
		}
		w.Header().Add("Location", r.URL.String())
		w.WriteHeader(http.StatusAccepted)

	}
}

// DeleteProfile обрабатывает DELETE запрос для профиля
func DeleteProfile(cl data.Editable) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			message := fmt.Sprintf("Incorrect ID: %v", err)
			http.Error(w, message, http.StatusBadRequest)
			log.Println(message)
			return
		}

		err = cl.RemoveProfile(id)
		if err != nil {
			message := fmt.Sprintf("Incorrect ID: %v", err)
			http.Error(w, message, http.StatusBadRequest)
			log.Println(message)
			return
		}
		w.Header().Add("Location", r.URL.String())
		w.WriteHeader(http.StatusNoContent)
	}
}

//GetPhotos обрабатывает запросы на получение списка фотографий
func GetPhotos(cl data.Editable) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(cl.GetPhotos())
		if err != nil {
			message := fmt.Sprintf("Unable to encode data: %v", err)
			http.Error(w, message, http.StatusInternalServerError)
			log.Println(message)
			return
		}
	}
}

//AddPhoto обрабатывает POST запрос для фотографий
func AddPhoto(cl data.Editable) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var photo data.Photo
		err := json.NewDecoder(r.Body).Decode(&photo)
		if err != nil {
			message := fmt.Sprintf("Unable to decode POST data: %v", err)
			http.Error(w, message, http.StatusUnsupportedMediaType)
			log.Println(message)
			return
		}
		id := cl.AddPhoto(photo)
		w.Header().Add("Location", r.URL.String()+"/"+strconv.Itoa(id))
		w.WriteHeader(http.StatusCreated)
	}
}

//EditPhoto обрабатывает PUT запрос для фотографий
func EditPhoto(cl data.Editable) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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

		err = cl.EditPhoto(photo, id)
		if err != nil {
			message := fmt.Sprintf("Incorrect ID: %v", err)
			http.Error(w, message, http.StatusBadRequest)
			log.Println(message)
			return
		}
		w.Header().Add("Location", r.URL.String())
		w.WriteHeader(http.StatusAccepted)

	}
}

// DeletePhoto обрабатывает DELETE запрос для фотографий
func DeletePhoto(cl data.Editable) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			message := fmt.Sprintf("Incorrect ID: %v", err)
			http.Error(w, message, http.StatusBadRequest)
			log.Println(message)
			return
		}

		err = cl.RemovePhoto(id)
		if err != nil {
			message := fmt.Sprintf("Incorrect ID: %v", err)
			http.Error(w, message, http.StatusBadRequest)
			log.Println(message)
			return
		}
		w.Header().Add("Location", r.URL.String())
		w.WriteHeader(http.StatusNoContent)
	}
}
