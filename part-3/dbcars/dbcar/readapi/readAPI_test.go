package readapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOne(t *testing.T) {
	/*Set destiny file*/
	items, err := GetAllElements()
	assert.Equal(t, len(items), 12, "Error to set a destiny file")

	/*Read and load empty file*/
	item, err := GetElement("3")
	assert.NoError(t, err, "Error to read all elements from the API")
	assert.Equal(t, item.ID, "3", "Error to read all elements from the API")
	assert.Equal(t, item.Title, "Cookies", "Error to read all elements from the API")
}
