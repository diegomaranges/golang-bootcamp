package fileinteraction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOne(t *testing.T) {
	file := CreateNewFInstance("")
	var mapExample map[string]*Items
	mapExample = make(map[string]*Items)
	var quantity int

	/*Set destiny file*/
	resultDestiny := file.ReturnDestinyFile()
	assert.Equal(t, resultDestiny, "db.json", "Error to set a destiny file")

	/*Read and load empty file*/
	err := file.ReadFile(mapExample)
	assert.Equal(t, len(mapExample), 3, "Error to read and load a non-empty file")
	assert.NoError(t, err, "Error to read and load a non-empty file")
	quantity = mapExample["1"].Quantity

	mapExample["1"].Quantity++

	/*write 2 elements in the file*/
	err = file.WriteFile(mapExample)
	assert.NoError(t, err, "Error to write the map into the file")

	/*Read and load non-empty file*/
	err = file.ReadFile(mapExample)
	assert.Equal(t, len(mapExample), 3, "Error to read and load a non-empty file")
	assert.Equal(t, mapExample["1"].Quantity, quantity+1, "Error to write the map into the file")
	assert.NoError(t, err, "Error to read and load a non-empty file")

	err = file.WriteFile(mapExample)
	assert.NoError(t, err, "Error with the final write")
}
