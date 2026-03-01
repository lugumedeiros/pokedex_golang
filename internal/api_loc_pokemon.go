package internal

import (
	"encoding/json"
	"net/http"
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
}

var globalConfig = api_url_config {
	current: "https://pokeapi.co/api/v2/location-area/",
	previous: "",
	next: "",
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
	return LocPokeData.Results
}