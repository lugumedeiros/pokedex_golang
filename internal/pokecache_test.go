package internal

import (
	"testing"
	"time"
)

// import "fmt"

func TestCache(t *testing.T) {
	t.Logf("Testing PokeCache:\n")

	key := "Key"
	val := "Value"

	cache := NewCache(time.Second * 5)
	cache.Add(key, []byte(val))
	bytes_from_cache, ok := cache.Get(key)
	if !ok {
		t.Errorf("Test Failed, no data found from cache")
	}
	if string(bytes_from_cache) != val {
		t.Errorf("Test Failed, wrong data found from cache - Expected: '%v' - Actual: '%v'", val, string(bytes_from_cache))
	}

	time.Sleep(time.Second * 6)
	bytes_from_cache, ok = cache.Get(key)
	if !ok {
		t.Logf("Test Pass\n")
	} else {
		t.Errorf("Test Failed, data not cleared")
	}
}
