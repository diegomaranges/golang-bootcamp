package db

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

const destinyFile string = "info.txt"

func TestAdd(t *testing.T) {
	myDataBase := CreateNewDBInstance(destinyFile)

	tables := []struct {
		key           string
		value         string
		expectedValue string
		length        int
		err           error
		element       string
	}{
		{"1", "1", "1", 1, nil, "first"},
		{"2", "2", "2", 2, nil, "second"},
		{"2", "5", "2", 2, errors.New("Key is already exist"), "second"},
		{"3", "3", "3", 3, nil, "third"},
		{"4", "4", "4", 4, nil, "fourth"},
	}

	for _, table := range tables {
		err := myDataBase.Add(table.key, table.value)
		assert.Equal(t, table.length, len(myDataBase.mapInformation), "error: element was not added", table.element)
		assert.Equal(t, table.err, err, "error: element added error", table.element)
		assert.Equal(t, table.expectedValue, myDataBase.mapInformation[table.key], "error: element add a wrong value")
	}
}

func TestRetrieve(t *testing.T) {
	myDataBase := CreateNewDBInstance(destinyFile)
	myDataBase.mapInformation["1"] = "1"
	myDataBase.mapInformation["2"] = "2"
	myDataBase.mapInformation["3"] = "3"
	myDataBase.mapInformation["4"] = "4"
	myDataBase.mapInformation["124"] = "34"

	tables := []struct {
		key     string
		value   string
		length  int
		err     error
		element string
	}{
		{"1", "1", 5, nil, "first"},
		{"2", "2", 5, nil, "second"},
		{"3", "3", 5, nil, "third"},
		{"4", "4", 5, nil, "fourth"},
		{"124", "34", 5, nil, "fifth"},
		{"125", "", 5, errors.New("Key does not exist"), "fifth"},
	}

	for _, table := range tables {
		value, err := myDataBase.Retrieve(table.key)
		assert.Equal(t, table.value, value, "error: do not retrieve a correct value", table.element)
		assert.Equal(t, table.length, len(myDataBase.mapInformation), "error: retrieve edit the map", table.element)
		assert.Equal(t, table.err, err, "error: do not retrieved", table.element)
	}
}

func TestUpdate(t *testing.T) {
	myDataBase := CreateNewDBInstance(destinyFile)
	myDataBase.mapInformation["1"] = "1"
	myDataBase.mapInformation["2"] = "2"
	myDataBase.mapInformation["3"] = "3"
	myDataBase.mapInformation["4"] = "4"
	myDataBase.mapInformation["124"] = "34"

	tables := []struct {
		key      string
		newValue string
		isUsed   bool
		length   int
		err      error
		element  string
	}{
		{"1", "5", true, 5, nil, "1"},
		{"2", "6", true, 5, nil, "2"},
		{"3", "7", true, 5, nil, "3"},
		{"4", "8", true, 5, nil, "4"},
		{"124", "9", true, 5, nil, "5"},
		{"125", "", false, 5, errors.New("Key does not exist"), "6"},
	}

	for _, table := range tables {
		err := myDataBase.Update(table.key, table.newValue)
		assert.Equal(t, table.length, len(myDataBase.mapInformation), "error: update element edit size of map", table.element)
		assert.Equal(t, table.err, err, "error: element was not updated", table.element)
		value, isUsed := myDataBase.mapInformation[table.key]
		assert.Equal(t, table.isUsed, isUsed, "error: element was not updated correctly", table.element)
		if table.isUsed {
			assert.Equal(t, table.newValue, value, "error: value changed when update the element", table.element)
		}
	}
}

func TestDelete(t *testing.T) {
	myDataBase := CreateNewDBInstance(destinyFile)
	myDataBase.mapInformation["1"] = "1"
	myDataBase.mapInformation["2"] = "2"
	myDataBase.mapInformation["3"] = "3"
	myDataBase.mapInformation["4"] = "4"
	myDataBase.mapInformation["124"] = "34"

	tables := []struct {
		key     string
		length  int
		err     error
		element string
	}{
		{"1", 4, nil, "first"},
		{"2", 3, nil, "second"},
		{"3", 2, nil, "third"},
		{"4", 1, nil, "fourth"},
		{"125", 1, errors.New("Key does not exist"), "fifth"},
	}

	for _, table := range tables {
		err := myDataBase.Delete(table.key)
		assert.Equal(t, table.length, len(myDataBase.mapInformation), "error: element was not deleted", table.element)
		assert.Equal(t, table.err, err, "error: element deleted error", table.element)
	}
}
