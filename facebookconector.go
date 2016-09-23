package main

import (
	"strings"

	fb "github.com/huandu/facebook"
)

// Make a connection with facebook API
type FacebookConector struct {
	Token string
}

func (facebook *FacebookConector) SetDefaults() {
	facebook.Token = "5d3e0d2d2ec44ab0bd0f99665847aabc"
}

// Get the required fields
func (facebook *FacebookConector) Get(fields []string, profileID string) string {
	url := "/" + profileID
	res, _ := fb.Get(url, fb.Params{
		"fields":       strings.Join(fields, ", "),
		"access_token": facebook.Token,
	})
	result := ""

	for _, key := range fields {
		result = strings.Join(res[key], ", ")
	}

	return result
}
