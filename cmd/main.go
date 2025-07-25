package main

import (
	"LearningAPI/controller"
	"LearningAPI/db"
	"LearningAPI/repository"
	"LearningAPI/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// camada de repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	// camada usecase
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)
	// Camada de controllers
	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.POST("/product", ProductController.CreateProduct)

	server.GET("/product/:productID", ProductController.GetProductById)

	server.Run(":8000")

}
