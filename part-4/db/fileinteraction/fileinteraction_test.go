package fileinteraction

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const destinyFile = "cars/dbtest.json"

func TestCreateFile(t *testing.T) {
	os.Remove(destinyFile)
	testInstance := CreateNewFInstance("test")
	err := testInstance.CreateFile()
	assert.NoError(t, err, "error: return a error when try create a new file non existent")

	err = testInstance.CreateFile()
	assert.Error(t, err, "error: no return a error when try create a new file and this already exist")
}

func TestReadFile(t *testing.T) {
	testInstance := CreateNewFInstance("testRead")
	testMap := make(map[string]*Items)

	err := testInstance.ReadFile(testMap)
	assert.NoError(t, err, "error: unexpected error tring read a test file")
	assert.Equal(t, 2, len(testMap), "error: ReadFile does not load all elements")
}

func TestWriteFile(t *testing.T) {
	testInstance := CreateNewFInstance("testRead")
	testMap := make(map[string]*Items)

	err := testInstance.ReadFile(testMap)
	assert.NoError(t, err, "error: unexpected error tring read a test file in write test")
	assert.Equal(t, 2, len(testMap), "error: ReadFile does not load all elements in write test")

	testInstanceWrite := CreateNewFInstance("testWrite")

	err = testInstanceWrite.WriteFile(testMap)
	assert.NoError(t, err, "error: unexpected error tring write a test file")
	assert.Equal(t, 2, len(testMap), "error: WriteFile modify quantity of items")
}

func TestDeleteFile(t *testing.T) {
	testInstance := CreateNewFInstance("testDelete")
	err := testInstance.CreateFile()
	assert.NoError(t, err, "error: return a error when try create a new file non existent in delete test")

	err = testInstance.DeleteFile()
	assert.NoError(t, err, "error: does not delete the test file")

	err = testInstance.DeleteFile()
	assert.Error(t, err, "error: does not error occurred when try delete the non existent file")
}

func TestReturnDestinyFile(t *testing.T) {
	testInstance := CreateNewFInstance("test")
	destiny := testInstance.ReturnDestinyFile()
	assert.Equal(t, destinyFile, destiny, "error: does not is a correct directory or file")
}
