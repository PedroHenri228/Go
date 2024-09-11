package main

import (
	"api_music/server/service"
	"api_music/server/web"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/music_api"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	MusicService := service.NewMusicService(db)
	MusicHandlers := web.NewMusicHandler(MusicService)

	router := mux.NewRouter()
	router.HandleFunc("/music", MusicHandlers.GetMusic).Methods("GET")
	router.HandleFunc("/music", MusicHandlers.SetMusic).Methods("POST")
	router.HandleFunc("/music/{id}", MusicHandlers.GetMusicByID).Methods("GET")
	router.HandleFunc("/music/{id}", MusicHandlers.UpdateMusic).Methods("PUT")
	router.HandleFunc("/music/{id}", MusicHandlers.DeleteMusic).Methods("DELETE")


	fmt.Println("Servidor Rodando na porta 8080.....")
	http.ListenAndServe(":8080", router)
	
}