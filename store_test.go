package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewList(t *testing.T) {
	c := NewStore()
	assert.Equal(t, 10000, c.maxLength)
	assert.Equal(t, 100, cl.maxLength)
}
