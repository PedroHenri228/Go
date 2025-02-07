package controllers

import (
	"net/http"

	use_cases "github.com/PedroHenri228/Go/categories/internal/use-cases"
	"github.com/gin-gonic/gin"
)

type createdCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategory(ctx *gin.Context) {

	var body createdCategoryInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"success": false,
				"erro":    err.Error(),
			})
		return
	}

	useCase := use_cases.NewCreateCategoryUseCase()

	err := useCase.Execute(body.Name)

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
	})
}
