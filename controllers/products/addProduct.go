package controllers

import (
	"net/http"

	"github.com/Ribesh/billing_final/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var products models.Product
		if err := ctx.ShouldBindJSON(&products); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		}

		db.Create(&products)
		ctx.JSON(200, products)
	}
}
