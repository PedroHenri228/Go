package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	ID        int    `json:"ID"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	BirthDate string `json:"BirthDate"`
	Email     string `json:"Email"`
}

func getConnection() (*sql.DB, context.Context, context.CancelFunc, error) {
	dsn := "meu_usuario:minha_senha@tcp(127.0.0.1:3306)/student_db"

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Erro ao criar a conexão com o MySQL: ", err)
		return nil, nil, nil, err
	}

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal("Erro ao verificar a conexão com o MySQL: ", err)
		return nil, nil, nil, err
	}

	log.Println("Conexão bem-sucedida com o MySQL!")
	return db, ctx, cancel, nil
}

func Insert(table string, data interface{}) error {
	db, ctx, cancel, err := getConnection()
	if err != nil {
		return fmt.Errorf("erro ao conectar ao banco de dados: %w", err)
	}
	defer db.Close()
	defer cancel()

	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)

	var columns []string
	var placeholders []string
	var values []interface{}

	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i).Interface()
		fieldType := t.Field(i).Type.Name()
		columnName := t.Field(i).Tag.Get("json") // Usando a tag 'json' para o nome da coluna

		if fieldType == "Time" {
			if timeValue, ok := fieldValue.(time.Time); ok {
				fieldValue = timeValue.Format("2006-01-02")
			}
		}

		columns = append(columns, columnName) // Usando o nome da coluna
		placeholders = append(placeholders, "?")
		values = append(values, fieldValue)
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		table,
		join(columns, ", "),
		join(placeholders, ", "),
	)

	_, err = db.ExecContext(ctx, query, values...)
	if err != nil {
		return fmt.Errorf("erro ao executar o INSERT: %w", err)
	}

	log.Println("Registro inserido com sucesso!")
	return nil
}

func join(elements []string, sep string) string {
	return strings.Join(elements, sep)
}

func SelectStudents() ([]Student, error) {
	db, ctx, cancel, err := getConnection() 
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %w", err)
	}
	defer db.Close()
	defer cancel()

	query := "SELECT id, first_name, last_name, birth_date, email FROM students"
	rows, err := db.QueryContext(ctx, query) 
	if err != nil {
		return nil, fmt.Errorf("erro ao executar o SELECT: %w", err)
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var student Student
		err := rows.Scan(&student.ID, &student.FirstName, &student.LastName, &student.BirthDate, &student.Email)
		if err != nil {
			return nil, fmt.Errorf("erro ao escanear os dados: %w", err)
		}
		students = append(students, student)
	}

	return students, nil
}


