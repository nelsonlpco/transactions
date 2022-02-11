package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nelsonlpco/transactions/api/adapter/rest/controller"
	"github.com/nelsonlpco/transactions/api/adapter/rest/middlewares"
	"github.com/nelsonlpco/transactions/infrastructure"
	"github.com/nelsonlpco/transactions/shared/dependencies"
	"github.com/sirupsen/logrus"
)

type Server struct {
	router              *mux.Router
	environment         *infrastructure.Environment
	dependencyContainer *dependencies.DependencyContainer
	controllers         *controller.Controllers
}

func NewServer(
	env *infrastructure.Environment,
	dependencyContainer *dependencies.DependencyContainer,
) *Server {
	return &Server{
		router:              mux.NewRouter(),
		environment:         env,
		dependencyContainer: dependencyContainer,
		controllers:         controller.NewControllers(dependencyContainer.Services),
	}
}

func (s *Server) registerRoutes() {
	logrus.Info("registering routes...")
	s.router.HandleFunc("/accounts", s.controllers.AccountController.CreatAccount).Methods("POST")
	s.router.HandleFunc("/accounts/{accountId}", s.controllers.AccountController.GetAccount).Methods("GET")
	s.router.HandleFunc("/transactions", s.controllers.TransactionController.CreateTransaction).Methods("POST")
}

func (s *Server) setMiddlewares() {
	logrus.Info("setting middlewares...")
	s.router.Use(middlewares.DefaultHeadersMidleware)
}

func (s *Server) InitConfigs() {
	s.setMiddlewares()
	s.registerRoutes()
}

func (s *Server) Start() {
	s.InitConfigs()

	logrus.WithFields(logrus.Fields{"server_port": s.environment.GetServerPort()}).Info("serving listening on http://localhost:", s.environment.GetServerPort())

	if err := http.ListenAndServe(fmt.Sprintf(":%v", s.environment.GetServerPort()), s.router); err != nil {
		logrus.Panic("Error on start server %v", err)
	}
}
