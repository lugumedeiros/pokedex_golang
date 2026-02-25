package main

import (
	"strings"
)

// import "regexp"

func cleanInput(text string) []string {
	str_slice := strings.Fields(text)
	var new_slice []string

	for i := range str_slice {
		lower := strings.ToLower(str_slice[i])
		new_slice = append(new_slice, lower)
	}
	return new_slice
}
