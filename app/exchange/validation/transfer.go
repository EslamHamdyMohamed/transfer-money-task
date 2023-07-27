package validation

import (
	"errors"
	"github.com/gin-gonic/gin"
	"transfer_money/app/exchange/requests"
	"transfer_money/helpers"
)

func ValidateTransferRequest(c *gin.Context) (error, requests.TransferRequest) {
	var request requests.TransferRequest
	err := c.ShouldBindJSON(&request)
	if helpers.CheckError(err) {
		return err, request
	}
	if request.Amount <= 0 {
		return errors.New("Failed Amount should not be negative or zero"), request
	}
	return nil, request
}
