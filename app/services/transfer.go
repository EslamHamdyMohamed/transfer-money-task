package services

import (
	"errors"
	"transfer_money/app/exchange/requests"
	appHelper "transfer_money/app/helpers"
	"transfer_money/app/models"
	"transfer_money/database/models/transfer"
	"transfer_money/helpers"
)

func TransferService(request requests.TransferRequest) error {
	err, fromAccount, toAccount := checkAccounts(request.From, request.To)
	if helpers.CheckError(err) {
		return err
	}
	if checkBalanceAmount != nil {
		return errors.New("Failed not enough balance")
	}
	fromAccount, toAccount = appHelper.SetTransfer(fromAccount, toAccount, request.Amount)
	err = models.TransferMoney(fromAccount, toAccount)
	if helpers.CheckError(err) {
		return err
	}
	return nil
}

func checkAccounts(from, to string) (error, transfer.Account, transfer.Account) {
	err, fromAccount := models.FindOrFailAccount(from)
	if helpers.CheckError(err) {
		return err, transfer.Account{}, transfer.Account{}
	}
	err, toAccount := models.FindOrFailAccount(to)
	if helpers.CheckError(err) {
		return err, transfer.Account{}, transfer.Account{}
	}
	return nil, fromAccount, toAccount
}

func checkBalanceAmount(fromAccount transfer.Account, amount float64) error {
	if fromAccount.Balance < amount || amount <= 0 {
		return errors.New("Failed not enough balance or amount not valid ")
	}
	return nil
}
