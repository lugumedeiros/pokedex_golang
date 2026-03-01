package internal

import (
	"encoding/json"
	"net/http"
	"time"
)


type Location struct {
	Name string
	Url string
}

type LocationApiStruct struct {
	Count int
	Next string
	Previous string
	Results []Location
}

type api_url_config struct {
	current string
	previous string
	next string
	cache *Cache
}

var globalConfig = api_url_config {
	current: "https://pokeapi.co/api/v2/location-area/",
	previous: "",
	next: "",
	cache: NewCache(time.Minute * 10),
}

func GetLocationNext() []Location{
	globalConfig.current = globalConfig.next
	return GetLocation()
}

func GetLocationBack() []Location{
	globalConfig.current = globalConfig.previous
	return GetLocation()
}

func GetLocation() []Location{
	var LocPokeData LocationApiStruct
	// Get from cache
	cache_entry, ok := globalConfig.cache.Get(globalConfig.current)
	if ok {
		globalConfig.next = cache_entry.next
		globalConfig.previous = cache_entry.previous
		return cache_entry.val
	}

	res, err_get := http.Get(globalConfig.current)
	if err_get != nil {
		return nil
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	err_dec := decoder.Decode(&LocPokeData)
	if err_dec != nil {
		return nil
	}
	globalConfig.next = LocPokeData.Next
	globalConfig.previous = LocPokeData.Previous

	// Store in Cache
	globalConfig.cache.Add("current", LocPokeData.Results, LocPokeData.Next, LocPokeData.Previous)
	return LocPokeData.Results
}