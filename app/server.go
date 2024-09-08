package app

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/khairililmi2468gmailcom/toko_online_golang/app/controllers"
)


func getEnv(key, fallback string) string{
	if value, ok := os.LookupEnv(key); ok{
		return value
	}
	return fallback
}


func Run(){
	var server = controllers.Server{}
	var appConfig = controllers.AppConfig{}
	var dbConfig = controllers.DBConfig{}
	err := godotenv.Load()
	if err != nil{
		log.Fatalf("Error on loading .env file")
	}
	appConfig.AppName = getEnv("APP_NAME", "Toko Online Golang")
	appConfig.AppEnv =getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "9999")
	appConfig.AppURL = getEnv("APP_URL", "http://localhost:9000")

	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBUser = getEnv("DB_USER", "user")
	dbConfig.DBPassword = getEnv("DB_PASSWORD", "password")
	dbConfig.DBName = getEnv("DB_NAME", "dbname")
	dbConfig.DBPort = getEnv("DB_PORT", "5432")
	dbConfig.DBDriver = getEnv("DB_DRIVER", "postgres")

	flag.Parse()
	arg := flag.Arg(0)

	if arg != "" {
		server.InitCommands(appConfig, dbConfig)
	} else {
		server.Initialize(appConfig, dbConfig)
		server.Run(":" + appConfig.AppPort)
	}

}