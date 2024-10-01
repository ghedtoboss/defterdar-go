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

	r.HandleFunc("/users/register", controllers.Register).Methods("POST") // kayıt
	r.HandleFunc("/users/login", controllers.Login).Methods("POST")       //giriş

	r.Handle("/shop", middleware.JWTAuth(middleware.Authorize("owner")(http.HandlerFunc(controllers.CreateShop)))).Methods("POST")   //shop acma
	r.Handle("/shop", middleware.JWTAuth(middleware.Authorize("owner")(http.HandlerFunc(controllers.GetShop)))).Methods("GET")       //shop bilgileri
	r.Handle("/shop", middleware.JWTAuth(middleware.Authorize("owner")(http.HandlerFunc(controllers.UpdateShop)))).Methods("PUT")    //shop güncelleme
	r.Handle("/shop", middleware.JWTAuth(middleware.Authorize("owner")(http.HandlerFunc(controllers.DeleteShop)))).Methods("DELETE") //shop silme

	r.Handle("/customers", middleware.JWTAuth(http.HandlerFunc(controllers.AddCustomer))).Methods("POST")                    //müşteri oluşturma
	r.Handle("/customers/{customer_id}", middleware.JWTAuth(http.HandlerFunc(controllers.UpdateCustomer))).Methods("PUT")    //müşteri profili güncelleme
	r.Handle("/customers/{customer_id}", middleware.JWTAuth(http.HandlerFunc(controllers.DeleteCustomer))).Methods("DELETE") //müşteri silme
	r.Handle("/customers", middleware.JWTAuth(http.HandlerFunc(controllers.GetCustomers))).Methods("GET")                    //müşterileri çekme
	r.Handle("/customers/{customer_id}", middleware.JWTAuth(http.HandlerFunc(controllers.GetCustomer))).Methods("GET")       //müşteriyi çekme

	r.Handle("/transactions", middleware.JWTAuth(middleware.Authorize("owner", "employee")(http.HandlerFunc(controllers.CreateTransaction)))).Methods("POST")                    //işlem oluşturma
	r.Handle("/transactions/{transaction_id}", middleware.JWTAuth(middleware.Authorize("owner", "employee")(http.HandlerFunc(controllers.UpdateTransaction)))).Methods("PUT")    //işlem güncelleme
	r.Handle("/transactions", middleware.JWTAuth(middleware.Authorize("owner", "employee")(http.HandlerFunc(controllers.GetTransactions)))).Methods("GET")                       //işlemleri çekme
	r.Handle("/transactions/{transaction_id}", middleware.JWTAuth(middleware.Authorize("owner", "employee")(http.HandlerFunc(controllers.GetTransaction)))).Methods("GET")       //işlem çekme
	r.Handle("/transactions/{transaction_id}", middleware.JWTAuth(middleware.Authorize("owner", "employee")(http.HandlerFunc(controllers.DeleteTransaction)))).Methods("DELETE") //işlem silme

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/docs/swagger.json"),
	))

	// Static files endpoint for serving the swagger docs
	fs := http.FileServer(http.Dir("./docs"))
	r.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", fs))

	return r
}
