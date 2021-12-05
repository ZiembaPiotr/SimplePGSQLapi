package router

import (
	"github.com/gorilla/mux"
	"github.com/mux/router/api"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()

	playerRoutes(router)
	clubRoutes(router)

	return router
}

func playerRoutes(router *mux.Router) {
	playerRouter := router.PathPrefix("/players").Subrouter()

	playerRouter.HandleFunc("/get-all", api.GetAllPlayers()).Methods("GET")
	playerRouter.HandleFunc("/get/{name}", api.GetPlayer()).Methods("GET")
	playerRouter.HandleFunc("/create-new", api.CreateNewPlayer()).Methods("POST")
	playerRouter.HandleFunc("/delete/{name}", api.DeletePlayer()).Methods("DELETE")
	playerRouter.HandleFunc("/update/{name}", api.UpdatePlayer()).Methods("PUT")
}

func clubRoutes(router *mux.Router) {
	clubRouter := router.PathPrefix("/clubs").Subrouter()

	clubRouter.HandleFunc("/get-all", api.GetAllClubs()).Methods("GET")
	clubRouter.HandleFunc("/get/{name}", api.GetClub()).Methods("GET")
	clubRouter.HandleFunc("/create-new", api.CreateNewClub()).Methods("POST")
	clubRouter.HandleFunc("/delete/{name}", api.DeleteClub()).Methods("DELETE")
	clubRouter.HandleFunc("/update/{name}", api.UpdateClub()).Methods("PUT")
}
