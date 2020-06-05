package readapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllElements(t *testing.T) {
	items, err := GetAllElements()
	assert.Equal(t, len(items), 12, "Error to set a destiny file")
	assert.NoError(t, err, "Error to set a destiny file")
}

func TestGetElement(t *testing.T) {
	items := []struct {
		id    string
		title string
		price string
	}{
		{
			id:    "1",
			title: "Bannana",
			price: "2.50",
		},
		{
			id:    "2",
			title: "Apple",
			price: "3.20",
		},
		{
			id:    "3",
			title: "Cookies",
			price: "10.40",
		},
		{
			id:    "4",
			title: "Noodles",
			price: "23.50",
		},
		{
			id:    "5",
			title: "Olive Oil",
			price: "13.00",
		},
		{
			id:    "6",
			title: "Water",
			price: "0.50",
		},
		{
			id:    "7",
			title: "Beer",
			price: "1.50",
		},
		{
			id:    "8",
			title: "Vodka",
			price: "10.50",
		},
		{
			id:    "9",
			title: "Bread",
			price: "0.20",
		},
		{
			id:    "10",
			title: "Grapes",
			price: "0.50",
		},
		{
			id:    "11",
			title: "Rice",
			price: "3.50",
		},
		{
			id:    "12",
			title: "Pizza",
			price: "13.10",
		},
	}

	for _, item := range items {
		itemReturned, err := GetElement(item.id)
		assert.NoError(t, err, "Error to read a element from the API")
		assert.Equal(t, item.title, itemReturned.Title, "Error to read a element from the API")
		assert.Equal(t, item.price, itemReturned.Price, "Error to read a element from the API")

	}

}
