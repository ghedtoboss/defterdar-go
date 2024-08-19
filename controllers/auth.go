package controllers

import (
	"defterdar-go/database"
	"defterdar-go/models"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// Register godoc
// @Summary Register a new user
// @Description Register a new user with name, email, password, role
// @Tags User
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

	if user.Role == "" {
		http.Error(w, "Role is required.", http.StatusBadRequest)
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password.", http.StatusInternalServerError)
		return
	}

	user.Password = string(hashedPass)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if result := database.DB.Create(&user); result.Error != nil {
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
// @Tags User
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
	if result := database.DB.Where("email = ?", login.Email).First(&user); result.Error != nil {
		http.Error(w, "Invalid email or password.", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		http.Error(w, "Invalid email or password.", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(720 * time.Hour)
	claims := models.Claim{
		Email:  user.Email,
		UserID: user.ID,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte("jwtKey"))
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
