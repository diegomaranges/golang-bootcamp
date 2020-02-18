package car

import (
	"testing"
)

func testOne(t *testing.T) {
	carv := CreateNewCarInstance()
	if carv == nil {
		t.Errorf("error creating a new car")
	}
	err := carv.AddItem("22")
	if err != nil {
		t.Errorf("error creating a new car")
	}
}
