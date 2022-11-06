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
	var balanceUpdate int
	utils.ParseBody(r, balanceUpdate)
	vars := mux.Vars(r)
	balanceId := vars["balanceId"]
	ID, err := strconv.ParseInt(balanceId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	balanceDetails, db := models.GetBalanceById(ID)
	balanceDetails.Funds += balanceUpdate
	db.Save(&balanceDetails)
	res, _ := json.Marshal(balanceDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
