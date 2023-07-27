package helpers

import (
	"transfer_money/database/models/transfer"
)

func SetTransfer(fromAccount, toAccount transfer.Account, amount float64) (transfer.Account, transfer.Account) {
	fromAccount = decreaseAccountBalance(fromAccount, amount)
	toAccount = increaseAccountBalance(toAccount, amount)
	return fromAccount, toAccount
}

func increaseAccountBalance(account transfer.Account, amount float64) transfer.Account {
	account.Balance += amount
	return account
}

func decreaseAccountBalance(account transfer.Account, amount float64) transfer.Account {
	account.Balance -= amount
	return account
}
