package server

import (
	"github.com/ComputePractice2018/photohosting/backend/data"

	"github.com/gorilla/mux"
)

func NewRouter(profileList data.Editable, photoList data.Editable) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/photohosting/profile", GetProfiles(profileList)).Methods("GET")
	router.HandleFunc("/api/photohosting/profile", AddProfile(profileList)).Methods("POST")
	router.HandleFunc("/api/photohosting/profile/{id}", EditProfile(profileList)).Methods("PUT")
	router.HandleFunc("/api/photohosting/profile/{id}", DeleteProfile(profileList)).Methods("DELETE")
	router.HandleFunc("/api/photohosting/profile/{id}/photos", GetPhotos(photoList)).Methods("GET")
	router.HandleFunc("/api/photohosting/profile/{id}/photos", AddPhoto(photoList)).Methods("POST")
	router.HandleFunc("/api/photohosting/profile/{id}/photos/{id}", EditPhoto(photoList)).Methods("PUT")
	router.HandleFunc("/api/photohosting/profile/{id}/photos/{id}", DeletePhoto(photoList)).Methods("DELETE")
	return router
}
