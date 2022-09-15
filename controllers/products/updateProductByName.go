package controllers

import (
	"net/http"

	"github.com/Ribesh/billing_final/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateProductByName(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product models.Product
		name := ctx.Param("name")

		// if err := db.First(&product, "product_name = ?", name); err != nil {
		// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		// }

		if err := ctx.ShouldBindJSON(&product); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		db.Model(&product).Where("product_name = ?", name).Updates(product)
		ctx.JSON(200, product)
	}
}
