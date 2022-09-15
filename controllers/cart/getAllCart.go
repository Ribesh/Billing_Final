package controllers

import (
	"github.com/Ribesh/billing_final/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllCart(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var getall []models.Cart
		db.Find(&getall)
		ctx.JSON(200, &getall)
	}
}
