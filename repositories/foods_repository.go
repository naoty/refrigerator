package repositories

import "github.com/naoty/refrigerator/models"

// FoodsRepository is a repository for Food.
type FoodsRepository struct {
}

// GetFoods returns foods.
func (repo FoodsRepository) GetFoods() []models.Food {
	// TODO: Get data from database
	return []models.Food{
		models.Food{Name: "Apple", Quantity: models.Quantity{Value: 1}},
		models.Food{Name: "Orange", Quantity: models.Quantity{Value: 2}},
	}
}
