package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://pokeapi.co/api/v2/location-area",
			val: []byte("testdata"),
		},
		{
			key: "https://pokeapi.co/api/v2/location-area/1",
			val: []byte("testdata2"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected %s to be in the cache", c.key)
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected %s, got %s", c.val, val)
				return
			}
		})
	}
}

func TestClean(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://pokeapi.co/api/v2/location-area", []byte("testdata"))

	_, ok := cache.Get("https://pokeapi.co/api/v2/location-area")
	if !ok {
		t.Errorf("expected https://pokeapi.co/api/v2/location-area to be in the cache")
		return
	}
	
	time.Sleep(waitTime)

	_, ok = cache.Get("https://pokeapi.co/api/v2/location-area")
	if ok {
		t.Errorf("expected https://pokeapi.co/api/v2/location-area to be removed from the cache")
		return
	}
}
