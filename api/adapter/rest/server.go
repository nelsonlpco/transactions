package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nelsonlpco/transactions/api/adapter/rest/core/dependencies"
	"github.com/nelsonlpco/transactions/api/adapter/rest/middlewares"
	"github.com/nelsonlpco/transactions/infrastructure"
	"github.com/nelsonlpco/transactions/infrastructure/db_manager"
	"github.com/sirupsen/logrus"
)

type Server struct {
	dbManager           *db_manager.DBManager
	router              *mux.Router
	environment         *infrastructure.Environment
	dependencyContainer *dependencies.DependencyContainer
}

func NewServer(
	env *infrastructure.Environment,
	dependencyContainer *dependencies.DependencyContainer,
	dbManager *db_manager.DBManager,
) *Server {
	return &Server{
		router:              mux.NewRouter(),
		environment:         env,
		dependencyContainer: dependencyContainer,
		dbManager:           dbManager,
	}
}

func (s *Server) registerRoutes() {
	logrus.Info("registering routes...")
	s.router.HandleFunc("/accounts", s.dependencyContainer.Controllers.AccountController.CreatAccount).Methods("POST")
	s.router.HandleFunc("/accounts/{accountId}", s.dependencyContainer.Controllers.AccountController.GetAccount).Methods("GET")
	s.router.HandleFunc("/transactions", s.dependencyContainer.Controllers.TransactionController.CreateTransaction).Methods("POST")
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

func (s *Server) Finish() {
	logrus.Info("Finishing server...")
	s.dbManager.GetDB().Close()
}
