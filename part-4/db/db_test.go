package db

import (
	"testing"
)

func TestOne(t *testing.T) {
	myDataBase := new(Database)
	myDataBase.Init()

	result := myDataBase.Add("xxx", "zzz")
	if result != 0 || len(myDataBase.mapInformation) != 1 {
		t.Errorf("error add the first element")
	}
	result = myDataBase.Add("xxy", "zzz")
	if result != 0 || len(myDataBase.mapInformation) != 2 {
		t.Errorf("error add the second element")
	}
	result = myDataBase.Add("xxx", "zzz")
	if result != -1 || len(myDataBase.mapInformation) != 2 {
		t.Errorf("error add the thert element, result = %d \n and lengt = %d", result, len(myDataBase.mapInformation))
	}
}

func TestTwo(t *testing.T) {
	var result int
	var tempString string
	myDataBase := new(Database)
	myDataBase.Init()

	/*add 9 items*/
	result = myDataBase.Add("xxx", "zzz")
	if result != 0 || len(myDataBase.mapInformation) != 1 {
		t.Errorf("error trying to add an item (1)")
	}
	result = myDataBase.Add("xxy", "zzz")
	if result != 0 || len(myDataBase.mapInformation) != 2 {
		t.Errorf("error trying to add an item (2)")
	}
	result = myDataBase.Add("xxz", "zzz")
	if result != 0 || len(myDataBase.mapInformation) != 3 {
		t.Errorf("error trying to add an item (3)")
	}
	result = myDataBase.Add("xyx", "zzz")
	if result != 0 || len(myDataBase.mapInformation) != 4 {
		t.Errorf("error trying to add an item (4)")
	}
	result = myDataBase.Add("xyy", "zzz")
	if result != 0 || len(myDataBase.mapInformation) != 5 {
		t.Errorf("error trying to add an item (5)")
	}
	result = myDataBase.Add("xyz", "zz2")
	if result != 0 || len(myDataBase.mapInformation) != 6 {
		t.Errorf("error trying to add an item (6)")
	}
	result = myDataBase.Add("xzx", "zzz")
	if result != 0 || len(myDataBase.mapInformation) != 7 {
		t.Errorf("error trying to add an item (7)")
	}
	result = myDataBase.Add("xzy", "zzz")
	if result != 0 || len(myDataBase.mapInformation) != 8 {
		t.Errorf("error trying to add an item (8)")
	}
	result = myDataBase.Add("xzz", "zzz")
	if result != 0 || len(myDataBase.mapInformation) != 9 {
		t.Errorf("error trying to add an item (9)")
	}

	/*try to add item with existing key*/
	result = myDataBase.Add("xxx", "zzz")
	if result != -1 || len(myDataBase.mapInformation) != 9 {
		t.Errorf("error trying to add a same item")
	}

	/*retrive 3 items, the first and last to add, and one more in the middle*/
	result, tempString = myDataBase.Retrieve("xxx")
	if result != 0 || len(myDataBase.mapInformation) != 9 || tempString != "zzz" {
		t.Errorf("error to retrive an item, result = %d \n and lengt = %d", result, len(myDataBase.mapInformation))
	}
	result, tempString = myDataBase.Retrieve("xyz")
	if result != 0 || len(myDataBase.mapInformation) != 9 || tempString != "zz2" {
		t.Errorf("error to retrive an item, result = %d \n and lengt = %d", result, len(myDataBase.mapInformation))
	}
	result, tempString = myDataBase.Retrieve("xzz")
	if result != 0 || len(myDataBase.mapInformation) != 9 || tempString != "zzz" {
		t.Errorf("error to retrive an item, result = %d \n and lengt = %d", result, len(myDataBase.mapInformation))
	}

	/*update a item*/
	result, tempString = myDataBase.Update("xzz", "222")
	if result != 0 || len(myDataBase.mapInformation) != 9 || tempString != "222" {
		t.Errorf("error trying to update an item")
	}

	/*try to update a non-existent item*/
	result, _ = myDataBase.Update("zzz", "333")
	if result != -1 || len(myDataBase.mapInformation) != 9 {
		t.Errorf("error trying to update an non-existent item")
	}

	/*delete 3 items,  the first and last to add, and one more in the middle*/
	result = myDataBase.Delete("xxx")
	if result != 0 || len(myDataBase.mapInformation) != 8 {
		t.Errorf("error trying to remove the first item")
	}
	result = myDataBase.Delete("xyx")
	if result != 0 || len(myDataBase.mapInformation) != 7 {
		t.Errorf("error trying to remove an middle item")
	}
	result = myDataBase.Delete("xzz")
	if result != 0 || len(myDataBase.mapInformation) != 6 {
		t.Errorf("error trying to remove the last item")
	}

	/*try delete a non-existent item*/
	result = myDataBase.Delete("xzz")
	if result != -1 || len(myDataBase.mapInformation) != 6 {
		t.Errorf("error trying to remove a non-existent item")
	}

}
