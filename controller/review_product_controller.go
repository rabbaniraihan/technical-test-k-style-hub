package controller

import (
	"k-style-test/database"
	"k-style-test/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReviewProductDB struct {
	database.DB
}

func (reviewDB *ReviewProductDB) GetAllReviewProduct(ctx *gin.Context) {
	db := reviewDB.GetDB()
	var reviewProducts []model.Review_product

	tx := db.Preload("Member").Preload("Like_review").Find(&reviewProducts)
	if tx.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": tx.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, reviewProducts)
}
