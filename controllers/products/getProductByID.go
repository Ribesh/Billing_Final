package controllers

import (
	"github.com/Ribesh/billing_final/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProductByName(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product_name models.Product
		name := ctx.Param("name")
		db.First(&product_name, "product_name =?", name)
		ctx.JSON(200, product_name)
	}
}
