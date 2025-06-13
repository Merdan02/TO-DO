package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"todo-app/internal/models"
	"todo-app/internal/service"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		service: userService,
	}
}

type UserResponse struct {
	ID       int       `json:"id"`
	Username string    `json:"username"`
	Role     string    `json:"role"`
	Created  time.Time `json:"created"`
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.UserModel
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.CreateUser(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
		Created:  user.CreatedAt,
	})
}
