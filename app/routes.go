package app

import (
	"github.com/gorilla/mux"
	"github.com/khairililmi2468gmailcom/toko_online_golang/app/controllers"
)
func (server *Server) initializeRoutes(){
    server.Router = mux.NewRouter()
    server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}