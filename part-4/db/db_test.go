package db

/*
import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

const directory = "testdb/"

func TestLoadFile(t *testing.T) {
	tempDB, err := CreateNewDBInstance(directory, "testRead", false)
	assert.NoError(t, err, "error: tring create a new db instance")

	expectedValues := []struct {
		id       string
		title    string
		price    string
		quantity int
	}{
		{
			id:       "12",
			title:    "Pizza",
			price:    "13.10",
			quantity: 1,
		},
		{
			id:       "2",
			title:    "Apple",
			price:    "3.20",
			quantity: 3,
		},
	}

	err = tempDB.LoadFile()
	assert.NoError(t, err, "error: tring load file")

	for _, item := range expectedValues {
		value, isUsed := tempDB.mapInformation[item.id]
		assert.True(t, isUsed, "error: item does not exist in the map")
		assert.Equal(t, item.title, value.Title, "error: wrong item title")
		assert.Equal(t, item.price, value.Price, "error: wrong item price")
		assert.Equal(t, item.quantity, value.Quantity, "error: wrong item quantity")
	}
}

func TestAdd(t *testing.T) {
	tempDB, err := CreateNewDBInstance(directory, "testRead", false)
	assert.NoError(t, err, "error: tring create a new db instance")

	expectedValues := []struct {
		key              string
		expectedTitle    string
		expectedPrice    string
		expectedQuantity int
		length           int
		err              error
	}{
		{
			key:              "1",
			expectedTitle:    "Bannana",
			expectedPrice:    "2.50",
			expectedQuantity: 1,
			length:           1,
			err:              nil,
		},
		{
			key:              "4",
			expectedTitle:    "Noodles",
			expectedPrice:    "23.50",
			expectedQuantity: 1,
			length:           2,
			err:              nil,
		},
		{
			key:              "4",
			expectedTitle:    "Noodles",
			expectedPrice:    "23.50",
			expectedQuantity: 2,
			length:           2,
			err:              nil,
		},
		{
			key:              "12",
			expectedTitle:    "Pizza",
			expectedPrice:    "13.10",
			expectedQuantity: 1,
			length:           3,
			err:              nil,
		},
		{
			key:              "15",
			expectedTitle:    "",
			expectedPrice:    "",
			expectedQuantity: 1,
			length:           3,
			err:              errors.New("item does not exist"),
		},
	}

	for _, item := range expectedValues {
		err := tempDB.Add(item.key)
		assert.Equal(t, item.err, err, "error: item added error")
		assert.Equal(t, item.length, len(tempDB.mapInformation), "error: add a wrong quantity of items")
		if err == nil {
			assert.Equal(t, item.expectedTitle, tempDB.mapInformation[item.key].Title, "error: item was not have a correct title")
			assert.Equal(t, item.expectedPrice, tempDB.mapInformation[item.key].Price, "error: item was not have a correct price")
			assert.Equal(t, item.expectedQuantity, tempDB.mapInformation[item.key].Quantity, "error: item was not have a correct quantity")
		}
	}
}

func TestReturnMap(t *testing.T) {
	tempDB, err := CreateNewDBInstance(directory, "testRead", false)
	assert.NoError(t, err, "error: tring create a new db instance")

	expectedValues := []struct {
		id       string
		title    string
		price    string
		quantity int
	}{
		{
			id:       "12",
			title:    "Pizza",
			price:    "13.10",
			quantity: 1,
		},
		{
			id:       "2",
			title:    "Apple",
			price:    "3.20",
			quantity: 3,
		},
	}

	err = tempDB.LoadFile()
	assert.NoError(t, err, "error: tring load file")

	myMap, err := tempDB.ReturnMap()
	assert.NoError(t, err, "error: tring return map")

	for _, item := range expectedValues {
		value, isUsed := myMap[item.id]
		assert.True(t, isUsed, "error: item does not exist in the map")
		assert.Equal(t, item.title, value.Title, "error: wrong item title")
		assert.Equal(t, item.price, value.Price, "error: wrong item price")
		assert.Equal(t, item.quantity, value.Quantity, "error: wrong item quantity")
	}
}

func TestRetrieve(t *testing.T) {
	tempDB, err := CreateNewDBInstance(directory, "testRead", false)
	assert.NoError(t, err, "error: tring create a new db instance")

	expectedValues := []struct {
		id       string
		title    string
		price    string
		quantity int
	}{
		{
			id:       "12",
			title:    "Pizza",
			price:    "13.10",
			quantity: 1,
		},
		{
			id:       "2",
			title:    "Apple",
			price:    "3.20",
			quantity: 3,
		},
	}

	err = tempDB.LoadFile()
	assert.NoError(t, err, "error: tring load file")

	for _, item := range expectedValues {
		myItem, err := tempDB.Retrieve(item.id)
		assert.NoError(t, err, "error: tring retrieve a Item")
		assert.Equal(t, item.title, myItem.Title, "error: wrong item title")
		assert.Equal(t, item.price, myItem.Price, "error: wrong item price")
		assert.Equal(t, item.quantity, myItem.Quantity, "error: wrong item quantity")
	}
}

func TestUpdate(t *testing.T) {
	tempDB, err := CreateNewDBInstance(directory, "testRead", false)
	assert.NoError(t, err, "error: tring create a new db instance")

	expectedValues := []struct {
		actualID string
		newID    string
		title    string
		price    string
		quantity int
	}{
		{
			actualID: "12",
			newID:    "4",
			title:    "Noodles",
			price:    "23.50",
			quantity: 1,
		},
		{
			actualID: "2",
			newID:    "3",
			title:    "Cookies",
			price:    "10.40",
			quantity: 3,
		},
	}

	err = tempDB.LoadFile()
	assert.NoError(t, err, "error: tring load file")

	for _, item := range expectedValues {
		err := tempDB.Update(item.actualID, item.newID)
		assert.NoError(t, err, "error: tring update a Item")
		value, isUsed := tempDB.mapInformation[item.newID]
		assert.True(t, isUsed, "error: item does not exist in the map")
		assert.Equal(t, item.title, value.Title, "error: wrong item title")
		assert.Equal(t, item.price, value.Price, "error: wrong item price")
		assert.Equal(t, item.quantity, value.Quantity, "error: wrong item quantity")
		_, isUsed = tempDB.mapInformation[item.actualID]
		assert.False(t, isUsed, "error: item exist in the map")
	}
}

func TestDelete(t *testing.T) {
	tempDB, err := CreateNewDBInstance(directory, "testRead", false)
	assert.NoError(t, err, "error: tring create a new db instance")

	expectedValues := []struct {
		id     string
		length int
	}{
		{
			id:     "12",
			length: 1,
		},
		{
			id:     "2",
			length: 0,
		},
	}

	err = tempDB.LoadFile()
	assert.NoError(t, err, "error: tring load file")

	for _, item := range expectedValues {
		_, isUsed := tempDB.mapInformation[item.id]
		assert.True(t, isUsed, "error: item does not exist in the map")
		err := tempDB.Delete(item.id)
		assert.Equal(t, item.length, len(tempDB.mapInformation), "error: does not delete the items")
		assert.NoError(t, err, "error: tring delete a Item")
		_, isUsed = tempDB.mapInformation[item.id]
		assert.False(t, isUsed, "error: item exist in the map")
	}
}

func TestSaveFile(t *testing.T) {
	tempDB, err := CreateNewDBInstance(directory, "testWrite", false)
	assert.NoError(t, err, "error: tring create a new db instance")

	examplesValues := []struct {
		key string
	}{
		{
			key: "1",
		},
		{
			key: "4",
		},
		{
			key: "4",
		},
		{
			key: "12",
		},
	}

	for _, item := range examplesValues {
		err := tempDB.Add(item.key)
		assert.NoError(t, err, "error: item added error")
	}

	assert.NoError(t, tempDB.SaveFile(), "error: does not save the map")
}

func TestDeleteFile(t *testing.T) {
	_, err := CreateNewDBInstance(directory, "testDelete", true)
	assert.NoError(t, err, "error: tring create a new db file")

	tempDB, err := CreateNewDBInstance(directory, "testDelete", false)
	assert.NoError(t, err, "error: tring create a new db instance")

	assert.NoError(t, tempDB.DeleteFile(), "error: does not delete the map")
}*/
