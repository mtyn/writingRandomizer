package main

import _ "encoding/json"

type Thing struct {
	Title string `json:"title"`
	Description string `json:"description"`
}

type Item struct {
	Thing
}

type Character struct {
	Thing
	PossibleItems []string 		`json:"possibleItems"`
	RelatedCharacters []string 	`json:"relatedCharacters"`
}

type Place struct {
	Thing
	PossiblePeople []string `json:"possiblePeople"`
	PossibleItems []string	`json:"possibleItems"`
	RelatedPlaces []string  `json:"relatedPlaces"`
}