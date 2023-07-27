package models

import (
	"gorm.io/gorm"
	"transfer_money/database"
	"transfer_money/database/models/transfer"
	"transfer_money/helpers"
)

func FindOrFailAccount(acouuntId string) (error, transfer.Account) {
	db := database.DB
	var account transfer.Account
	if err := db.Where("id = ?", acouuntId).First(&account).Error; err != nil {
		return err, account
	}
	return nil, account
}

func TransferMoney(fromAccount, toAccount transfer.Account) error {
	db := database.DB

	err := db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Save(&fromAccount).Error; err != nil {
			return err
		}
		if err := tx.Save(&toAccount).Error; err != nil {
			return err
		}
		return nil
	})

	if helpers.CheckError(err) {
		return err
	}
	return nil
}
