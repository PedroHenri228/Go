package controllers

import (
	"net/http"

	"github.com/PedroHenri228/Go/categories/internal/repositories"
	use_cases "github.com/PedroHenri228/Go/categories/internal/use-cases"
	"github.com/gin-gonic/gin"
)


func ListCategory(ctx *gin.Context, repository repositories.ICategoryRepository ) {
	useCase := use_cases.NewListCategoryUseCase(repository)

	categories, err := useCase.Execute()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"success": false,
				"erro":    err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"categories": categories,
	})
}
