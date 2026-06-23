// Package server assembles the HTTP router from the application's features.
package server

import (
	"mockapi/internal/mockmock"
	"mockapi/internal/web"
	"mockapi/pkg/httputil"

	"github.com/gin-gonic/gin"
)

func NewRouter(mockmockH *mockmock.Handler) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(httputil.RequestID(), gin.Logger(), gin.Recovery())

	r.GET("/", gin.WrapF(web.Handler()))
	mockmockH.RegisterRoutes(r)

	return r
}
