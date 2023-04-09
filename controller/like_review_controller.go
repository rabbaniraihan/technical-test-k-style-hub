package controller

import (
	"k-style-test/database"
	"k-style-test/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type LikeDB struct {
	database.DB
}

func (likeDB *LikeDB) AddLikeReview(ctx *gin.Context) {
	db := likeDB.GetDB()
	var newLike model.Like_review

	err := ctx.ShouldBindJSON(&newLike)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	tx := db.Create(&newLike)
	if tx.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": tx.Error.Error(),
		})
		return
	}

	row := tx.Row()
	err = row.Scan(&newLike.Id, &newLike.ReviewId, &newLike.MemberId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, newLike)
}

func (likeDB *LikeDB) DeleteLikeReview(ctx *gin.Context) {
	db := likeDB.GetDB()
	stringId := ctx.Param("id")

	id, err := strconv.Atoi(stringId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var deletedLike model.Like_review
	deletedLike.Id = id

	tx := db.Clauses(clause.Returning{}).Delete(&deletedLike)
	if tx.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": tx.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, deletedLike)
}
