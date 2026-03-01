package internal

import (
	"testing"
	"time"
)

// import "fmt"

func TestCache(t *testing.T) {
	t.Logf("Testing PokeCache:\n")

	key := "Key"
	val := []Location{{Name: "Name", Url: "url.com",}}

	cache := NewCache(time.Second * 5)
	cache.Add(key, val, "", "")
	entry, ok := cache.Get(key)
	if !ok {
		t.Errorf("Test Failed, no data found from cache")
	}
	if entry.val[0].Name != val[0].Name {
		t.Errorf("Test Failed, wrong data found from cache - Expected: '%v' - Actual: '%v'", val[0].Name, entry.val[0].Name)
	}

	time.Sleep(time.Second * 6)
	entry, ok = cache.Get(key)
	if !ok {
		t.Logf("Test Pass\n")
	} else {
		t.Errorf("Test Failed, data not cleared")
	}
}
