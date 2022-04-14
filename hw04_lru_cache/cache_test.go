package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(3)

		c.Set("a", 100)
		c.Set("b", 200)
		c.Set("c", 300)
		c.Set("d", 400)

		valueA, okA := c.Get("a")
		require.False(t, okA)
		require.Nil(t, valueA)

		c.Get("b")
		c.Get("c")
		c.Set("b", 500)
		c.Set("d", 600)
		c.Get("b")
		c.Set("e", 700)

		valueB, okB := c.Get("b")
		require.True(t, okB)
		require.Equal(t, 500, valueB)

		valueC, okC := c.Get("c")
		require.False(t, okC)
		require.Nil(t, valueC)
		existsA := c.Set("a", 800)
		require.False(t, existsA)
	})

	t.Run("capacity = 0", func(t *testing.T) {
		c := NewCache(0)

		c.Set("a", 100)

		valueA, okA := c.Get("a")
		require.False(t, okA)
		require.Nil(t, valueA)
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
