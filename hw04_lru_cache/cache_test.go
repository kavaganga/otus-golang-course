package hw04lrucache

import (
	"crypto/rand"
	"math/big"
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

		r := c.Set("1", 100)
		require.False(t, r)

		r = c.Set("2", 200)
		require.False(t, r)

		r = c.Set("3", 300)
		require.False(t, r)

		v, r := c.Get("1")
		require.True(t, r)
		require.Equal(t, 100, v)

		v, r = c.Get("2")
		require.True(t, r)
		require.Equal(t, 200, v)

		v, r = c.Get("3")
		require.True(t, r)
		require.Equal(t, 300, v)

		c.Clear()
		_, r = c.Get("1")
		require.False(t, r)
		_, r = c.Get("2")
		require.False(t, r)
		_, r = c.Get("3")
		require.False(t, r)
	})
}

func TestCacheMultithreading(t *testing.T) {
	rnd, err := rand.Int(rand.Reader, big.NewInt(1_000_000))
	if err != nil {
		t.Fatal(err)
	}
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
			c.Get(Key(strconv.Itoa(int(rnd.Int64()))))
		}
	}()

	wg.Wait()
}
