package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func GetItems() []Item {
	jsonFile, err := os.Open("data/items.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var items []Item
	_ = json.Unmarshal(byteValue, &items)

	return items
}

func GetPlaces() []Place {
	jsonFile, err := os.Open("data/places.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var places []Place
	_ = json.Unmarshal(byteValue, &places)

	return places
}

func GetCharacters() []Character {
	jsonFile, err := os.Open("data/characters.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var characters []Character
	_ = json.Unmarshal(byteValue, &characters)
	return characters
}


func FindPlacesWithTitles(places []Place, title []string) []Place {
	// This is a bad implementation - can we do it any better?
	commonItems := []Place{}
	for i := 0; i < len(places); i++ {
		for j := 0; j < len(title); j++ {
			if places[i].Title == title[j] {
				commonItems = append(commonItems, places[i])
			}
		}
	}
	return commonItems
}

func FindCharactersWithTitles(characters []Character, title []string) []Character {
	commonItems := []Character{}
	for i := 0; i < len(characters); i++ {
		for j := 0; j < len(title); j++ {
			if characters[i].Title == title[j] {
				commonItems = append(commonItems, characters[i])
			}
		}
	}
	return commonItems
}

func FindItemsWithTitles(items []Item, title []string) []Item {
	commonItems := []Item{}
	for i := 0; i < len(items); i++ {
		for j := 0; j < len(title); j++ {
			if items[i].Title == title[j] {
				commonItems = append(commonItems, items[i])
			}
		}
	}
	return commonItems
}