package repositories

import "github.com/PedroHenri228/Go/categories/internal/entities"

type ICategoryRepository interface {
	Save(category *entities.Category) error
	List() ([]*entities.Category ,error)
}