package routes

import (
	"github.com/geges1101/go-balance-manager/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBalanceRoutes = func(router *mux.Router) {
	router.HandleFunc("/balance/", controllers.CreateBalance).Methods("POST")
	router.HandleFunc("/balance/", controllers.GetBalance).Methods("GET")
	router.HandleFunc("/order/", controllers.CreateReport).Methods("GET")
	router.HandleFunc("/balance/{balanceId}", controllers.GetBalanceById).Methods("GET")
	router.HandleFunc("/balance/{balanceId}", controllers.UpdateBalance).Methods("PUT")
	router.HandleFunc("/balance/{balanceId}", controllers.UpdateReserve).Methods("PUT")
	router.HandleFunc("/balance/{balanceId}", controllers.SubtractRevenue).Methods("PUT")
	router.HandleFunc("/balance/{fromId}", controllers.CreateTransfer).Methods("PUT")
	router.HandleFunc("/balance/{balanceId}", controllers.DeleteBalance).Methods("DELETE")
}
