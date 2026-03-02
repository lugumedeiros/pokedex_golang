package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Area struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type Location struct {
	Name string
	Url  string
}

type LocationApiStruct struct {
	Count    int
	Next     string
	Previous string
	Results  []Location
}

type api_url_config struct {
	current   string
	previous  string
	next      string
	cache     *Cache
	pokeCache *PokeCache
}

var globalConfig = api_url_config{
	current:   "https://pokeapi.co/api/v2/location-area/",
	previous:  "",
	next:      "",
	cache:     NewLocCache(time.Minute * 10),
	pokeCache: NewPokeCache(time.Minute * 10),
}

func GetLocationNext() []Location {
	globalConfig.current = globalConfig.next
	return GetLocation()
}

func GetLocationBack() []Location {
	globalConfig.current = globalConfig.previous
	return GetLocation()
}

func GetLocation() []Location {
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

func GetPokemonInArea(location string) []string {
	if pokemons, ok := globalConfig.pokeCache.Get(location); ok {
		return pokemons
	}

	var areaData Area
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v", location)
	res, err_get := http.Get(url)
	if err_get != nil {
		return nil
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	err_dec := decoder.Decode(&areaData)
	if err_dec != nil {
		return nil
	}

	var pokemons []string
	for _, enconter := range areaData.PokemonEncounters {
		pokemons = append(pokemons, enconter.Pokemon.Name)
	}
	// Store in Cache
	globalConfig.pokeCache.Add(location, pokemons)
	return pokemons
}
