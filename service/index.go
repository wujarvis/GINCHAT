package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GlobalIndexResponse struct {
	Message string `json: message`
}

// GetIndex godoc
// @Summary      首页
// @Description  获取首页信息
// @Tags         首页
// @Accept       json
// @Produce      json
// @Router       /index [get]
// @Success      200 {object} GlobalIndexResponse
func GetIndex(ctx *gin.Context) {
	response := GlobalIndexResponse{
		Message: "Welcome!",
	}
	ctx.JSON(http.StatusOK, response)
}
