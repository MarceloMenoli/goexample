package http

import (
	"goexample/internal/drink/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DrinkHandler struct {
	drinkUsecase usecase.DrinkUsecase
}

func NewDrinkHandler(u usecase.DrinkUsecase) *DrinkHandler {
	return &DrinkHandler{drinkUsecase: u}
}

// CreateDrink - Handler para criar um novo drink
func (h *DrinkHandler) CreateDrink(c *gin.Context) {
	var input struct {
		Name        string `json:"name"`
		Ingredients string `json:"ingredients"`
		Description string `json:"description"`
		IsAlcoholic bool   `json:"is_alcoholic"`
		Rating      int    `json:"rating"`
	}

	// Decodifica o JSON recebido na requisição
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao decodificar JSON"})
		return
	}

	// Chama o caso de uso para criar um novo drink
	drink, err := h.drinkUsecase.CreateDrink(input.Name, input.Ingredients, input.Description, input.IsAlcoholic, input.Rating)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar drink"})
		return
	}

	// Retorna o drink criado com sucesso
	c.JSON(http.StatusCreated, gin.H{
		"message": "Drink criado com sucesso",
		"drink":   drink,
	})
}

// ListDrinks - Handler para listar todos os drinks
func (h *DrinkHandler) ListDrinks(c *gin.Context) {
	drinks, err := h.drinkUsecase.GetAllDrinks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar drinks"})
		return
	}

	// Retorna a lista de drinks
	c.JSON(http.StatusOK, drinks)
}

// DeleteDrink - Handler para deletar um drink por ID
func (h *DrinkHandler) DeleteDrink(c *gin.Context) {
	var input struct {
		ID uint `json:"id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao decodificar JSON"})
		return
	}

	if err := h.drinkUsecase.DeleteDrinkByID(input.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar drink"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Drink deletado com sucesso"})
}
