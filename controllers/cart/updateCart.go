package controllers

import (
	"net/http"

	"github.com/Ribesh/billing_final/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateCart(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var updateCart models.Cart
		name := ctx.Param("name")

		if err := ctx.ShouldBindJSON(&updateCart); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		db.Model(&updateCart).Where("product_name = ?", name).Updates(updateCart)
		ctx.JSON(200, updateCart)

	}
}
