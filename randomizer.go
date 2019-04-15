package main

import (
	"flag"
	"github.com/fatih/color"
	"math/rand"
)

func getARandomPlace(places []Place, lastVisited *Place, totallyRandom bool) Place {
	var filteredPlaces []Place
	if lastVisited != nil && !totallyRandom {
		filteredPlaces = FindPlacesWithTitles(places, lastVisited.RelatedPlaces)
	} else {
		filteredPlaces = places
	}
	if len(filteredPlaces) == 0 {
		color.Red("No linked places available, choosing a random one")
		filteredPlaces = places
	}
	randomIndex := rand.Intn(len(filteredPlaces))
	return filteredPlaces[randomIndex]
}

func getARandomCharacter(characters []Character, relatedTo *Character, relatedPlace *Place, totallyRandom bool) Character {
	var filteredCharacters []Character
	if relatedPlace != nil && !totallyRandom {
		filteredCharacters = FindCharactersWithTitles(characters, relatedPlace.PossiblePeople)

	} else {
		filteredCharacters = characters
	}
	if relatedTo != nil && !totallyRandom {
		filteredCharacters = FindCharactersWithTitles(filteredCharacters, relatedTo.RelatedCharacters)
	}
	if len(filteredCharacters) == 0 {
		color.Red("No linked characters available, choosing a random one.")
		filteredCharacters = characters
	}
	randomIndex := rand.Intn(len(filteredCharacters))
	return filteredCharacters[randomIndex]
}

func getARandomItem(items []Item, character *Character, place *Place, totallyRandom bool) Item {
	var filteredItems []Item
	if place != nil && !totallyRandom {
		filteredItems = FindItemsWithTitles(items, place.PossibleItems)
	} else {
		filteredItems = items
	}
	if character != nil && !totallyRandom {
		filteredItems = FindItemsWithTitles(filteredItems, character.RelatedCharacters)
	}
	if len(filteredItems) == 0 {
		color.Red("No linked items available, choosing a random one.")
		filteredItems = items
	}
	randomIndex := rand.Intn(len(filteredItems))
	return filteredItems[randomIndex]
}

func printPlace(place Place) {
	color.Green("PLACE:\n%s\n%s\n", place.Title, place.Description)
}

func printCharacter(character Character) {
	color.Magenta("CHARACTER:\n%s\n%s\n", character.Title, character.Description)
}

func printItem(item Item) {
	color.Cyan("ITEM:\n%s\n%s\n", item.Title, item.Description)
}

func main() {
	// Set up the flags
	totallyRandom := flag.Bool("fullRandom", false, "Ignores relations between " +
		"items (just throw out random prompts," +
		"which may have no relation)")
	randomBool := *totallyRandom
	places := GetPlaces()
	items := GetItems()
	characters := GetCharacters()

	initialPlace := getARandomPlace(places, nil, randomBool)
	initialCharacter := getARandomCharacter(characters, nil, &initialPlace, randomBool)
	initialItem := getARandomItem(items, &initialCharacter, &initialPlace, randomBool)

	color.Blue("Welcome to the writing prompt randomizer!\n")
	color.Black("Your initial prompt is:\n")
	printPlace(initialPlace)
	println("")
	printCharacter(initialCharacter)
	println("")
	printItem(initialItem)
}