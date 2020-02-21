package router

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func generateRoutesTest(t *testing.T) {
	newRouter := NewRoute()
	assert.NotNil(t, newRouter, "error with create route (length)")
}
