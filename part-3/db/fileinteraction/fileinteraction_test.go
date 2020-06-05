package fileinteraction

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*Run this test without file existent*/
func TestSetAndReturnDF(t *testing.T) {
	file := CreateNewFInstance("text.txt")
	assert.Equal(t, "text.txt", file.ReturnDestinyFile(), "error: destiny file was not stated")
}

func TestReadFile(t *testing.T) {
	file := CreateNewFInstance("text.txt")
	emptyMap := make(map[string]string)
	mapExample := make(map[string]string)
	mapExample["1"] = "1"
	mapExample["2"] = "2"
	mapExample["3"] = "3"
	mapExample["4"] = "4"

	emplyTables := []struct {
		key     string
		value   string
		length  int
		isUsed  bool
		element int
		err     error
	}{
		{"1", "", 0, false, 1, nil},
		{"2", "", 0, false, 2, nil},
		{"3", "", 0, false, 3, nil},
		{"4", "", 0, false, 4, nil},
	}
	nonEmplyTables := []struct {
		key     string
		value   string
		length  int
		isUsed  bool
		element int
		err     error
	}{
		{"1", "1", 4, true, 1, nil},
		{"2", "2", 4, true, 2, nil},
		{"3", "3", 4, true, 3, nil},
		{"4", "4", 4, true, 4, nil},
	}

	myfile, err := os.OpenFile(file.destinyFile, os.O_APPEND|os.O_WRONLY, os.ModePerm)
	myfile.Truncate(0)
	myfile.Seek(0, 0)
	myfile.Close()
	assert.Equal(t, nil, err, "error: file does not exist and is created")

	err = file.ReadFile(emptyMap)
	assert.Equal(t, nil, err, "error: reading empty file with empty map")

	for _, table := range emplyTables {
		_, isUsed := emptyMap[table.key]
		assert.Equal(t, table.length, len(emptyMap), "error: map have element", table.element)
		assert.Equal(t, table.isUsed, isUsed, "error: map have element", table.element)
	}

	for k, v := range mapExample {
		emptyMap[k] = v
	}

	assert.Equal(t, 4, len(emptyMap), "error: wrong assing")

	err = file.ReadFile(emptyMap)
	assert.Equal(t, nil, err, "error: reading empty file with non empty map")

	for _, table := range emplyTables {
		_, isUsed := emptyMap[table.key]
		assert.Equal(t, table.length, len(emptyMap), "error: map have element", table.element)
		assert.Equal(t, table.isUsed, isUsed, "error: map have element")
	}

	file.WriteFile(mapExample)

	err = file.ReadFile(emptyMap)
	assert.Equal(t, nil, err, "error: reading non empty file")
	for _, table := range nonEmplyTables {
		value, isUsed := emptyMap[table.key]
		assert.Equal(t, table.length, len(emptyMap), "error: map does not have 4 element", table.element)
		assert.Equal(t, table.isUsed, isUsed, "error:  element does not exist", table.element)
		assert.Equal(t, table.value, value, "error: wrong element value", table.element)
	}
}

/*var mapExample map[string]string
mapExample = make(map[string]string)

/*Set destiny file
file.SetFile("text.txt")
resultDestiny := file.ReturnDestinyFile()
assert.Equal(t, resultDestiny, "text.txt", "error to set a destiny file")

/*Read and load empty file
err := file.ReadFile(mapExample)
assert.NoError(t, err, "Error to read and load a empty file")

mapExample["firstElement"] = "zzzzzz"
mapExample["secondElement"] = "000000"

/*write 2 elements in the file
err = file.WriteFile(mapExample)
assert.NoError(t, err, "Error to write the map into the file")

mapExample["false element"] = "error"

/*Read and load non-empty file
err = file.ReadFile(mapExample)
assert.Equal(t, len(mapExample), 2, "Error to read and load a non-empty file 1")
assert.NoError(t, err, "Error to read and load a non-empty file")

mapExample["truth element"] = "its fine"

err = file.WriteFile(mapExample)
assert.NoError(t, err, "Error with the final write")*/
