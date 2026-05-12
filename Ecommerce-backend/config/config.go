package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/lpernett/godotenv"
)

var configurations *Config

type DBConfig struct {
	Host         string
	Port         int
	Name         string
	User         string
	Password     string
	EnableSSMODE bool
}

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
	DB           *DBConfig
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to  load the env Variable")
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service name is required")
		os.Exit(1)

	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {

		fmt.Println("Http port is required")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println("Port must be Number ")
		os.Exit(1)
	}
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {

		fmt.Println("Jwt secrect key is required")
		os.Exit(1)
	}

	dbhost := os.Getenv("DB_HOST")
	if dbhost == "" {

		fmt.Println("DB host is required")
		os.Exit(1)
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {

		fmt.Println("Databse port is required")
		os.Exit(1)
	}

	dbPrt, err := strconv.ParseInt(dbPort, 10, 64)
	if err != nil {
		fmt.Println("DB port must be Number ")
		os.Exit(1)
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("DB name is Required ")
		os.Exit(1)
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		fmt.Println("DB user is Required ")
		os.Exit(1)
	}
	dbPass := os.Getenv("DB_PASSWORD")
	if dbPass == "" {
		fmt.Println("DB password is Required ")
		os.Exit(1)
	}
	enableSslMode := os.Getenv("DB_ENABLE_SSL_MODE")

	var enbleSslMode bool
	if enableSslMode != "" {
		enbleSslMode, err = strconv.ParseBool(enableSslMode)
		if err != nil {
			fmt.Println("Invalid enable ssl mode type", err)
			os.Exit(1)
		}
	}

	dbConfig := &DBConfig{

		Host:         dbhost,
		Port:         int(dbPrt),
		Name:         dbName,
		User:         dbUser,
		Password:     dbPass,
		EnableSSMODE: enbleSslMode,
	}

	configurations = &Config{

		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     int(port),
		JwtSecretKey: jwtSecretKey,
		DB:           dbConfig,
	}
}
func GetConfig() *Config {
	if configurations == nil {
		loadConfig()

	}

	return configurations
}
