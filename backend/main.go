package main

import (
	"fmt"
	"log"
	"net/http"
	"timex/api"
	"timex/database"
)

func main() {
	addr := ":8080"
	db, err := database.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("opened postgres db")
	app := api.NewApp(db)
	router := app.SetupRoutes()
	srv := http.Server{
		Addr:    addr,
		Handler: router,
	}
	fmt.Println("running server at ", addr)
	log.Fatal(srv.ListenAndServe())
}
