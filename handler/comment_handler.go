package handler

import (
	"fmt"
	"simple-block-api/models"
	"simple-block-api/utils"

	"github.com/gin-gonic/gin"
)

func CommentCreate(ctx *gin.Context) {
	comment := models.Comment{}
	error := ctx.ShouldBindJSON(&comment)

	if error != nil || comment.Content == "" || comment.PostID == 0 {
		fmt.Printf("comment: %+v\n", comment)
		utils.RespWithBizError(ctx, utils.ErrParam)
		return
	}

	comment.UserID = ctx.GetUint("userID")

	db := utils.GetDB()
	db.Create(&comment)

	utils.RespWithData(ctx, comment)
}

func CommentList(ctx *gin.Context) {

	comment := models.Comment{}
	error := ctx.ShouldBindJSON(&comment)

	if error != nil {
		utils.RespWithBizError(ctx, utils.ErrParam)
		return
	}

	db := utils.GetDB()
	var comments []models.Comment
	db.Where("post_id = ?", comment.PostID).Find(&comments)

	utils.RespWithData(ctx, comments)
}
