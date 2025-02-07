package use_cases

import (
	"log"

	"github.com/PedroHenri228/Go/categories/internal/entities"
)

type createCategoryUseCase struct {
	// db
}

func NewCreateCategoryUseCase() *createCategoryUseCase {

	return &createCategoryUseCase{}

}

func (u *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	//TODO: persistir entidade para db
	log.Println(category)

	return nil
}