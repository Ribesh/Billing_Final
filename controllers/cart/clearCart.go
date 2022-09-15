package controllers

import (
	"github.com/Ribesh/billing_final/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ClearCart(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var delete models.Cart
		item := "yes"
		db.Unscoped().Delete(&delete, "billing_item =?", item)
		ctx.JSON(200, delete)

	}
}
