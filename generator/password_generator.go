package generator

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	symbols = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
)

func GeneratePassword(minLength int, numbers int, specialChars int) (string, error) {
	lettersList := make([]string, minLength)

	// build the Letters list
	LettersLength := len(letters)
	for i := 0; i < minLength; i++ {
		Letter := string(letters[rand.Intn(LettersLength)])
		if convertVowelToNumber(rand.Intn(2), numbers, Letter) {
			lettersList[i] = strconv.Itoa(rand.Intn(10))
			numbers--
		} else {
			lettersList[i] = Letter
		}
	}

	// build the Numbers list
	numbersList := make([]string, numbers)
	for i := 0; i < numbers; i++ {
		numbersList[i] = strconv.Itoa(rand.Intn(10))
	}

	// build the Symbols list
	symbolsList := make([]string, specialChars)
	symbolsLength := len(symbols)
	for i := 0; i < specialChars; i++ {
		symbolsList[i] = string(symbols[rand.Intn(symbolsLength)])
	}

	// join the lists and shuffle the resulting/joined lists
	if len(numbersList) > 0 {
		LettersAndNumbersList := append(lettersList, numbersList...)
		return shuffleAndConvertToString(append(LettersAndNumbersList, symbolsList...))
	}
	return shuffleAndConvertToString(append(lettersList, symbolsList...))

}

func convertVowelToNumber(randomNumber int, numbers int, letter string) bool {
	if randomNumber == 1 && numbers > 0 && isVowel(letter) {
		return true
	}
	return false
}

func isVowel(Letter string) bool {
	vowels := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
		"A": true,
		"E": true,
		"I": true,
		"O": true,
		"U": true,
	}
	if vowels[Letter] {
		return true
	}
	return false
}

func shuffle(slice []string) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := range slice {
		j := r.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func shuffleAndConvertToString(passwordList []string) (string, error) {
	shuffle(passwordList)
	return strings.Join(passwordList, ""), nil
}
