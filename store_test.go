package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewList(t *testing.T) {
	c := NewStore()
	assert.Greater(t, c.maxLength, 0)
}
