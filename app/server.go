package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
type Server struct{
	DB		 *gorm.DB
	Router 	 *mux.Router	
}

type AppConfig struct{
	AppName string
	AppEnv string
	AppPort string
}

type DBConfig struct{
	DBHost string
	DBUser string
	DBPassword string
	DBName string
	DBPort string
}
func (server *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println("Selamat datang di " + appConfig.AppName)
	
	server.initializeDB(dbConfig)
	server.initializeRoutes()
	
}
func (server *Server) initializeDB(dbConfig DBConfig){
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta ", dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBPort)
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Gagal koneksi ke database server")
	}
}
func (server *Server) Run (addr string){
 	fmt.Printf("Listening port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func getEnv(key, fallback string) string{
	if value, ok := os.LookupEnv(key); ok{
		return value
	}
	return fallback
}


func Run(){
	var server = Server{}
	var appConfig = AppConfig{}
	var dbConfig = DBConfig{}
	err := godotenv.Load()
	if err != nil{
		log.Fatalf("Error on loading .env file")
	}
	appConfig.AppName = getEnv("APP_NAME", "Toko Online Golang")
	appConfig.AppEnv =getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "9999")
	
	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBUser = getEnv("DB_USER", "user")
	dbConfig.DBPassword = getEnv("DB_PASSWORD", "password")
	dbConfig.DBName = getEnv("DB_NAME", "dbname")
	dbConfig.DBPort = getEnv("DB_PORT", "5432")

	server.Initialize(appConfig, dbConfig)
	server.Run(":" + appConfig.AppPort)
}