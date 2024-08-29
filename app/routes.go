package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/khairililmi2468gmailcom/toko_online_golang/app/controllers"
)
func (server *Server) initializeRoutes(){
   server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/public/", http.FileServer(staticFileDirectory))
	server.Router.PathPrefix("/public/").Handler(staticFileHandler).Methods("GET")
}