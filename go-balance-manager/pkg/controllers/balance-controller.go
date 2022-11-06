package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/geges1101/go-balance-manager/pkg/models"
	"github.com/geges1101/go-balance-manager/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBalance models.Balance

func GetBalance(w http.ResponseWriter, r *http.Request) {
	newBalances := models.GetAllBalances()
	res, _ := json.Marshal(newBalances)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBalanceById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	balanceId := vars["balanceId"]
	ID, err := strconv.ParseInt(balanceId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	balanceDetails, _ := models.GetBalanceById(ID)
	res, _ := json.Marshal(balanceDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBalance(w http.ResponseWriter, r *http.Request) {
	CreateBalance := &models.Balance{}
	utils.ParseBody(r, CreateBalance)
	b := CreateBalance.CreateBalance()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	balanceId := vars["balanceId"]
	ID, err := strconv.ParseInt(balanceId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	balance := models.DeleteBalance(ID)
	res, _ := json.Marshal(balance)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	balanceId := vars["balanceId"]
	balanceUpdate := vars["amount"]
	ID, err := strconv.ParseInt(balanceId, 0, 0)
	amount, err1 := strconv.ParseInt(balanceUpdate, 0, 0)
	if err != nil || err1 != nil {
		fmt.Println("error while parsing")
	}
	balanceDetails, db := models.GetBalanceById(ID)
	if balanceDetails.Funds <= 10000000-amount ||
		balanceDetails.Funds >= amount {
		balanceDetails.Funds += amount
		db.Save(&balanceDetails)
	}
	res, _ := json.Marshal(balanceDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateReserve(w http.ResponseWriter, r *http.Request) {
	var req = &models.Order{}
	utils.ParseBody(r, req)
	vars := mux.Vars(r)
	balanceId := vars["balanceId"]
	balanceUpdate := vars["amount"]
	ID, err := strconv.ParseInt(balanceId, 0, 0)
	amount, err1 := strconv.ParseInt(balanceUpdate, 0, 0)
	if err != nil || err1 != nil {
		fmt.Println("error while parsing")
	}
	balanceDetails, db := models.GetBalanceById(ID)
	if balanceDetails.Funds >= amount {
		balanceDetails.Funds -= amount
		balanceDetails.Reserve += amount
		db.Save(&balanceDetails)
	}
	res, _ := json.Marshal(balanceDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func SubtractRevenue(w http.ResponseWriter, r *http.Request) {
	var req = &models.Order{}
	utils.ParseBody(r, req)
	vars := mux.Vars(r)
	balanceId := vars["balanceId"]
	balanceUpdate := vars["amount"]
	ID, err := strconv.ParseInt(balanceId, 0, 0)
	amount, err1 := strconv.ParseInt(balanceUpdate, 0, 0)
	if err != nil || err1 != nil {
		fmt.Println("error while parsing")
	}
	balanceDetails, db := models.GetBalanceById(ID)
	if balanceDetails.Reserve < amount {
		fmt.Println("insufficient funds")
		return

	}
	balanceDetails.Reserve -= amount
	db.Save(&balanceDetails)
	db.Save(req)

	res, _ := json.Marshal(balanceDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateTransfer(w http.ResponseWriter, r *http.Request) {
	var req = &models.Transfer{}
	utils.ParseBody(r, req)
	vars := mux.Vars(r)
	fromId, err := strconv.ParseInt(vars["fromId"], 0, 0)
	toId, err1 := strconv.ParseInt(vars["toId"], 0, 0)
	amount, err2 := strconv.ParseInt(vars["amount"], 0, 0)
	if err != nil || err1 != nil || err2 != nil {
		fmt.Println("error while parsing")
	}
	fromDetails, db := models.GetBalanceById(fromId)
	toDetails, _ := models.GetBalanceById(toId)
	if fromDetails.Funds < amount {
		fmt.Println("insufficent funds for transfer")
		return
	}
	if toDetails.Funds >= 10000000 {
		fmt.Println("reciever funds are at max")
		return
	}
	toDetails.Funds += amount
	fromDetails.Funds -= amount
	db.Save(&fromDetails)
	db.Save(&toDetails)
	left, _ := json.Marshal(fromDetails)
	right, _ := json.Marshal(toDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(left)
	w.Write(right)
}
