package cache

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

var cl = NewStore(WithLength(100))

func TestStore_Clear(t *testing.T) {
	for i := 0; i < 10; i++ {
		cl.SetCache(strconv.Itoa(i), strconv.Itoa(i))
	}
	cl.Clear()
	assert.Equal(t, 0, cl.l.Len())
	mLength := 0
	cl.m.Range(func(k, v interface{}) bool {
		mLength++
		return true
	})
	assert.Equal(t, 0, mLength)
}

func TestStore_AddCache(t *testing.T) {
	cl.Clear()
	for i := 0; i < 100; i++ {
		cl.SetCache(strconv.Itoa(i), strconv.Itoa(i))
		length := 0
		cl.m.Range(func(key, value interface{}) bool {
			length++
			return true
		})
		assert.Equal(t, length, cl.l.Len(), "err")
		assert.LessOrEqual(t, cl.l.Len(), cl.maxLength, "err")
		assert.LessOrEqual(t, cl.l.Len(), cl.maxLength, "err")
	}
}

func TestStore_RemoveCache(t *testing.T) {
	cl.Clear()
	for i := 0; i < 10; i++ {
		cl.SetCache(strconv.Itoa(i), strconv.Itoa(i))
	}
	cl.RemoveCache("1")
	a, ok := cl.GetCache("1")
	assert.Equal(t, nil, a, "err")
	assert.Equal(t, false, ok, "err")
	assert.Equal(t, 9, cl.l.Len())
	mLength := 0
	cl.m.Range(func(k, v interface{}) bool {
		mLength++
		return true
	})
	assert.Equal(t, 9, mLength)
}

func TestStore_GetCache(t *testing.T) {
	cl.Clear()
	for i := 0; i < 100; i++ {
		cl.SetCache(strconv.Itoa(i), strconv.Itoa(i))
	}
	a1, ok1 := cl.GetCache("95")
	a2, ok2 := cl.GetCache("5")
	assert.Equal(t, strconv.Itoa(95), a1, "err")
	v, _ := cl.m.Load(cl.l.Front().Value.(string))
	assert.Equal(t, v.(Cache).value, a2, "err")
	assert.True(t, ok1, "err")
	assert.True(t, ok2, "err")
}

func TestStore_SetExpired(t *testing.T) {
	cl.Clear()
	cl.SetCache("test", "test")
	cl.SetExpired("test", 1*time.Second)
	a1, ok1 := cl.GetCache("test")
	assert.Equal(t, "test", a1)
	assert.True(t, ok1)
	time.Sleep(1 * time.Second + time.Millisecond)
	_, ok2 := cl.GetCache("test")
	assert.False(t, ok2)
}

func TestStore_SetExpiredCache(t *testing.T) {
	cl.Clear()

	cl.SetExpiredCache("test", "test", 1*time.Second)
	a1, ok1 := cl.GetCache("test")
	assert.Equal(t, "test", a1)
	assert.True(t, ok1)
	time.Sleep(1 * time.Second + time.Millisecond)
	_, ok2 := cl.GetCache("test")
	assert.False(t, ok2)
}
