package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
)

// function panic and recovery
// panic
//   way to abruptly exit an executing function
//   ways a function panics:
//   - calling panic built-in function
//   - abnormal source code state
//   - accessing a nil value or out-of-bound array element
//   - concurrency deadlock
// recovery
//   after function panic way to regain control of the exution flow

var (
	emptyFileNameError = errors.New("File name cannot be empty")
)

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

func load(fname string) ([]string, error) {
	if fname == "" {
		return nil, errors.New("Dictionary file name cannot be empty.")
	}

	file, err := os.Open(fname)
	if err != nil {
		return nil, fmt.Errorf("Unable to open file %s: %s", fname, err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func write(fname string, anagrams map[string][]string) {
	if anagrams == nil {
		panic("Unable to write, anagrams missing.")
	}

	file, err := os.OpenFile(fname, os.O_WRONLY+os.O_CREATE+os.O_EXCL, 0644)
	if err != nil {
		msg := fmt.Sprintf("Unable to create output file: %v", err)
		panic(msg)
	}
	defer file.Close()
	for k, v := range anagrams {
		output := fmt.Sprint("%s -> %v\n", k, v)
		file.WriteString(output)
	}
}

func mapWords(words []string) map[string][]string {
	anagrams := make(map[string][]string)
	for _, word := range words {
		wordSig := sortRunes(word)
		anagrams[wordSig] = append(anagrams[wordSig], word)
	}
	return anagrams
}

func makeAnagrams(words []string, fname string) {
	// recovery func
	//   unamed func
	//   allows program to gracefully recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Failed ot make anagram:", r)
		}
	}()

	anagrams := mapWords(words)
	write(fname, anagrams)
}

func main() {
	words, err := load("dict.txt")
	if err != nil {
		fmt.Println("Unable to load file:", err)
		os.Exit(1)
	}

	// if there is already is an out.txt file program will panic
	makeAnagrams(words, "out.txt")
}
