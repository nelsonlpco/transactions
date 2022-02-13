package core

import (
	"database/sql"
	"flag"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/nelsonlpco/transactions/api/adapter/rest"
	"github.com/nelsonlpco/transactions/api/adapter/rest/core/dependencies"
	"github.com/nelsonlpco/transactions/infrastructure"
	"github.com/nelsonlpco/transactions/infrastructure/db_manager"
	"github.com/nelsonlpco/transactions/infrastructure/resources"
	"github.com/sirupsen/logrus"
)

func openDatabase(env *infrastructure.Environment, dbChan chan *sql.DB) {
	var db *sql.DB
	var err error
	for {
		db, err = sql.Open(env.GetSqlDrive(), env.GetConnectionString())
		if err != nil {
			logrus.Error("error on connect to database: ", err)
		}

		err := db.Ping()
		if err == nil {
			break
		}

		time.Sleep(5 * time.Second)
	}

	dbChan <- db
}

func configLogrus(env *infrastructure.Environment) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.Level(env.GetLogLevel()))
}

func cli() (bool, bool) {
	dumpDbRef := flag.Bool("d", false, "initialize database")
	seedOperationsRef := flag.Bool("s", false, "seed default operation types")
	flag.Parse()

	return *dumpDbRef, *seedOperationsRef
}

func Bootstrap() *rest.Server {
	dumpDB, seedOperations := cli()

	environment := infrastructure.NewEnvironment()
	configLogrus(environment)
	mysql.SetLogger(logrus.StandardLogger())

	dbCh := make(chan *sql.DB, 1)

	go openDatabase(environment, dbCh)

	db := <-dbCh

	dbManager := db_manager.NewDBManager(db, environment.GetTTL())

	if dumpDB {
		resources.CreateDatabase(dbManager.GetDB(), environment.GetSqlSchema())
		if seedOperations {
			resources.SeedOperationsType(dbManager.GetDB())
		}
	}

	return rest.NewServer(environment, dependencies.NewDependencyContainer(dbManager), dbManager)
}
