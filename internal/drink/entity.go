package drink

type Drink struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
	Description string `json:"description"`
	IsAlcoholic bool   `json:"is_alcoholic"`
	Rating      int    `json:"rating"`
	ImageURL    string `json:"image_url"`
}
