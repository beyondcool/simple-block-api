package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"simple-block-api/models"
	"simple-block-api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PostCreate(ctx *gin.Context) {
	post := models.Post{}
	ctx.ShouldBindJSON(&post)
	// 创建文章时需要提供文章的标题和内容:
	if post.Title == "" || post.Content == "" {
		utils.RespWithBizError(ctx, utils.ErrParam)
		return
	}
	post.UserID = ctx.GetUint("userID")
	db := utils.GetDB()
	db.Create(&post)
	utils.RespWithData(ctx, post)
}

func PostList(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	pageSizeStr := ctx.DefaultQuery("page_size", "25")
	db := utils.GetDB()
	var posts []models.Post

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "username", "email")
	}).Omit(" content").Scopes(utils.Paginate(int(page), int(pageSize))).Find(&posts)

	utils.RespWithData(ctx, posts)
}

func PostDetail(ctx *gin.Context) {
	postID := ctx.Param("id")
	db := utils.GetDB()
	var post models.Post
	result := db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Omit("password")
	}).First(&post, postID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("Post with ID %s not found\n", postID)
		utils.RespWithHttpError(ctx, http.StatusNotFound, "Post not found")
		return
	}
	if result.Error != nil {
		fmt.Printf("Error occurred while fetching post with ID %s: %v\n", postID, result.Error)
		utils.RespWithBizError(ctx, utils.ErrParam)
		return
	}
	utils.RespWithData(ctx, post)
}

func PostDelete(ctx *gin.Context) {
	postID := ctx.Param("id")
	userID := ctx.GetUint("userID")

	db := utils.GetDB()

	result := db.Where("id=? and user_id=?", postID, userID).Delete(&models.Post{}, postID)
	if result.Error != nil || result.RowsAffected == 0 {
		utils.RespWithBizError(ctx, utils.ErrParam)
		return
	}

	utils.RespWithMessage(ctx, "Post deleted successfully.")
}

func PostUpdate(ctx *gin.Context) {
	postID := ctx.Param("id")
	userID := ctx.GetUint("userID")
	var post models.Post
	if err := ctx.ShouldBindJSON(&post); err != nil {
		log.Printf("Failed to bind JSON: %v", err)
		utils.RespWithBizError(ctx, utils.ErrParam)
		return
	}

	post.ID = 0 // 确保 ID 不被覆盖

	db := utils.GetDB()
	result := db.Model(&models.Post{}).Where("id = ? AND user_id = ?", postID, userID).Updates(post)
	if result.Error != nil || result.RowsAffected == 0 {
		log.Printf("Failed to update post: %v", result.Error)
		utils.RespWithBizError(ctx, utils.ErrParam)
		return
	}

	utils.RespWithMessage(ctx, "Post updated successfully.")
}
