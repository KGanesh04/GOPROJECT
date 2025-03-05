package handlers

import (
	"go-mysql-api/database"
	"go-mysql-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ✅ Get All Users
func GetUsers(c *gin.Context) {
	var users []models.User
	err := database.DB.Select(&users, "SELECT * FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// ✅ Get User by ID
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	err := database.DB.Get(&user, "SELECT * FROM users WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// ✅ Create User
func CreateUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result, err := database.DB.Exec("INSERT INTO users (name, email) VALUES (?, ?)", newUser.Name, newUser.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	id, _ := result.LastInsertId()
	newUser.ID = int(id)

	c.JSON(http.StatusCreated, newUser)
}

// ✅ Update User
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result, err := database.DB.Exec("UPDATE users SET name=?, email=? WHERE id=?", updatedUser.Name, updatedUser.Email, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// ✅ Delete User
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	result, err := database.DB.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
