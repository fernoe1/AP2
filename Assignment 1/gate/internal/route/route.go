package route

import (
	"github.com/fernoe1/AP2/assignment-1/gate/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	return r
}
