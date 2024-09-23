package controllers

import (
	"defterdar-go/database"
	"defterdar-go/helpers"
	"defterdar-go/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Register godoc
// @Summary Register a new user
// @Description Register a new user with name, email, password, role (e.g. owner, employee)
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.User true "User"
// @Success 201 {string} string "User registered successfully."
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /user/register [post]
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input.", http.StatusBadRequest)
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password.", http.StatusInternalServerError)
		return
	}

	user.Password = string(hashedPass)
	user.Role = "owner"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if result := database.DBWrite.Create(&user); result.Error != nil {
		http.Error(w, "Failed to register.", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully.",
	})
	if err != nil {
		return
	}
}

// Login godoc
// @Summary Login a new user
// @Description Register a new user with name, email, password, role
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.LoginReq true "User"
// @Success 201 {string} string "token: "
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /user/login [post]
func Login(w http.ResponseWriter, r *http.Request) {
	var login models.LoginReq
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, "Invalid input.", http.StatusBadRequest)
		return
	}

	var user models.User
	if result := database.DBRead.Where("email = ?", login.Email).First(&user); result.Error != nil {
		http.Error(w, "Invalid email or password.", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		http.Error(w, "Invalid email or password.", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(720 * time.Hour)
	claims := models.Claims{
		Email:  user.Email,
		UserID: user.ID,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(helpers.GetJwtKey())
	if err != nil {
		http.Error(w, "Failed to create token.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"token": tokenStr})
	if err != nil {
		return
	}
}
