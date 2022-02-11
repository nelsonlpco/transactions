package infrastructure

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Environment struct {
	serverPort       int
	sqlDrive         string
	connectionString string
	logLevel         string
	ttl              time.Duration
}

func NewEnvironment() *Environment {
	envFile := getEnvOrDefault("ENVFILE", ".env")
	err := godotenv.Load(envFile)
	if err != nil {
		log.Println(fmt.Sprintf("envfile %v not found", envFile))
	}

	serverPort, err := strconv.Atoi(getEnvOrDefault("SERVER_PORT", "5555"))
	if err != nil {
		serverPort = 5555
	}

	timeoutSeconds, err := strconv.Atoi(getEnvOrDefault("TIMEOUT_SECODS", "1"))
	if err != nil {
		timeoutSeconds = 1
	}
	timeoutDuration := time.Duration(timeoutSeconds) * time.Second

	return &Environment{
		serverPort:       serverPort,
		sqlDrive:         getEnvOrDefault("SQL_DRIVER", "mysql"),
		connectionString: getEnvOrDefault("CONNECTION_STRING", ""),
		logLevel:         getEnvOrDefault("LOG_LEVEL", "5"),
		ttl:              timeoutDuration,
	}
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
	level, err := strconv.Atoi(e.logLevel)
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

func (e *Environment) GetSqlDrive() string {
	return e.sqlDrive
}

func (e *Environment) GetTTL() time.Duration {
	return e.ttl
}
