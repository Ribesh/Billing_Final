package main

import (
	"github.com/Ribesh/billing_final/middleware"

	cart "github.com/Ribesh/billing_final/controllers/cart"
	products "github.com/Ribesh/billing_final/controllers/products"

	"github.com/Ribesh/billing_final/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	db = database.Connect()
}

func main() {

	router := gin.Default()

	router.Use(middleware.Cors())
	api := router.Group("api")
	apiProducts := api.Group("/products")
	{
		apiProducts.GET("", products.GetAllProducts(db))
		apiProducts.GET("/:name", products.GetProductByName(db))
		apiProducts.POST("/", products.CreateProduct(db))
		apiProducts.DELETE("/:name", products.DeleteProduct(db))
		apiProducts.PUT("/:name", products.UpdateProductByName(db))
	}
	apiCart := api.Group("/cart")
	{
		apiCart.GET("/", cart.GetAllCart(db))
		apiCart.POST("/:name", cart.AddToCart(db))
		apiCart.PUT("/:name", cart.UpdateCart(db))
		apiCart.DELETE("/:name", cart.RemoveFromCart(db))
		apiCart.DELETE("/", cart.ClearCart(db))
	}

	router.Run(":8090")
}
