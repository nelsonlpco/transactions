package infrastructure

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Environment struct {
	serverPort        int
	sqlDrive          string
	connectionString  string
	logLevel          uint32
	ttl               time.Duration
	connectionTimeout time.Duration
	schemasPath       string
}

func NewEnvironment() *Environment {
	envFile := getEnvOrDefault("ENVFILE", ".env")
	err := godotenv.Load(envFile)
	if err != nil {
		log.Println(fmt.Sprintf("envfile %v not found", envFile))
	}

	return &Environment{
		serverPort:        formatServerPort(),
		sqlDrive:          getEnvOrDefault("SQL_DRIVER", "sqlite3"),
		connectionString:  getEnvOrDefault("CONNECTION_STRING", "file:./transaction.db"),
		logLevel:          formatLogLevel(),
		ttl:               formatTTL(),
		connectionTimeout: formatConnectionTimeout(),
		schemasPath:       getEnvOrDefault("SCHEMAS_PATH", "./schemas/v1_schema.sql"),
	}
}

func formatServerPort() int {
	serverPort, err := strconv.Atoi(getEnvOrDefault("SERVER_PORT", "5555"))
	if err != nil {
		serverPort = 5555
	}

	return serverPort
}

func formatTTL() time.Duration {
	timeoutSeconds, err := strconv.Atoi(getEnvOrDefault("TIMEOUT_SECODS", "1"))
	if err != nil {
		timeoutSeconds = 1
	}
	return time.Duration(timeoutSeconds) * time.Second
}

func formatLogLevel() uint32 {
	logLevel := getEnvOrDefault("LOG_LEVEL", "5")
	level, err := strconv.Atoi(logLevel)
	if err != nil {
		return 5
	}
	if level < 1 {
		return 1
	}
	if level > 6 {
		return 6
	}

	return uint32(level)
}

func formatConnectionTimeout() time.Duration {
	value := getEnvOrDefault("CONNECTION_TIMEOUT", "0")
	timeout, err := strconv.ParseInt(value, 10, 8)
	if err != nil {
		logrus.Panic(err)
	}

	return time.Duration(timeout) * time.Second

}

func getEnvOrDefault(envName string, defaultValue string) string {
	env, ok := os.LookupEnv(envName)
	if !ok {
		return defaultValue
	}
	return env
}

func (e *Environment) GetServerPort() int {
	return e.serverPort
}

func (e *Environment) GetConnectionString() string {
	return e.connectionString
}

func (e *Environment) GetLogLevel() uint32 {
	return e.logLevel
}

func (e *Environment) GetSqlDrive() string {
	return e.sqlDrive
}

func (e *Environment) GetTTL() time.Duration {
	return e.ttl
}

func (e *Environment) GetConnectionTimeout() time.Duration {
	return e.connectionTimeout
}

func (e *Environment) GetSqlSchema() string {
	return e.schemasPath
}
