package http

import (
	"goexample/internal/user/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(u usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: u}
}

// CreateUser - Handler para criar um novo usuário
func (h *UserHandler) CreateUser(c *gin.Context) {
	var input struct {
		Name string `json:"name"`
	}

	// Decodifica o JSON recebido na requisição
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao decodificar JSON"})
		return
	}

	// Chama o caso de uso para criar um novo usuário
	user, err := h.userUsecase.CreateUser(input.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário"})
		return
	}

	// Retorna a resposta com o usuário criado
	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuário criado com sucesso",
		"user":    user,
	})
}

// ListUsers - Handler para listar todos os usuários
func (h *UserHandler) ListUsers(c *gin.Context) {
	// Chama o caso de uso para listar os usuários
	users, err := h.userUsecase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar usuários"})
		return
	}

	// Retorna a lista de usuários
	c.JSON(http.StatusOK, users)
}
