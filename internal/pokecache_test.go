package internal

import (
	"testing"
	"time"
)

// import "fmt"

func TestLocationCache(t *testing.T) {
	t.Logf("Testing Location Cache:\n")

	key := "Key"
	val := []Location{{Name: "Name", Url: "url.com"}}

	cache := NewLocCache(time.Second * 1)
	cache.Add(key, val, "", "")
	entry, ok := cache.Get(key)
	if !ok {
		t.Errorf("Test Failed, no data found from cache")
	}
	if entry.val[0].Name != val[0].Name {
		t.Errorf("Test Failed, wrong data found from cache - Expected: '%v' - Actual: '%v'", val[0].Name, entry.val[0].Name)
	}

	time.Sleep(time.Second * 2)
	entry, ok = cache.Get(key)
	if !ok {
		t.Logf("Test Pass\n")
	} else {
		t.Errorf("Test Failed, data not cleared")
	}
}

func TestAreaCache(t *testing.T) {
	t.Logf("Testing AreaC ache:\n")
	cache := NewPokeCache(time.Second * 1)

	key := "test"
	val := []string{"a", "b"}

	cache.Add(key, val)
	val_form_cache, ok := cache.Get(key)
	if !ok {
		t.Errorf("Test Failed, no data found from cache")
	}
	if val_form_cache[0] != val[0] {
		t.Errorf("Test Failed, wrong data found from cache - Expected: '%v' - Actual: '%v'", val_form_cache[0], val[0])
	}

	time.Sleep(time.Second * 2)
	val_form_cache, ok = cache.Get(key)
	if !ok {
		t.Logf("Test Pass\n")
	} else {
		t.Errorf("Test Failed, data not cleared")
	}

}
