package route

import (
	"github.com/fernoe1/AP2/assignment-1/order/internal/adapter/http/handler"
	"github.com/fernoe1/AP2/assignment-1/order/internal/adapter/http/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	return r
}

func RegisterOrderRoute(r *gin.Engine, uc handler.OrderUsecase) {
	orderHandler := handler.OrderHandler{OrderUsecase: uc}

	orderRoute := r.Group("/orders")
	orderRoute.POST("", orderHandler.Post)
	orderRoute.GET("/:id", orderHandler.Get)
	orderRoute.PATCH("/:id/cancel", orderHandler.Patch)
	orderRoute.PATCH("/:id", orderHandler.PatchStatus)
}
