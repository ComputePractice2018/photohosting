package main

import (
	"log"
	"net/http"

	"github.com/ComputePractice2018/photohosting/backend/data"
	"github.com/ComputePractice2018/photohosting/backend/server"
)

func main() {
	//name := flag.String("name", "slimz", "имя для преветствия")
	//flag.Parse()
	profileList := data.NewProfileList()
	router := server.NewRouter(profileList)
	photoList := data.NewPhotoList()
	router := server.NewRouter(photoList)

	log.Fatal(http.ListenAndServe(":8080", router))
}
