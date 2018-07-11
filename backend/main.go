package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/ComputePractice2018/photohosting/backend/data"
	"github.com/ComputePractice2018/photohosting/backend/server"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	connection := flag.String("connection", "photohosting:SuperSecretPassword@tcp(db:3306)/photohosting", "mysql connection string")
	flag.Parse()

	profileList, err := data.NewDBProfileList(*connection)
	defer profileList.Close()
	//profileList := data.NewProfileList()

	if err != nil {
		log.Fatalf("DB error: %+v", err)
	}

	router := server.NewRouter(profileList)

	log.Fatal(http.ListenAndServe(":8080", router))
}
