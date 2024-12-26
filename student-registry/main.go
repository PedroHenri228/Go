package main

import (
	"fmt"
	"log"
	"net/http"
	"student/handlers"
)

func main() {
	// Define a rota que será usada
	http.HandleFunc("/", handlers.TableStudent)
	http.HandleFunc("/register", handlers.StudentHandler)
	http.HandleFunc("/students", handlers.GetStudentHandler)

	// Inicia o servidor na porta 8080
	fmt.Println("Servidor rodando na porta 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
