package main

import (
	"defterdar-go/database"
	"defterdar-go/routes"
	"github.com/joho/godotenv"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

//@title Defterdar-go
//@version 1.0
//@description finance app for tradesman
//@termsOfService http://swagger.io/terms/

// @contact.name API SUPPORT
// @contact.url http://www.swagger.io/support
// @contact.email esezer@egetechno.com

//license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()
	database.Migrate()

	r := routes.Initroutes()

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
