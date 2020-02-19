package fileinteraction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*Run this test without file existent*/
func TestOne(t *testing.T) {
	file := CreateNewFInstance()
	var mapExample map[string]*Items
	mapExample = make(map[string]*Items)

	/*Set destiny file*/
	file.SetFile("text")
	resultDestiny := file.ReturnDestinyFile()
	assert.Equal(t, resultDestiny, "text.json", "error to set a destiny file")

	/*Read and load empty file*/
	err := file.ReadFile(mapExample)
	assert.NoError(t, err, "Error to read and load a empty file")

	/*write 2 elements in the file*/
	err = file.WriteFile(mapExample)
	assert.NoError(t, err, "Error to write the map into the file")

	/*Read and load non-empty file*/
	err = file.ReadFile(mapExample)
	assert.NoError(t, err, "Error to read and load a non-empty file")

	err = file.WriteFile(mapExample)
	assert.NoError(t, err, "Error with the final write")
}
