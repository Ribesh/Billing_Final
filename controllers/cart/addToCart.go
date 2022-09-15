package controllers

import (
	"log"
	"net/http"

	"github.com/Ribesh/billing_final/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddToCart(db *gorm.DB) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var searchProduct models.Product
		name := ctx.Param("name")
		log.Println(name)
		result := db.First(&searchProduct, "product_name = ?", name)

		var addCart models.Cart

		addCart.ProductName = searchProduct.ProductName
		addCart.PricePiece = searchProduct.PricePiece
		addCart.PriceDozen = searchProduct.PriceDozen
		addCart.PriceBox = searchProduct.PriceBox
		addCart.Status = searchProduct.Status
		addCart.BillingItem = "yes"

		if result.Error == nil {
			log.Println(&result)
			db.Create(&addCart)
			ctx.JSON(http.StatusOK, &addCart)
		}
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, "Item not found")
		}

		// if err := ctx.ShouldBindJSON(&addCart); err != nil {
		// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "error adding to cart"})
		// }

		//db.Model(&addCart).Where("product_name = ?", name).Updates(searchProduct)
		//log.Println(searchProduct.ProductName)
		//db.Debug().Select(&searchProduct).Create(&addCart)
		// db.Where("product_name = ?", name).Preload("Product").Create(&addCart)
		// log.Println(addCart.ProductName)
		// ctx.JSON(200, addCart)
	}
}
