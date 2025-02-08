package use_cases

import (

	"github.com/PedroHenri228/Go/categories/internal/entities"
	"github.com/PedroHenri228/Go/categories/internal/repositories"
)

type createCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewCreateCategoryUseCase(repository repositories.ICategoryRepository) *createCategoryUseCase {

	return &createCategoryUseCase{repository}

}

func (u *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	//TODO: Verifique se o nome da categoria já existe
	err = u.repository.Save(category)

	if err != nil {
		return err
	}

	return nil
}