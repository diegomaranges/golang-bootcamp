package router

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRoute(t *testing.T) {
	newRouter := NewRoute()
	assert.NotNil(t, newRouter, "error with create route (length)")

	/*expected values*/
	tables := []struct {
		name string
	}{
		{
			name: "SignIn",
		},
		{
			name: "Refresh",
		},
		{
			name: "CreateNewCar",
		},
		{
			name: "getSpecificCar",
		},
		{
			name: "DeleteCar",
		},
		{
			name: "getSpecificItem",
		},
		{
			name: "addItem",
		},
		{
			name: "updateItem",
		},
		{
			name: "deleteItem",
		},
	}

	for _, table := range tables {
		myRoute := newRouter.GetRoute(table.name)
		assert.NotNil(t, myRoute, "error: Url does not created")
	}
}
