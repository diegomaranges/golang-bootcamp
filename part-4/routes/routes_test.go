package routes

import (
	"testing"
)

func testOne(t *testing.T) {
	var router *Router
	router.NewRouter()
	if router == nil {
		t.Errorf("error creating a new route")
	}
}
