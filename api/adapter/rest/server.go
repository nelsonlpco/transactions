package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nelsonlpco/transactions/api/adapter/rest/config/fabric"
	"github.com/nelsonlpco/transactions/api/adapter/rest/controller"
	"github.com/nelsonlpco/transactions/api/adapter/rest/middlewares"
)

type Server struct {
	router            *mux.Router
	accountController *controller.AccountController
}

func NewServer() *Server {
	return &Server{
		router:            mux.NewRouter(),
		accountController: fabric.MakeAccountController(),
	}
}

func (s *Server) registerRoutes() {
	s.router.HandleFunc("/accounts", s.accountController.CreatAccount).Methods("POST")
	s.router.HandleFunc("/accounts/{accountId}", s.accountController.GetAccount).Methods("GET")

}

func (s *Server) setMiddlewares() {
	s.router.Use(middlewares.DefaultHeadersMidleware)
}

func (s *Server) Start() {
	fmt.Println("running on")

	s.setMiddlewares()
	s.registerRoutes()

	log.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", s.router); err != nil {
		log.Panicf("Error on start server %v", err)
	}

}
