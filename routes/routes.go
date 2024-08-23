package routes

import (
	"defterdar-go/controllers"
	"defterdar-go/middleware"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/user/register", controllers.Register).Methods("POST")
	r.HandleFunc("/user/login", controllers.Login).Methods("POST")
	r.Handle("/user/get-profile", middleware.JWTAuth(http.HandlerFunc(controllers.GetProfile))).Methods("GET")
	r.Handle("/user/update-profile", middleware.JWTAuth(http.HandlerFunc(controllers.UpdateProfile))).Methods("PUT")
	r.Handle("/user/update-password", middleware.JWTAuth(http.HandlerFunc(controllers.UpdatePassword))).Methods("PUT")
	r.Handle("/users/{user_id}/delete", middleware.JWTAuth(middleware.Authorize("admin")(http.HandlerFunc(controllers.DeleteUser)))).Methods("DELETE")
	r.Handle("/users/close-account", middleware.JWTAuth(http.HandlerFunc(controllers.CloseAccount))).Methods("DELETE")

	r.Handle("/shop", middleware.JWTAuth(middleware.Authorize("owner")(http.HandlerFunc(controllers.CreateShop)))).Methods("POST")
	r.Handle("/shop/{shop_id}", middleware.JWTAuth(middleware.Authorize("owner")(http.HandlerFunc(controllers.UpdateShop)))).Methods("PUT")
	r.Handle("/shop/{shop_id}", middleware.JWTAuth(middleware.Authorize("owner")(http.HandlerFunc(controllers.DeleteShop)))).Methods("DELETE")

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/docs/swagger.json"), // The url pointing to API definition
	))

	// Static files endpoint for serving the swagger docs
	fs := http.FileServer(http.Dir("./docs"))
	r.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", fs))

	return r
}
