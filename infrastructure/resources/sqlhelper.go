package resources

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nelsonlpco/transactions/infrastructure"
	"github.com/sirupsen/logrus"
)

func CreateDatabase(env *infrastructure.Environment) {
	logrus.WithFields(logrus.Fields{"drive": env.GetSqlDrive(), "connection": env.GetConnectionString()}).Info("connection to database...")
	db, err := sql.Open(env.GetSqlDrive(), env.GetConnectionString())
	if err != nil {
		logrus.Fatal("error on connect to database: ", err)
	}

	defer func() {
		log.Println("encerrando database")
		err := db.Close()
		if err != nil {
			log.Fatalf("erro ao encerrar database: %v", err)
		}
	}()

	workDir, err := os.Getwd()
	if err != nil {
		log.Panicf("error on get workdir: %v", err)
	}

	data, err := os.ReadFile(fmt.Sprintf("%v/infrastructure/resources/schemas/transactions_schame.sql", workDir))
	if err != nil {
		log.Panicf("error on open file: %v", err)
	}

	_, err = db.Exec(string(data))
	if err != nil {
		log.Panicf("error on exec sql: %v", err)
	}

}
