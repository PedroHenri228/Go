package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"student/controllers"
	"student/db"
	"time"
)

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "erro ao fazer o parse do form: %v", err)
			return
		}

		firstName := r.PostForm.Get("first_name")
		lastName := r.PostForm.Get("last_name")
		email := r.PostForm.Get("email")
		birthDateStr := r.PostForm.Get("birth_date")

		birthDate, err := time.Parse("2006-01-02", birthDateStr)
		if err != nil {
			fmt.Fprintf(w, "Erro ao parsear a data: %v", err)
			return
		}

		err = controllers.Create(firstName, lastName, birthDate, email)
		if err != nil {
			fmt.Fprintf(w, "Erro ao salvar os dados: %v", err)
			return
		}
		http.Redirect(w, r, "/success", http.StatusSeeOther)
		return
	}

	http.ServeFile(w, r, "handlers/templates/student_form.html")
}

func TableStudent(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "handlers/templates/table_student.html")

}
func GetStudentHandler(w http.ResponseWriter, r *http.Request) {

	students, err :=  db.SelectStudents()

	if err != nil {
		http.Error(w, "Erro ao obter os dados dos alunos", http.StatusInternalServerError)
		log.Printf("Erro ao obter os dados dos alunos: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(students)
	if err != nil {
		http.Error(w, "Erro ao codificar os dados", http.StatusInternalServerError)
		log.Printf("Erro ao codificar os dados: %v", err)
	}

}
