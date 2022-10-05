package util

import (
	"errors"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"

	"github.com/Korppi/love/pkg/models"
)

// String of lowercase alphabets
const Alphabets string = "abcdefghijklmnopqrstuvwxyzåäö"

// Calculates lovescore for given two names.
//
// The idea how to calculate score is taken from here:
// https://www.codebrainer.com/blog/love-calculator-android
func CalculateLove(s1, s2 string) (models.Love, error) {
	if len(s1) == 0 || len(s2) == 0 {
		return models.Love{}, errors.New("ONE OF NAMES IS EMPTY")
	}
	text := s1 + "loves" + s2
	numberString, err := stringOccurencesToString(text)
	if err != nil {
		return models.Love{}, err
	}
	for numberString != "100" && len(numberString) >= 3 {
		n := numberString
		numberString = ""
		for n != "" {
			if len(n) == 1 {
				numberString = numberString + n
				break
			}
			n1, _ := strconv.Atoi(string(n[0]))
			n2, _ := strconv.Atoi(string(n[len(n)-1]))
			sum := n1 + n2
			numberString = numberString + strconv.Itoa(sum)
			n = n[1 : len(n)-1]
		}
	}
	percentage, _ := strconv.Atoi(numberString)
	description := getDescription(percentage)
	return models.Love{Percent: percentage, Description: description}, nil
}

// Returns description for percentage
func getDescription(percentage int) string {
	switch {
	case percentage > 95:
		return "You are perfect couple!"
	case percentage > 80:
		return "You are good for each other!"
	case percentage > 70:
		return "Not perfect but you can make it!"
	case percentage > 60:
		return "There might be some problems but give it a try!"
	case percentage > 50:
		return "Not good, not bad..."
	case percentage > 40:
		return "Things are not looking good..."
	case percentage > 30:
		return "You both need to do some changes!"
	case percentage > 20:
		return "Looking really bad!"
	case percentage > 10:
		return "Just forget whole thing!"
	default:
		return "You are DOOMED!!!11"
	}
}

// Calculates how many times letters are found in given text.
//
// Example, "Aabc" returs list [2,1,1,0,0...].
// NOTE: changes given text to lowercase, so "A" and "a" will be the same.
func countCharacters(text string) ([]int, error) {
	alphabetCount := make([]int, len(Alphabets))
	lettersAsList := strings.Split(strings.ToLower(text), "")
	for _, s := range lettersAsList {
		index := strings.Index(Alphabets, s)
		if index == -1 {
			return []int{}, errors.New("ILLEGAL CHARACTER IN NAME")
		}
		alphabetCount[index]++

		if !slices.Contains(lettersAsList, s) {
			lettersAsList = append(lettersAsList, s)
		}
	}
	return alphabetCount, nil
}

// Makes set of text.
//
// Example: "ADbca" returns ["a","d","b","c"].
// NOTE: changes given text to lowercase before doing set.
func stringToSet(text string) []string {
	list := strings.Split(strings.ToLower(text), "")
	setOfLetters := make([]string, 0)
	for _, s := range list {
		if !slices.Contains(setOfLetters, s) {
			setOfLetters = append(setOfLetters, s)
		}
	}
	return setOfLetters
}

// Returns string which shows how many times letter were in text
//
// Example: "abca" return "211" because letter "a" was 2 times in text, and b and c
// only 1 time.
func stringOccurencesToString(text string) (string, error) {
	text = strings.ReplaceAll(text, "-", "")
	alphabetCount, err := countCharacters(text)
	if err != nil {
		return "", err
	}
	stringSet := stringToSet(text)
	numberString := ""
	for _, v := range stringSet {
		numberString = numberString + strconv.Itoa(alphabetCount[strings.Index(Alphabets, v)])
	}
	return numberString, nil
}
