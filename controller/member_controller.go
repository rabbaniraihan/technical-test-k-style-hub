package controller

import (
	"k-style-test/database"
	"k-style-test/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type MemberDB struct {
	database.DB
}

func (memberDB *MemberDB) GetAllMember(ctx *gin.Context) {
	db := memberDB.GetDB()

	var members []model.Member
	tx := db.Find(&members)
	if tx.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": tx.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, members)
}

func (memberDB *MemberDB) AddMember(ctx *gin.Context) {
	db := memberDB.GetDB()
	var newMember model.Member

	err := ctx.ShouldBindJSON(&newMember)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	tx := db.Create(&newMember)
	if tx.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": tx.Error.Error(),
		})
		return
	}

	row := tx.Row()
	err = row.Scan(&newMember.Id, &newMember.Username, &newMember.Gender, &newMember.SkinType, &newMember.SkinColor)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, newMember)
}

func (memberDB *MemberDB) UpdateMember(ctx *gin.Context) {
	db := memberDB.GetDB()
	stringId := ctx.Param("id")

	id, err := strconv.Atoi(stringId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var updatedMember model.Member

	err = ctx.ShouldBindJSON(&updatedMember)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	tx := db.Clauses(clause.Returning{
		Columns: []clause.Column{
			{
				Name: "id",
			},
		}}).Where("id = ?", id).Updates(&updatedMember)

	if tx.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": tx.Error.Error(),
		})
		return
	}

	row := db.Row()
	err = row.Scan(&updatedMember.Id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, updatedMember)
}

func (memberDB *MemberDB) DeleteMember(ctx *gin.Context) {
	db := memberDB.GetDB()
	stringId := ctx.Param("id")

	id, err := strconv.Atoi(stringId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var deletedMember model.Member
	deletedMember.Id = id

	tx := db.Clauses(clause.Returning{}).Delete(&deletedMember)
	if tx.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": tx.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, "message : Member deleted sucesfully")
}
