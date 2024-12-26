package controllers

import (
	"student/db"
	"time"
	"log"
)

type Students_Values struct {
	First_name string
	Last_name  string
	Birth_date time.Time
	Email      string
}

func Create(f_name string, l_name string, b_date time.Time, email string) error {

	s := Students_Values{
		First_name: f_name,
		Last_name:  l_name,
		Birth_date: b_date,
		Email:      email,
	}

	err := db.Insert("students", s)
	if err != nil {
		log.Printf("Erro ao inserir aluno: %v", err)
		return err
	}

	log.Println("Aluno inserido com sucesso!")
	return nil
}


