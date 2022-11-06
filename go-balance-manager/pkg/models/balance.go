package models

import (
	"github.com/geges1101/go-balance-manager/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Balance struct {
	gorm.Model
	Id        int64  `json:"id"`
	FirstName string `json:"name"`
	LastName  string `json:"surname"`
	Funds     int64  `json:"funds"`
	Reserve   int64  `json:"reserve"`
}

type Order struct {
	BalanceId int64 `json:"balanceId"`
	ServiceId int64 `json:"serviceId"`
	Amount    int64 `json:"amount"`
}

type Transfer struct {
	FromAccountID string `json:"fromId" binding:"required,min=1"`
	ToAccountID   string `json:"toId" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Balance{})
}

func (b *Balance) CreateBalance() *Balance {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBalances() []Balance {
	var Balances []Balance
	db.Find(&Balances)
	return Balances
}

func GetBalanceById(Id int64) (*Balance, *gorm.DB) {
	var getBalance Balance
	db := db.Where("ID=?", Id).Find(&getBalance)
	return &getBalance, db
}

func DeleteBalance(ID int64) Balance {
	var balance Balance
	db.Where("ID=?", ID).Delete(balance)
	return balance
}
