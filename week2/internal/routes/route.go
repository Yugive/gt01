package routes

import (
	"github.com/gin-gonic/gin"
	v1 "week2/internal/routes/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	userApi := v1.NewUserApi()
	apiV1 := r.Group("/api/v1")
	apiV1.Use()
	{
		apiV1.GET("/user/getUserInfo/:id", userApi.GetUserInfo)
	}
	return r
}
