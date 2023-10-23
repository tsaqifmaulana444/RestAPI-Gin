package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tsaqifmaulana444/go-gin-restapi/models"
	"github.com/tsaqifmaulana444/go-gin-restapi/controllers/productController"
)

func main() {
	r := gin.Default()
	models.ConnectDB()

	r.GET("api/products", productController.Index)
	r.GET("api/products/:id", productController.Show)
	r.POST("api/products", productController.Create)
	r.PUT("api/products/:id", productController.Update)
	r.DELETE("api/products/:id", productController.Delete)

	r.Run()
}