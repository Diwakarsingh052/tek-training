package models

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

// The table tests can certain cases like: empty array, nil array and array with actual data.
func TestCalculateTotalCost(t *testing.T) {
	tt := []struct {
		name              string
		inventories       []Inventory
		category          string
		expectedTotalCost float64
		expectedErr       error
	}{
		{
			name:              "No Items",
			inventories:       nil,
			category:          "Electronics",
			expectedTotalCost: 0,
			expectedErr:       errors.New("inventory not found"),
		},
		{
			name: "Multiple Items",
			inventories: []Inventory{
				{Category: "Electronics", CostPerItem: 10, Quantity: 1},
				{Category: "Electronics", CostPerItem: 30, Quantity: 2},
				{Category: "Electronics", CostPerItem: 10, Quantity: 1},
			},
			category:          "Electronics",
			expectedTotalCost: 80,
			expectedErr:       nil,
		},
		{
			name: "Empty Category",
			inventories: []Inventory{
				{Category: "Electronics", CostPerItem: 10, Quantity: 1},
				{Category: "Electronics", CostPerItem: 30, Quantity: 2},
				{Category: "Electronics", CostPerItem: 10, Quantity: 1},
			},
			category:          "",
			expectedTotalCost: 0,
			expectedErr:       errors.New("category doesn't exist"),
		},
	}

	for _, testCase := range tt {
		t.Run(testCase.name, func(t *testing.T) {
			actualCost, err := CalculateTotalCost(testCase.inventories, testCase.category)
			require.Equal(t, testCase.expectedTotalCost, actualCost)
			require.Equal(t, testCase.expectedErr, err)
		})
	}
}
