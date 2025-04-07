package api

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CheckUserInChannel обрабатывает запрос на проверку пользователя в канале
func (a *Api) checkUserInChannel(c *gin.Context) {
	var req struct {
		BotToken    string `json:"bot_token" binding:"required"`
		ChannelLink string `json:"channel_link" binding:"required"`
		UserID      string `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request: " + err.Error(),
		})
		return
	}

	// Преобразуем ID пользователя в int64
	userID, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid user ID: " + err.Error(),
		})
		return
	}

	// Отправляем запрос в gRPC-сервис
	resp, err := a.userChecker.CheckUserInChannel(
		context.Background(),
		req.BotToken,
		req.ChannelLink,
		userID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Internal server error: " + err.Error(),
		})
		return
	}

	// Отправляем ответ клиенту
	c.JSON(http.StatusOK, gin.H{
		"is_member": resp.IsMember,
	})
}
