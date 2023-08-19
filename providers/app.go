package providers

import (
	"github.com/gin-gonic/gin"
	"os"
	"transfer_money/app/routes"
	"transfer_money/helpers"
)

func Run() {
	r := gin.Default()

	routes.TransferRoutes(r)

	err := r.Run(":" + os.Getenv("APP_PORT_LOCAL"))
	helpers.CheckError(err)
}
