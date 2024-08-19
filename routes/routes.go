package routes

import (
	"defterdar-go/controllers"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func Initroutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/user/register", controllers.Register).Methods("POST")
	r.HandleFunc("/user/login", controllers.Login).Methods("POST")

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/docs/swagger.json"), // The url pointing to API definition
	))

	// Static files endpoint for serving the swagger docs
	fs := http.FileServer(http.Dir("./docs"))
	r.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", fs))

	return r
}
