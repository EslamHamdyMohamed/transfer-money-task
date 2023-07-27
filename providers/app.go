package providers

import (
	"github.com/gin-gonic/gin"
	"transfer_money/app/routes"
)

func Run() {
	r := gin.Default()

	routes.TransferRoutes(r)

	r.Run()
}
