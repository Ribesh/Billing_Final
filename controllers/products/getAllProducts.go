package controllers

import (
	"github.com/Ribesh/billing_final/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllProducts(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var products []models.Product
		db.Find(&products)
		ctx.JSON(200, products)

	}
}
