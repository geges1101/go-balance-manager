package routes

import(
	"github.com/gorilla/mux"
	"github.com/geges1101/go-balance-manager/pkg/controllers"
)

var RegisterBalanceRoutes = func (router *mux.Router)  {
	router.HandleFunc("/balance/", controllers.CreateBalance).Methods("POST")
	router.HandleFunc("/balance/", controllers.GetBalance).Methods("GET")
	router.HandleFunc("/balance/{balanceId}", controllers.GetBalanceById).Methods("GET")
	router.HandleFunc("/balance/{balanceId}", controllers.UpdateBalance).Methods("PUT")
	router.HandleFunc("/balance/{balanceId}", controllers.DeleteBalance).Methods("DELETE")
}