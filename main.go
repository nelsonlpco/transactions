package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nelsonlpco/transactions/api/adapter/rest/core"
)

func main() {
	server := core.Bootstrap()
	server.Start()

	defer server.Finish()
}
