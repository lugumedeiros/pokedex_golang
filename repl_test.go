package main

import "testing"

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
