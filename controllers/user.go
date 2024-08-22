package controllers

import (
	"defterdar-go/database"
	"defterdar-go/helpers"
	"defterdar-go/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// GetProfile godoc
// @Summary GetProfile a user
// @Description Pulls the user's own profile information.
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.User true "User"
// @Success 201 {object} models.User
// @Failure 404 {string} string "User not found."
// @Failure 400 {string} string "Invalid input."
// @Router /user/get-profile [get]
func GetProfile(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("user").(models.Claim)

	var user models.User
	if result := database.DB.First(&user, claims.UserID); result.Error != nil {
		http.Error(w, "User not found.", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		return
	}
}

// UpdateProfile godoc
// @Summary UpdateProfile a user
// @Description Update the user's profile.
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.User true "User"
// @Success 201 {string} string "User updated successfully."
// @Failure 404 {string} string "User not found."
// @Failure 400 {string} string "Invalid input."
// @Router /user/update-profile [put]
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("user").(models.Claim)

	var user models.User
	if result := database.DB.First(&user, claims.UserID); result.Error != nil {
		http.Error(w, "User not found.", http.StatusNotFound)
		return
	}

	var input models.User
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input.", http.StatusBadRequest)
		return
	}

	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Email != "" {
		user.Email = input.Email
	}

	if result := database.DB.Save(&user); result.Error != nil {
		http.Error(w, "Failed to update user.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully."})
	if err != nil {
		return
	}
}

// UpdatePassword godoc
// @Summary UpdatePassword a user
// @Description Update the user's password.
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.PasswordUpdateReq true "User"
// @Success 201 {string} string "Password updated successfully."
// @Failure 404 {string} string "User not found."
// @Failure 400 {string} string "Invalid input."
// @Router /user/update-password [put]
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("user").(models.Claim)

	var user models.User
	if result := database.DB.First(&user, claims.UserID); result.Error != nil {
		http.Error(w, "User not found.", http.StatusNotFound)
		return
	}

	var passData models.PasswordUpdateReq
	err := json.NewDecoder(r.Body).Decode(&passData)
	if err != nil {
		http.Error(w, "Invalid input.", http.StatusBadRequest)
		return
	}

	if !helpers.CheckPasswordHash(passData.OldPassword, user.Password) {
		http.Error(w, "Invalid password.", http.StatusBadRequest)
		return
	}

	hashedPassword, err := helpers.HashPassword(passData.NewPassword)
	if err != nil {
		http.Error(w, "Failed to update password.", http.StatusInternalServerError)
		return
	}

	user.Password = hashedPassword

	if result := database.DB.Save(&user); result.Error != nil {
		http.Error(w, "Failed to update password.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Password updated successfully."})
	if err != nil {
		return
	}
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by ID
// @Tags User
// @Param   id path int true "User ID"
// @Success 204 {string} string "User deleted successfully"
// @Failure 400 {string} string "Invalid user ID"
// @Failure 404 {string} string "User not found"
// @Router /users/{user_id}/delete [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["user_id"])
	if err != nil {
		http.Error(w, "Invalid user ID.", http.StatusBadRequest)
		return
	}

	if result := database.DB.Delete(&models.User{}, userID); result.Error != nil {
		http.Error(w, "Failed to delete user.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully."})
	if err != nil {
		return
	}
}

// CloseAccount godoc
// @Summary Close a user own account
// @Description Close account
// @Tags User
// @Param   id path int true "User ID"
// @Success 204 {string} string "Account closed successfully."
// @Failure 400 {string} string "Invalid user ID"
// @Failure 404 {string} string "User not found"
// @Router /users/close-account [delete]
func CloseAccount(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("user").(models.Claim)

	var user models.User
	if result := database.DB.First(&user, claims.UserID); result.Error != nil {
		http.Error(w, "User not found.", http.StatusNotFound)
		return
	}

	if result := database.DB.Delete(&user); result.Error != nil {
		http.Error(w, "Failed to close account.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]string{"message": "Account closed successfully."})
	if err != nil {
		return
	}
}
