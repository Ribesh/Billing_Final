package controllers

import (
	"github.com/Ribesh/billing_final/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product models.Product
		name := ctx.Param("name")
		db.Delete(&product, "product_name = ?", name)
		ctx.JSON(200, product)
	}
}
