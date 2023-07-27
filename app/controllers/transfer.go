package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"transfer_money/app/exchange/validation"
	"transfer_money/app/services"
	"transfer_money/helpers"
)

func Transfer(c *gin.Context) {
	err, request := validation.ValidateTransferRequest(c)
	if helpers.CheckError(err) {
		helpers.Response(c, http.StatusBadRequest, "error", nil, err.Error())
		return
	}
	err = services.TransferService(request)
	if helpers.CheckError(err) {
		helpers.Response(c, http.StatusBadRequest, "error", nil, err.Error())
		return
	}
	helpers.Response(c, http.StatusOK, "Process run successfully", "", nil)
}
