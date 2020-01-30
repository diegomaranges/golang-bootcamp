package db

import (
	"bufio"
	"os"
	"testing"

	"github.corp.globant.com/diego-maranges/golang-bootcamp/parte2-4/db"
)

func TestOne(t *testing.T) {
	myDataBase := new(db.Database)
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
	var myDataBase database
	myDataBase.mapInformation = make(map[string]string)
	myDataBase.inputType = bufio.NewScanner(os.Stdin)

	result = myDataBase.add("xxx", "zzz")
	if result != 0 || len(myDataBase.mapInformation) != 1 {
		t.Errorf("error add the first element")
	}
	result = myDataBase.add("xxy", "zzz")
	if result != 0 || len(myDataBase.mapInformation) != 2 {
		t.Errorf("error add the second element")
	}
	result, tempString = myDataBase.retrieve("xxx")
	if result != 0 || len(myDataBase.mapInformation) != 2 || tempString != "zzz" {
		t.Errorf("error add the thert element, result = %d \n and lengt = %d", result, len(myDataBase.mapInformation))
	}
	result, tempString = myDataBase.update("xxx", "222")
	if result != 0 || len(myDataBase.mapInformation) != 2 || tempString != "222" {
		t.Errorf("error add the ford element")
	}
	result = myDataBase.delete("xxy")
	if result != 0 || len(myDataBase.mapInformation) != 1 {
		t.Errorf("error add the last element")
	}

}
