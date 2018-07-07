package main

import (
	"log"
	"net/http"

	"github.com/ComputePractice2018/photohosting/backend/server"
)

func main() {
	//name := flag.String("name", "slimz", "имя для приветствия")
	//flag.Parse()

	http.HandleFunc("/api/photohosting/profile", server.GetProfile)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
