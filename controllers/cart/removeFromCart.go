package controllers

import (
	"github.com/Ribesh/billing_final/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RemoveFromCart(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var removeItem models.Cart
		name := ctx.Param("name")
		db.Delete(&removeItem, "product_name = ?", name)
		ctx.JSON(200, removeItem)
	}

}
