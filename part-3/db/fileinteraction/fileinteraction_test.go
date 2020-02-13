package fileinteraction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*Run this test without file existent*/
func TestOne(t *testing.T) {
	file := CreateNewFInstance()
	var mapExample map[string]string
	mapExample = make(map[string]string)

	/*Set destiny file*/
	file.SetFile("text.txt")
	resultDestiny := file.ReturnDestinyFile()
	assert.Equal(t, resultDestiny, "text.txt", "error to set a destiny file")

	/*Read and load empty file*/
	err := file.ReadFile(mapExample)
	assert.NoError(t, err, "Error to read and load a empty file")

	mapExample["firstElement"] = "zzzzzz"
	mapExample["secondElement"] = "000000"

	/*write 2 elements in the file*/
	err = file.WriteFile(mapExample)
	assert.NoError(t, err, "Error to write the map into the file")

	mapExample["false element"] = "error"

	/*Read and load non-empty file*/
	err = file.ReadFile(mapExample)
	assert.Equal(t, len(mapExample), 2, "Error to read and load a non-empty file 1")
	assert.NoError(t, err, "Error to read and load a non-empty file")

	mapExample["truth element"] = "its fine"

	err = file.WriteFile(mapExample)
	assert.NoError(t, err, "Error with the final write")
}
