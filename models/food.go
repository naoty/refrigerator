package models

import "time"

// Food is a food to be stored in the refrigerator.
type Food struct {
	// Name is the name of the food.
	Name string `json:"name"`

	// Quantity is the quantity of the food.
	Quantity Quantity `json:"quantity"`

	// UseByDate is the use-by-date of the food.
	UseByDate time.Time `json:"useByDate"`

	// BestByDate is the best-by-date of the food.
	BestByDate time.Time `json:"bestByDate"`
}

// Quantity is the quantity used to represent the quantity of foods.
type Quantity struct {
	// Value is the value of quantity.
	Value int `json:"value"`

	// Unit is the unit of quantity.
	Unit string `json:"unit"`
}
