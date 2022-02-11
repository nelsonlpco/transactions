package db_manager

import (
	"database/sql"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type DBManager struct {
	db  *sql.DB
	ttl time.Duration
}

func NewDBManager(db *sql.DB, ttl time.Duration) *DBManager {
	return &DBManager{
		db:  db,
		ttl: ttl,
	}
}

func (d *DBManager) ExecSqlFromResource(resourcePath string) {
	resource, err := os.ReadFile(resourcePath)
	if err != nil {
		logrus.Panic("error on read sql resource: ", err)
	}

	sqlQuery := string(resource)

	_, err = d.db.Exec(sqlQuery)
	if err != nil {
		logrus.Panic("error on exec sql resource: ", err)
	}
}

func (s *DBManager) Finish() {
	logrus.Info("Disconnecting from database")
	err := s.db.Close()
	if err != nil {
		logrus.Error("error on disconnect to database: ", err)
	}
}

func (s *DBManager) GetDB() *sql.DB {
	return s.db
}

func (s *DBManager) GetTTL() time.Duration {
	return s.ttl
}
