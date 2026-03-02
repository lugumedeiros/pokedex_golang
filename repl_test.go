package main

import "testing"
import "pokedexcligo/internal"

// import "fmt"

type CaseStruct struct {
	input    string
	expected []string
}

func TestCleanInput(t *testing.T) {
	case_1 := CaseStruct{input: " Hello world! ", expected: []string{"hello", "world!"}}
	case_2 := CaseStruct{input: "   ", expected: []string{}}
	case_3 := CaseStruct{input: "test a b", expected: []string{"test", "a", "b"}}
	all_cases := []CaseStruct{case_1, case_2, case_3}
	t.Logf("Testing Input CLI:\n")
	for i, c := range all_cases {
		t.Logf("Testing case: %v\n", i+1)
		actual := cleanInput(c.input)

		if len(c.expected) != len(actual) {
			t.Errorf("Test Failed (LEN) - Expected: %v - Actual: %v\n", c.expected, actual)
			continue
		}

		for j := range len(c.expected) {
			if actual[j] != c.expected[j] {
				t.Errorf("Test Failed (STR) - Expected: '%v' - Actual: '%v'\n", c.expected[j], actual[j])
			}
		}
		t.Logf("Test Pass\n")
	}
}

func TestPokeAPILocationLogic(t *testing.T) {
	locations := internal.GetLocation()
	has_passed := true
	t.Logf("Testing Location API:\n")
	if locations == nil {
		t.Errorf("Test Failed - No location retrieved")
		has_passed = false
	} else {
		if locations[0].Name != "canalave-city-area"{
			t.Errorf("Test Failed - Expected: '%v' - Actual: '%v'\n", "canalave-city-area", locations[0].Name)
			has_passed = false
		}
	}

	new_locations := internal.GetLocationNext()
	if new_locations == nil {
		t.Errorf("Test Failed - No location retrieved")
		has_passed = false
	} else {
		if new_locations[0].Name != "mt-coronet-1f-route-216"{
			t.Errorf("Test Failed - Expected: '%v' - Actual: '%v'\n", "mt-coronet-1f-route-216", new_locations[0].Name)
			has_passed = false
		}
	}
	if has_passed {
		t.Logf("Test Pass\n")
	}
}

func TestPokeAPIAreaLogic(t *testing.T){
	t.Logf("Testing Area API:\n")
	pokemons := internal.GetPokemonInArea("1")
	if len(pokemons) == 0 {
		t.Errorf("Test Failed - Failed to find any pokemon in area 1")
	}
	if pokemons[0] != "tentacool" {
		t.Errorf("Test Failed - Expected: '%v' - Actual: '%v'\n", "tentacool", pokemons[0])
	}
	t.Logf("Test Pass\n")
}

func TestPokeCatchLogic(t *testing.T){
	t.Logf("Testing Catch API:\n")
	pokemon_name := "pikachu"
	success, err := internal.GetPokemon(pokemon_name)
	if err {
		t.Errorf("Error happened in the API")
	}
	if success {
		t.Errorf("Expected to fail, but pikachu was caught")
	}
	for i:=0; i < 20; i++ {
		success, err = internal.GetPokemon(pokemon_name)
		
		if err {
			t.Errorf("Error happened in the API")
		}
		if success {
			t.Logf("Test Pass\n")
			return
		}
	}
	t.Errorf("Expected to succes, but pikachu was not caught")
}
