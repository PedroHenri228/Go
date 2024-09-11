package main

import (
	"api_music/server/service"
	"api_music/server/web"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
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

	router := http.NewServeMux()
	router.HandleFunc("GET /music", MusicHandlers.GetMusic)
	router.HandleFunc("POST /music", MusicHandlers.SetMusic)
	router.HandleFunc("GET /music/{id}", MusicHandlers.GetMusicByID)
	router.HandleFunc("PUT /music/{id}", MusicHandlers.UpdateMusic)
	router.HandleFunc("DELETE /music/{id}", MusicHandlers.DeleteMusic)


	fmt.Println("Servidor Rodando na porta 8080.....")
	http.ListenAndServe(":8080", router)
	
}