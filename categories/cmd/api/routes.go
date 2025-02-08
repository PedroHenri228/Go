package main

import (
	"github.com/PedroHenri228/Go/categories/cmd/api/controllers"
	"github.com/PedroHenri228/Go/categories/internal/repositories"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("/categories")
	
	inMemoryCategoryRepository := repositories.NewInMemoryCategoryRepository()


	categoryRoutes.POST("/", func(ctx *gin.Context) {
		controllers.CreateCategory(ctx, inMemoryCategoryRepository)
	} )


	categoryRoutes.GET("/", func(ctx *gin.Context) {
		controllers.ListCategory(ctx, inMemoryCategoryRepository)
	} )
}