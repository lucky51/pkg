package test

import (
	"testing"

	"github.com/lucky51/pkg/ext"
	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {

	assert.Equal(t, ext.Min(3, 4, 5, 10, 9), 3)
	assert.Equal(t, ext.Min(3, 4, 5, 10, 9, 1), 1)
	assert.Equal(t, ext.Min(3, 4, 5, 10, 9, 1, 10), 1)
	assert.Equal(t, ext.Min(10, 9, 6, 10, 20), 6)
	assert.NotEqual(t, ext.Min(10, 9, 6, 10, 20), 10)
}
