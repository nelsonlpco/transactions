package main

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nelsonlpco/transactions/api/adapter/rest"
	"github.com/nelsonlpco/transactions/infrastructure"
	"github.com/nelsonlpco/transactions/infrastructure/db_manager"
	"github.com/nelsonlpco/transactions/infrastructure/resources"
	"github.com/nelsonlpco/transactions/shared/dependencies"
	"github.com/sirupsen/logrus"
)

func main() {
	environment := infrastructure.NewEnvironment()
	configLogrus(environment)

	resources.CreateDatabase(environment)

	dbManager := openDatabase(environment)

	defer dbManager.Finish()

	server := rest.NewServer(environment, dependencies.NewDependencyContainer(dbManager))

	server.Start()
}

func openDatabase(env *infrastructure.Environment) *db_manager.DBManager {
	db, err := sql.Open(env.GetSqlDrive(), env.GetConnectionString())
	if err != nil {
		logrus.Fatal("error on connect to database: ", err)
	}

	dbManager := db_manager.NewDBManager(db, env.GetTTL())

	return dbManager
}

func configLogrus(env *infrastructure.Environment) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.Level(env.GetLogLevel()))
}
