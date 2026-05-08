package handler

import (
	"simple-block-api/middleware"
	"simple-block-api/models"
	"simple-block-api/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserRegister(ctx *gin.Context) {
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")

	if name == "" || password == "" {
		utils.RespWithBizError(ctx, utils.ErrUsernameOrPasswordInvalid)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("password error")
	}

	db := utils.GetDB()
	user := models.User{
		Username: name,
		Password: string(hashedPassword),
	}
	db.Create(&user)
	utils.RespWithMessageAndData(ctx, "User registered successfully", user)
}

func UserLogin(ctx *gin.Context) {
	// 1. bind json to paramUser struct
	var paramUser models.User
	if err := ctx.ShouldBindJSON(&paramUser); err != nil {
		utils.RespWithBizError(ctx, utils.ErrParam)
		return
	}

	// 2. check record exist by username
	db := utils.GetDB()
	var dbUser models.User
	err := db.Where("username = ?", paramUser.Username).First(&dbUser).Error
	if err != nil {
		utils.RespWithBizError(ctx, utils.ErrUsernameOrPasswordInvalid)
		return
	}
	// 3. check password
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(paramUser.Password)); err != nil {
		utils.RespWithBizError(ctx, utils.ErrUsernameOrPasswordInvalid)
		return
	}

	// 4. generate JWT token
	token, err := middleware.GenerateToken(dbUser.ID, dbUser.Username, 24)
	if err != nil {
		utils.RespWithBizError(ctx, utils.ErrServer)
		return
	}

	// 5. return token in response
	utils.RespWithMessageAndData(ctx, "Login successful", gin.H{"token": token})

}
