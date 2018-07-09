package main

import (
	"log"
	"net/http"

	"github.com/ComputePractice2018/photohosting/backend/server"
	"github.com/gorilla/mux"
)

func main() {
	//name := flag.String("name", "slimz", "имя для приветствия")
	//flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/api/photohosting/profile", server.GetProfile).Methods
	("GET")
	router.HandleFunc("/api/photohosting/profile", server.AddProfile).Methods
	("POST")
	router.HandleFunc("/api/photohosting/profile/{id}", server.EditProfile).Methods
	("PUT")
	router.HandleFunc("/api/photohosting/profile/{id}", server.DeleteProfile).Methods
	("DELETE")
	router.HandleFunc("/api/photohosting/profile/{id}/photos", server.GetProfile).Methods
	("GET")
	router.HandleFunc("/api/photohosting/profile/{id}/photos", server.AddProfile).Methods
	("POST")
	router.HandleFunc("/api/photohosting/profile/{id}/photos/{id}", server.EditProfile).Methods
	("PUT")
	router.HandleFunc("/api/photohosting/profile/{id}/photos/{id}", server.DeleteProfile).Methods
	("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
