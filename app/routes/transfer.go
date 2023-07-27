package routes

import (
	"github.com/gin-gonic/gin"
	"transfer_money/app/controllers"
)

func TransferRoutes(r *gin.Engine) {
	r.POST("/transfer", controllers.Transfer)
}
