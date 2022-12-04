package router

import (
	"ginchat/docs"
	"ginchat/service"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/index", service.GetIndex)
	r.GET("/user/list", service.GetUserList)
	r.POST("/user/create", service.CreateUser)
	r.POST("/user/delete", service.DeleteUser)
	r.POST("/user/update", service.UpdateUser)
	r.POST("/user/login", service.UserLogin)

	r.POST("/user/sendMsg", service.SendMsg)
	return r
}
