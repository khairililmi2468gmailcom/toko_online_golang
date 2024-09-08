package controllers

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/khairililmi2468gmailcom/toko_online_golang/app/models"
	"github.com/khairililmi2468gmailcom/toko_online_golang/database/seeders"
	"github.com/urfave/cli"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type PageLink struct {
	Page          int32
	Url           string
	IsCurrentPage bool
}

type PaginationLinks struct {
	CurrentPage string
	NextPage    string
	PrevPage    string
	TotalRows   int32
	TotalPages  int32
	Links       []PageLink
}

type PaginationParams struct {
	Path        string
	TotalRows   int32
	PerPage     int32
	CurrentPage int32
}

type Server struct{
	DB		 *gorm.DB
	Router 	 *mux.Router
    AppConfig *AppConfig

}

type AppConfig struct{
	AppName string
	AppEnv string
	AppPort string
    AppURL  string

}

type DBConfig struct{
	DBHost string
	DBUser string
	DBPassword string
	DBName string
	DBPort string
    DBDriver   string

}
func (server *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println("Selamat datang di " + appConfig.AppName)

	server.initializeDB(dbConfig)
    server.initializeAppConfig(appConfig)
	server.initializeRoutes()
	// seeders.DBSeed(server.DB)

}
func (server *Server) initializeAppConfig(appConfig AppConfig) {
	server.AppConfig = &appConfig
}

func (server *Server) initializeDB(dbConfig DBConfig){
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta ", dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBPort)
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Gagal koneksi ke database server")
	}

}
func (server *Server) dbMigrate() {
	for _, model := range models.RegisterModels() {
		err := server.DB.Debug().AutoMigrate(model.Model)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Database migrated successfully.")
}
func (server *Server) Run (addr string){
 	fmt.Printf("Listening port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func (server *Server) InitCommands(config AppConfig, dbConfig DBConfig) {
	server.initializeDB(dbConfig)

	cmdApp := cli.NewApp()
	cmdApp.Commands = []cli.Command{
		{
			Name: "db:migrate",
			Action: func(c *cli.Context) error {
				server.dbMigrate()
				return nil
			},
		},
		{
			Name: "db:seed",
			Action: func(c *cli.Context) error {
				err := seeders.DBSeed(server.DB)
				if err != nil {
					log.Fatal(err)
				}

				return nil
			},
		},
	}

	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPaginationLinks(config *AppConfig, params PaginationParams) (PaginationLinks, error) {
	var links []PageLink

	totalPages := int32(math.Ceil(float64(params.TotalRows) / float64(params.PerPage)))

	for i := 1; int32(i) <= totalPages; i++ {
		links = append(links, PageLink{
			Page:          int32(i),
			Url:           fmt.Sprintf("%s/%s?page=%s", config.AppURL, params.Path, fmt.Sprint(i)),
			IsCurrentPage: int32(i) == params.CurrentPage,
		})
	}

	var nextPage int32
	var prevPage int32

	prevPage = 1
	nextPage = totalPages

	if params.CurrentPage > 2 {
		prevPage = params.CurrentPage - 1
	}

	if params.CurrentPage < totalPages {
		nextPage = params.CurrentPage + 1
	}

	return PaginationLinks{
		CurrentPage: fmt.Sprintf("%s/%s?page=%s", config.AppURL, params.Path, fmt.Sprint(params.CurrentPage)),
		NextPage:    fmt.Sprintf("%s/%s?page=%s", config.AppURL, params.Path, fmt.Sprint(nextPage)),
		PrevPage:    fmt.Sprintf("%s/%s?page=%s", config.AppURL, params.Path, fmt.Sprint(prevPage)),
		TotalRows:   params.TotalRows,
		TotalPages:  totalPages,
		Links:       links,
	}, nil
}