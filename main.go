package main

import (
	"fmt"
	"math/rand"
	//"slices"
	"strings"

	"github.com/fatih/color"
)

func main() {
	// 1. Get a randon word from list of words
	words := []string{  
		"Happy",
		"Sunny",
		"Jump",
		"Laugh",
		"Friend",
		"Book",
		"Blue",
		"Cat",
		"Rain",
		"Small",
		"Run",
		"Sweet",
		"Green",
		"Tree",
		"Sing",
		"Dog",
		"Warm",
		"Home",
		"Play",
		"Star",
	}
	word := strings.ToLower(words[rand.Intn(len(words))])

	lives := 2 * len(word)

	// 2. Generate the word blanks
	blanks := []string{}
	for range word {
		blanks = append(blanks, "_")
	}

	for {
		// 3. Show the word blanks and ask user for letters
		fmt.Printf("❤️ %d, Word: %s - Letter: ", lives, strings.Join(blanks, ""))
		var input string
		fmt.Scanln(&input)
		input = strings.ToLower(input)

		// 4. Check provided letters
		for _, inputLetter := range input {
			correctGuess := false

			for i, wordLetter := range word {
				if inputLetter == wordLetter {
					correctGuess = true
					color.Blue("%s was correct guess!\n", string(inputLetter))
					// if slices.Contains(blanks, string(inputLetter)) {
					// 	color.Yellow("%s has already been guessed!\n", string(inputLetter))
					// }
					blanks[i] = string(inputLetter)
				}
			}

			if !correctGuess {
				color.Red("%s was wrong guess!\n", string(inputLetter))
				lives--
			}
		}

		// 5. If no more lives, you lost
		if lives <= 0 {
			color.Red("❤️ 0, Word: %s - Sorry you lost!\n", word)
			break
		}

		// 6- if guess correct
		if word == strings.Join(blanks, "") {
			color.Green("❤️ %d, Word: %s - You won, Congarts!\n", lives, word)
			break
		}
	}
	
}
