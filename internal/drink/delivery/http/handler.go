package http

import (
	"goexample/internal/drink/usecase"
	"net/http"
	"strconv"

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

	contentType := c.ContentType()
	if contentType != "multipart/form-data" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tipo de conteúdo deve ser multipart/form-data"})
		return
	}

	name := c.PostForm("name")
	ingredients := c.PostForm("ingredients")
	description := c.PostForm("description")
	isAlcoholicStr := c.PostForm("is_alcoholic")
	ratingStr := c.PostForm("rating")

	if name == "" || ingredients == "" || description == "" || isAlcoholicStr == "" || ratingStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos os campos são obrigatórios"})
		return
	}

	// Converter isAlcoholic para bool
	isAlcoholic, err := strconv.ParseBool(isAlcoholicStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Valor inválido para is_alcoholic"})
		return
	}

	// Converter rating para int
	rating, err := strconv.Atoi(ratingStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Valor inválido para rating"})
		return
	}

	// Obter o arquivo de imagem
	fileHeader, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Imagem é obrigatória"})
		return
	}

	// Abrir o arquivo
	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao abrir o arquivo de imagem"})
		return
	}
	defer file.Close()

	// Chamar o caso de uso para criar um novo drink
	drink, err := h.drinkUsecase.CreateDrink(name, ingredients, description, isAlcoholic, rating, file, fileHeader.Filename)
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
