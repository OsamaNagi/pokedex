package internals

import (
	"fmt"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://pokeapi.co/api/v2/location-area",
			val: []byte("testdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestGet(t *testing.T) {
	const interval = 5 * time.Second
	cache := NewCache(interval)
	cache.Add("https://pokeapi.co/api/v2/location-area", []byte("testdata"))

	val, ok := cache.Get("https://pokeapi.co/api/v2/location-area")
	if !ok {
		t.Errorf("expected to find key")
		return
	}
	if string(val) != "testdata" {
		t.Errorf("expected to find value")
		return
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Second
	const waitTime = baseTime + 5*time.Second
	cache := NewCache(baseTime)
	cache.Add("https://pokeapi.co/api/v2/location-area", []byte("testdata"))

	_, ok := cache.Get("https://pokeapi.co/api/v2/location-area")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://pokeapi.co/api/v2/location-area")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
