package route

import (
	"github.com/fernoe1/AP2/assignment-1/payment/internal/adapter/http/server/handler"
	"github.com/fernoe1/AP2/assignment-1/payment/internal/middleware"
	"github.com/fernoe1/AP2/assignment-1/payment/internal/usecase"
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	return r
}

func RegisterPaymentRoute(r *gin.Engine, uc usecase.PaymentUsecase) {
	paymentHandler := handler.PaymentHandler{PaymentUsecase: &uc}

	paymentRoute := r.Group("/payments")
	paymentRoute.POST("", paymentHandler.Post)
	paymentRoute.GET("/:order_id", paymentHandler.Get)
}
