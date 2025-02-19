/*

_____________§_§§_______§__§_____§§__§_________________
____________§___§§§____§___§__§§_____§_________________
____________§_____§§§_§____§_§§______§______§§_________
_____________§_______§§_____§________§__§§§§_§_________
_____________§_______________________§§§_____§_________
_____________§______________________________§__________
___§§§§§§_____§_____________________________§__________
____§______§§§_____________________________§___________
____§§_____________________________________§§§_________
_____§________________________________________§§§§_____
______§________________________________________ _§§_____
_______§_______________________________________§_______
_______§§____________________§_______________§§________
_______§_____________§§§§§§§§§_______________§_________
______§_____________§§§§§§§§§§§_______§_______§§_______
____§§____________§§§§§§§§§§§§§§________________§§_____
___§_____________§§§§§§____§§§§§§___§___§_________§§___
__§_____________§§§§§______§§§§§§____§__§§__________§__
§§_____________§§§§__§______§§§§§§§__§§__§________§§___
§_____________§§§§__§§______§§§§§§§__§§§_§§___§§§§_____
______________§§§__§_§§_____§§§§§§§§__§§§_§§__§________
_____________§§§___§§§_____§§§§§§§§§§_§§§§§§§__§_______
_____________§§____§§§____§§§§§§§§§_§§§§§§§§§___§______
§§§__________§_____§§§§_§§§§§§________§__§§§§§___§_____
___§§_______§§§____§§_§§§§§§§_____________§§§§§__§§____
____§_______§§§§___§_§§§§§__§________§____§§§§§§§§§§§§§
___§___§____§§§§§__§§§§§§__§______§§§________§__§§§§§§§
___§___§____§§§§§§§§§§§___§_____§§_________§_§§_§§§§§§§
__§____§____§§§§§§§§§§_________§§______§§§§§_§§__§§§§§§
_§_____§____§§§§§§§§§__________§______§_____§§§__§__§§§
_§____§§____§§§§§§§§§_______________________§§§__§_____
§______§____§§§§§§§§§§§§§____________________§§§_§_____
___§§§§_____§§§§§§_______§______________§§_§§§_§§______
______§__§__§§§§§§_§§§§§_____________________§§_§§_____
______§_§§__§§§§§_§§_§§_§_____________________§§_§§____
______§_§_§_§§§§§_§_§§§§§_________________§§§____§§____
______§_§_§_§§§§§_§_§§§§_____________________§____§____
______§§§_§§§§§§§_§§§§§§_______§_____§§§________§__§___
______§§§__§_§§§§__§§§§______§_____§§_________§§§_§§§__
______§_§___§_§_§___________§_____§___§§§_____§§__§§§§§
________§__§§____§______________§§__§§__§_____§§§_§§§§§
__________§____§_§§____________§§__§___§______§_§§_§§§§
_________§§_____§_§_§_________§§_§§__§§_______§__§§§_§§
_________§_______§_§__________§_§___§________§§__§§§___
_________________§§§_________§_§__§§_________§§__§§§§__
_________________§_§§_§_____§_§_§§___________§§_§§§§§__
__________________§§§§§_____§§§§_____§______§§§§§§§§§__
___________________§§§____________§§§______§§§§§§§§§§§_
____________________§§§__§________________§§§§§§§§§§§§§
_____________________§§§_§_______________§§§§§§§§§§§§§§
______________________§_§§______________§§§§§§§§§§§§§§§
________________________§§§____________§§§§§§§§§§§§§§§§
__________________________§§§§________§§§§§§§§§§§§§§§§§
_____________________________§§§§__§§§§§§§§§§§§§§§§§§§§

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func countWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}


func countCharacters(text string, includeSpaces bool) int {
	if includeSpaces {
		return len(text)
	}
	count := 0
	for _, char := range text {
		if !unicode.IsSpace(char) {
			count++
		}
	}
	return count
}


func countSentences(text string) int {
	sentences := strings.Split(text, ".")
	count := 0
	for _, s := range sentences {
		if strings.TrimSpace(s) != "" {
			count++
		}
	}
	return count
}




func wordFrequency(text string) map[string]int {
	words := strings.Fields(strings.ToLower(text))
	frequency := make(map[string]int)
	for _, word := range words {
		word = strings.Trim(word, ".,!?;:\"()[]{}")
		if word != "" {
			frequency[word]++
		}
	}
	return frequency
}


func readFromFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var text string
	for scanner.Scan() {
		text += scanner.Text() + " "
	}
	return text, scanner.Err()
}



func main() {
	fmt.Println("Enter text (or type 'file:<filename>' tovread from a file):")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	var text string
	var err error 

	if strings.HasPrefix(input, "file:") {
		filename := strings.TrimPrefix(input, "file:")
		text, err = readFromFile(strings.TrimSpace(filename))
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
	} else {
		text = input
	}


	wordCount := countWords(text)
	charCountWithSpaces := countCharacters(text, true)
	charCountWithoutSpaces := countCharacters(text, false)
	sentenceCount := countSentences(text)
	frequency := wordFrequency(text)


	fmt.Println("\n--- Word Counter Result ---")
	fmt.Println("Word Count:", wordCount)
	fmt.Println("Character Count (with spaces):", charCountWithSpaces)
	fmt.Println("Character Count (without spaces):", charCountWithoutSpaces)
	fmt.Println("Senetence Count:", sentenceCount)
	fmt.Println("Most Frequent Words:")

	count := 0
	for word, freq := range frequency {
		fmt.Printf("%s: %d\n", word, freq)
		count++
		if count == 5 {
			break
		}
	}
}
