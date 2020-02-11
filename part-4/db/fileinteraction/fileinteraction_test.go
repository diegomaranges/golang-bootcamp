package fileinteraction

import (
	"testing"
)

/*Run this test without file existent*/
func TestOne(t *testing.T) {
	file := new(DestinyFile)
	var mapExample map[string]string
	mapExample = make(map[string]string)

	/*Set destiny file*/
	file.SetFile("text.txt")
	resultDestiny := file.ReturnDestinyFile()
	if resultDestiny != "text.txt" {
		t.Errorf("error to set a destiny file")
	}

	/*Read and load empty file*/
	result := file.ReadFile(mapExample)
	if result != 0 || len(mapExample) != 0 {
		t.Errorf("Error to read and load a empty file")
	}

	mapExample["firstElement"] = "zzzzzz"
	mapExample["secondElement"] = "000000"

	/*write 2 elements in the file*/
	result = file.WriteFile(mapExample)
	if result != 0 {
		t.Errorf("error to write the map into the file")
	}

	mapExample["false element"] = "error"

	/*Read and load non-empty file*/
	result = file.ReadFile(mapExample)
	if result != 0 || len(mapExample) != 2 {
		t.Errorf("Error to read and load a non-empty file")
	}

	mapExample["truth element"] = "its fine"

	result = file.WriteFile(mapExample)
	if result != 0 {
		t.Errorf("error with the final write")
	}
}
