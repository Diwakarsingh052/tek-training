package models

import (
	"context"
	"errors"
)

// Define the function CreatInventory, which belongs to the struct 'Service'.
// This function takes in 3 parameters: a context `ctx` of type `Context`, `ni` of type `NewInventory`, and `userId` of type `uint`.
// This function will return an `Inventory` and an `error`.

func (s *Service) CreatInventory(ctx context.Context, ni NewInventory, userId uint) (Inventory, error) {
	// Create a new 'Inventory' struct named 'inv'.
	// Initialize it with parameters from the 'NewInventory' struct and the `userId` passed to the function.
	inv := Inventory{
		ItemName:    ni.ItemName,
		Quantity:    ni.Quantity,
		Category:    ni.Category,
		UserId:      userId,
		CostPerItem: ni.CostPerItem,
	}

	// Create a new database transaction using `ctx` as the context.
	// Within this transaction, create a new row in the database for the 'inv' struct.
	tx := s.db.WithContext(ctx).Create(&inv)

	// If there's an error with the database transaction.
	if tx.Error != nil {
		// Return an empty 'Inventory' struct and the error.
		return Inventory{}, tx.Error
	}

	// If there was no error with the database transaction, return 'inv' and nil as the error.
	return inv, nil
}

func (s *Service) ViewInventory(ctx context.Context, userId string) ([]Inventory, float64, error) {
	var inv = make([]Inventory, 0, 10)
	tx := s.db.WithContext(ctx).Where("user_id = ?", userId)
	err := tx.Find(&inv).Error
	if err != nil {
		return nil, 0, err
	}

	totalCost, err := CalculateTotalCost(inv, "shirts")
	if err != nil {
		return nil, 0, err
	}
	return inv, totalCost, nil

}

func CalculateTotalCost(inventories []Inventory, category string) (float64, error) {
	if category == "" {
		return 0, errors.New("category doesn't exist")
	}
	// Compute the total cost
	var totalCost float64
	for _, inventory := range inventories {
		totalCost += inventory.CostPerItem * float64(inventory.Quantity)
	}

	return totalCost, nil
}
