package controller

import (
	"k-style-test/database"
	"k-style-test/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductDB struct {
	database.DB
}

func (productDB *ProductDB) GetProductById(ctx *gin.Context) {
	db := productDB.GetDB()
	stringId := ctx.Param("id")

	id, err := strconv.Atoi(stringId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var getProduct model.Product
	getProduct.Id = id

	tx := db.Find(&getProduct)
	if tx.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": tx.Error.Error(),
		})
		return
	}

	row := tx.Row()
	err = row.Scan(&getProduct.Id, &getProduct.NameProduct, &getProduct.Price)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, getProduct)
}
