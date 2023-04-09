package router

import (
	"k-style-test/controller"
	"k-style-test/database"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	db := &database.DB{}

	db.StartDB()

	memberDB := &controller.MemberDB{
		DB: *db,
	}

	productDB := &controller.ProductDB{
		DB: *db,
	}

	reviewProductDB := &controller.ReviewProductDB{
		DB: *db,
	}

	likeDB := &controller.LikeDB{
		DB: *db,
	}

	router.GET("/member", memberDB.GetAllMember)
	router.POST("/member", memberDB.AddMember)
	router.PUT("/member/:id", memberDB.UpdateMember)
	router.DELETE("/member/:id", memberDB.DeleteMember)

	router.GET("/product/:id", productDB.GetProductById)

	router.GET("/review_product", reviewProductDB.GetAllReviewProduct)

	router.POST("/like_review", likeDB.AddLikeReview)
	router.DELETE("/like_review/:id", likeDB.DeleteLikeReview)

	return router
}
