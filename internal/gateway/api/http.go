package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/tg-checker/gen/proto"
)

type UserChecker interface {
	CheckUserInChannel(
		ctx context.Context,
		botToken,
		channelLink string,
		userID int64) (*pb.CheckUserResponse, error)
}

type Api struct {
	userChecker UserChecker
}

func NewApi(userChecker UserChecker) *Api {
	return &Api{
		userChecker: userChecker,
	}
}

func (a *Api) RegisterRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	v1 := router.Group("/api/v1")
	{
		v1.POST("/check-user", a.checkUserInChannel)
	}

	return router
}
