package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOne(t *testing.T) {
	myDataBase := new(Database)
	myDataBase.Init()

	err := myDataBase.Add("xxx", "zzz")
	assert.Equal(t, len(myDataBase.mapInformation), 1, "error add the first element")
	assert.NoError(t, err, "error add the first element")

	err = myDataBase.Add("xxy", "zzz")
	assert.Equal(t, len(myDataBase.mapInformation), 2, "error add the second element")
	assert.NoError(t, err, "error add the second element")

	err = myDataBase.Add("xxx", "zzz")
	assert.Equal(t, len(myDataBase.mapInformation), 2, "error add the third element")
	assert.Error(t, err, "error add the third element")
}

func TestTwo(t *testing.T) {
	myDataBase := CreateNewDBInstance()
	myDataBase.Init()

	/*add 9 items*/
	err := myDataBase.Add("xxx", "zzz")
	assert.Equal(t, len(myDataBase.mapInformation), 1, "error trying to add an item (1)")
	assert.NoError(t, err, "error trying to add an item (1)")

	err = myDataBase.Add("xxy", "zzz")
	assert.Equal(t, len(myDataBase.mapInformation), 2, "error trying to add an item (2)")
	assert.NoError(t, err, "error trying to add an item (2)")

	err = myDataBase.Add("xxz", "zzz")
	assert.Equal(t, len(myDataBase.mapInformation), 3, "error trying to add an item (3)")
	assert.NoError(t, err, "error trying to add an item (3)")

	err = myDataBase.Add("xyx", "zzz")
	assert.Equal(t, len(myDataBase.mapInformation), 4, "error trying to add an item (4)")
	assert.NoError(t, err, "error trying to add an item (4)")

	err = myDataBase.Add("xyy", "zzz")
	assert.Equal(t, len(myDataBase.mapInformation), 5, "error trying to add an item (5)")
	assert.NoError(t, err, "error trying to add an item (5)")

	err = myDataBase.Add("xyz", "zz2")
	assert.Equal(t, len(myDataBase.mapInformation), 6, "error trying to add an item (6)")
	assert.NoError(t, err, "error trying to add an item (6)")

	err = myDataBase.Add("xzx", "zzz")
	assert.Equal(t, len(myDataBase.mapInformation), 7, "error trying to add an item (7)")
	assert.NoError(t, err, "error trying to add an item (7)")

	err = myDataBase.Add("xzy", "zzz")
	assert.Equal(t, len(myDataBase.mapInformation), 8, "error trying to add an item (8)")
	assert.NoError(t, err, "error trying to add an item (8)")

	err = myDataBase.Add("xzz", "zz3")
	assert.Equal(t, len(myDataBase.mapInformation), 9, "error trying to add an item (9)")
	assert.NoError(t, err, "error trying to add an item (9)")

	/*try to add item with existing key*/
	err = myDataBase.Add("xxx", "zzz")
	assert.Equal(t, len(myDataBase.mapInformation), 9, "error trying to add a same item")
	assert.Error(t, err, "error trying to add a same item")

	/*retrive 3 items, the first and last to add, and one more in the middle*/
	tempString, err := myDataBase.Retrieve("xxx")
	assert.Equalf(t, len(myDataBase.mapInformation), 9, "error to retrive an item, lengt = %d", len(myDataBase.mapInformation))
	assert.Equal(t, tempString, "zzz", "error to retrive an item")
	assert.NoError(t, err, "error to retrive an item")

	tempString, err = myDataBase.Retrieve("xyz")
	assert.Equalf(t, len(myDataBase.mapInformation), 9, "error to retrive an item, lengt = %d", len(myDataBase.mapInformation))
	assert.Equal(t, tempString, "zz2", "error to retrive an item")
	assert.NoError(t, err, "error to retrive an item")

	tempString, err = myDataBase.Retrieve("xzz")
	assert.Equalf(t, len(myDataBase.mapInformation), 9, "error to retrive an item, lengt = %d", len(myDataBase.mapInformation))
	assert.Equal(t, tempString, "zz3", "error to retrive an item")
	assert.NoError(t, err, "error to retrive an item")

	/*update a item*/
	err = myDataBase.Update("xzz", "222")
	assert.Equalf(t, len(myDataBase.mapInformation), 9, "error trying to update an existent item")
	assert.NoError(t, err, "error trying to update an existent item")

	/*try to update a non-existent item*/
	err = myDataBase.Update("zzz", "333")
	assert.Equalf(t, len(myDataBase.mapInformation), 9, "error trying to update an non-existent item")
	assert.Error(t, err, "error trying to update an non-existent item")

	/*delete 3 items,  the first and last to add, and one more in the middle*/
	err = myDataBase.Delete("xxx")
	assert.Equalf(t, len(myDataBase.mapInformation), 8, "error trying to remove the first item")
	assert.NoError(t, err, "error trying to remove the first item")

	err = myDataBase.Delete("xyx")
	assert.Equalf(t, len(myDataBase.mapInformation), 7, "error trying to remove an middle item")
	assert.NoError(t, err, "error trying to remove an middle item")

	err = myDataBase.Delete("xzz")
	assert.Equalf(t, len(myDataBase.mapInformation), 6, "error trying to remove the last item")
	assert.NoError(t, err, "error trying to remove the last item")

	/*try delete a non-existent item*/
	err = myDataBase.Delete("xzz")
	assert.Equalf(t, len(myDataBase.mapInformation), 6, "error trying to remove a non-existent item")
	assert.Error(t, err, "error trying to remove a non-existent item")

}
