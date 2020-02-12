package db

import (
	"testing"
)

func TestOne(t *testing.T) {
	myDataBase := new(Database)
	myDataBase.Init()

	err := myDataBase.Add("xxx", "zzz")
	if len(myDataBase.mapInformation) != 1 || err != nil {
		t.Errorf("error add the first element")
	}
	err = myDataBase.Add("xxy", "zzz")
	if len(myDataBase.mapInformation) != 2 || err != nil {
		t.Errorf("error add the second element")
	}
	err = myDataBase.Add("xxx", "zzz")
	if len(myDataBase.mapInformation) != 2 || err == nil {
		t.Errorf("error add the thert element, lengt = %d", len(myDataBase.mapInformation))
	}
}

func TestTwo(t *testing.T) {
	var tempString string
	myDataBase := CreateNewDBInstance()
	myDataBase.Init()

	/*add 9 items*/
	err := myDataBase.Add("xxx", "zzz")
	if len(myDataBase.mapInformation) != 1 || err != nil {
		t.Errorf("error trying to add an item (1)")
	}
	err = myDataBase.Add("xxy", "zzz")
	if len(myDataBase.mapInformation) != 2 || err != nil {
		t.Errorf("error trying to add an item (2)")
	}
	err = myDataBase.Add("xxz", "zzz")
	if len(myDataBase.mapInformation) != 3 || err != nil {
		t.Errorf("error trying to add an item (3)")
	}
	err = myDataBase.Add("xyx", "zzz")
	if len(myDataBase.mapInformation) != 4 || err != nil {
		t.Errorf("error trying to add an item (4)")
	}
	err = myDataBase.Add("xyy", "zzz")
	if len(myDataBase.mapInformation) != 5 || err != nil {
		t.Errorf("error trying to add an item (5)")
	}
	err = myDataBase.Add("xyz", "zz2")
	if len(myDataBase.mapInformation) != 6 || err != nil {
		t.Errorf("error trying to add an item (6)")
	}
	err = myDataBase.Add("xzx", "zzz")
	if len(myDataBase.mapInformation) != 7 || err != nil {
		t.Errorf("error trying to add an item (7)")
	}
	err = myDataBase.Add("xzy", "zzz")
	if len(myDataBase.mapInformation) != 8 || err != nil {
		t.Errorf("error trying to add an item (8)")
	}
	err = myDataBase.Add("xzz", "zzz")
	if len(myDataBase.mapInformation) != 9 || err != nil {
		t.Errorf("error trying to add an item (9)")
	}

	/*try to add item with existing key*/
	err = myDataBase.Add("xxx", "zzz")
	if len(myDataBase.mapInformation) != 9 || err == nil {
		t.Errorf("error trying to add a same item")
	}

	/*retrive 3 items, the first and last to add, and one more in the middle*/
	tempString, err = myDataBase.Retrieve("xxx")
	if len(myDataBase.mapInformation) != 9 || tempString != "zzz" || err != nil {
		t.Errorf("error to retrive an item, lengt = %d", len(myDataBase.mapInformation))
	}
	tempString, err = myDataBase.Retrieve("xyz")
	if len(myDataBase.mapInformation) != 9 || tempString != "zz2" || err != nil {
		t.Errorf("error to retrive an item, lengt = %d", len(myDataBase.mapInformation))
	}
	tempString, err = myDataBase.Retrieve("xzz")
	if len(myDataBase.mapInformation) != 9 || tempString != "zzz" || err != nil {
		t.Errorf("error to retrive an item, lengt = %d", len(myDataBase.mapInformation))
	}

	/*update a item*/
	tempString, err = myDataBase.Update("xzz", "222")
	if len(myDataBase.mapInformation) != 9 || tempString != "222" || err != nil {
		t.Errorf("error trying to update an item")
	}

	/*try to update a non-existent item*/
	_, err = myDataBase.Update("zzz", "333")
	if len(myDataBase.mapInformation) != 9 || err == nil {
		t.Errorf("error trying to update an non-existent item")
	}

	/*delete 3 items,  the first and last to add, and one more in the middle*/
	err = myDataBase.Delete("xxx")
	if len(myDataBase.mapInformation) != 8 || err != nil {
		t.Errorf("error trying to remove the first item")
	}
	err = myDataBase.Delete("xyx")
	if len(myDataBase.mapInformation) != 7 || err != nil {
		t.Errorf("error trying to remove an middle item")
	}
	err = myDataBase.Delete("xzz")
	if len(myDataBase.mapInformation) != 6 || err != nil {
		t.Errorf("error trying to remove the last item")
	}

	/*try delete a non-existent item*/
	err = myDataBase.Delete("xzz")
	if len(myDataBase.mapInformation) != 6 || err == nil {
		t.Errorf("error trying to remove a non-existent item")
	}

}
