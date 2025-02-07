package entities

import (
	"fmt"
	"time"
)

type Category struct {
	ID       uint `json:"id"`
	Name     string `json:"name"`
	CreateAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"update_at"`
}


func NewCategory(name string) (*Category, error) {
	category := &Category{
		Name: name,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	// Regras de negócios

	err := category.isValid()

	if err != nil {
		return nil, err
	}

	return category, nil

}

func (c *Category) isValid() error {

	if(len(c.Name) < 5) {
		return fmt.Errorf("o nome deve ser maior que 5. obtido %d", len(c.Name))
	}

	return nil
}