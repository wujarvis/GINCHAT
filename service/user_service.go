package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GlobalUserListResponse struct {
	Msg  string              `json:"msg"`
	Data []*models.UserBasic `json:"data"`
}

// GetUserList godoc
// @Summary      用户列表
// @Description  获取用户信息
// @Tags         用户模块
// @Accept       json
// @Produce      json
// @Router       /usr/list [get]
// @Success      200 {object} GlobalUserListResponse
func GetUserList(ctx *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	response := GlobalUserListResponse{
		Msg:  "ok",
		Data: data,
	}
	ctx.JSON(http.StatusOK, response)
}

// CreateUser
// @Summary      创建用户
// @Description  创建一个新用户
// @Tags         用户模块
// @param name query string true "用户名"
// @param password query string true "密码"
// @param repassword query string true "确认密码"
// @Router       /usr/create [get]
// @Success      200 {string} json{"code","message"}
func CreateUser(ctx *gin.Context) {
	/*
		用户注册
	*/
	user := models.UserBasic{}
	user.Name = ctx.PostForm("username")
	password := ctx.PostForm("password")
	repassword := ctx.PostForm("repassword")

	// 密码校验
	if password != repassword {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": "两次密码不一致",
		})
		return
	}

	// 用户重复注册校验
	err := models.FindUserByName(user.Name)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": "该用户已存在",
		})
		return
	}

	// 加密
	salt := fmt.Sprintf("%06d", rand.Int31())
	user.Password = utils.MakePassword(password, salt)
	user.Salt = salt

	createUser := models.CreateUser(&user)
	if createUser.Error != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": fmt.Sprintf("新增用户失败：%s", createUser.Error),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "新增用户成功",
		})
	}

}

func DeleteUser(ctx *gin.Context) {
	/*
		删除用户
	*/
	user := models.UserBasic{}
	id, err := strconv.Atoi(ctx.PostForm("id"))
	if err != nil {
		return
	}
	user.ID = uint(id)
	models.DeleteUser(&user)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "删除用户成功",
	})
}

func UpdateUser(ctx *gin.Context) {
	/*
		更新用户
	*/
	userID := ctx.PostForm("id")
	userName := ctx.PostForm("userName")
	password := ctx.PostForm("password")
	id, _ := strconv.Atoi(userID)
	//oldPassword := ctx.Query("oldPassword")
	//newPassword := ctx.Query("newPassword")
	//user := models.UserBasic{}
	//userInfo := utils.DB.Where("id = ?", userID).First(&user)
	user := models.UserBasic{
		Name:     userName,
		Password: password,
	}
	user.ID = uint(id)
	models.UpdateUser(&user)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "更新用户成功",
	})
}

func UserLogin(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	user := models.FindUserByName(username)
	if user == nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": "用户不存在",
		})
		return
	}
	// 密码校验，用户传入的密码与myql中存储的密码是否一致
	validPassword := utils.ValidPassword(password, user.Salt, user.Password)
	if !validPassword {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": "用户密码错误",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
	})
}
