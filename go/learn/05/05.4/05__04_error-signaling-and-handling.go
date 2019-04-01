package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
)

// Error signaling and handling
//   exception handling within code when an undesirable state arises
//   handle errors immediately after a called func returns with type error

// signaling errors
//   builtin type error signals error(s) from a func
//   func must return a value with type error to indicate something went wrong

// sortRunes is a simple insertion sort that sorts the runes of the given string.
// By sorting the unicode chars of the string (i.e. "morning" -> "gimnnor")
// the result can be used as a signature to create words in the same class.
func sortRunes(str string) string {
	runes := bytes.Runes([]byte(str))
	var temp rune
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[j] < runes[i] {
				temp = runes[i]
				runes[i], runes[j] = runes[j], temp
			}
		}
	}
	return string(runes)
}

// load loads the content of the specified file's name into
// memory as an slice (array) of strings.  Notice the function
// also returns an error value that will be non-nil if an error
// occured during the load process.
// Scanner.Split() uses a function that knows
// to split the file records.  Here the code uses a provided function bufio.ScanLines
// to do the job of splitting the records from the file and returns a word
// for each line read and saves in slice of words.
func load(fname string) ([]string, error) {
	if fname == "" {
		// func signature has 2 values
		//   return value nil for type string
		//   return value errors.New for type error
		//  errors.New is part of builtin error handling type
		//  fmt.Errorf lets you parametize error values
		return nil, errors.New("Dictionary file name cannot be empty")
	}

	// attempt to load dict file.
	file, err := os.Open(fname)
	if err != nil {
		// func signature has 2 values
		//   return value nil for type string
		//   return value err from os.Open
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	// func signature has 2 values
	//   return value lines for type string
	//   return value scanner.Err() for type error
	return lines, scanner.Err()
}

func main() {
	// loads dictionary words in memory
	words, err := load("dict.txt")

	// test err for non-nil value.
	//   programmer knows where errors is from in execution
	//   programmers must always test for error
	//   forces programmer to create robust error checking
	// If err != nil, then the error is handled.
	//   value nil type error means non-erroneus state
	// In this context, the program ends.
	if err != nil {
		fmt.Println("Unable to load file:", err)
		// handle error by terminating program
		os.Exit(1)
	}

	// the next code snippet goes through the words slice
	// and does the followings for each word in the slice:
	// -  Creates a word signature by soriting the letters in the word itself.
	// -  Maps the word to its signature
	// So, words with the same signature will be mapped together (i.e. abers -> sabre, bares, bears, saber, etc)
	anagrams := make(map[string][]string)
	for _, word := range words {
		wordSig := sortRunes(word)
		anagrams[wordSig] = append(anagrams[wordSig], wordSig)
	}

	// prints the results
	for k, v := range anagrams {
		fmt.Println(k, "->", v)
	}
}
