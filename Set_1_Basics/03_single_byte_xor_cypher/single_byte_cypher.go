package main

import (
	"encoding/hex"
	"fmt"
	"regexp"
)

const hexEncoded = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func CountLetters(content []byte) (map[byte]int, int, error) {
	reg, _ := regexp.Compile("[a-z]")
	mapCounterLetter := make(map[byte]int)
	totalLetter := 0

	for _, r := range content {

		_, found := mapCounterLetter[r]
		match := reg.MatchString(string(r))

		if found && match {
			mapCounterLetter[r]++
			totalLetter++
		} else if !found && match {
			mapCounterLetter[r] = 1
			totalLetter++
		}
	}

	return mapCounterLetter, totalLetter, nil
}

func GetFrequencyMapLetterCounter(mapCounterLetter map[byte]int, totalLetter int) map[byte]float64 {
	mapFreqLetter := make(map[byte]float64)
	for key, value := range mapCounterLetter {
		mapFreqLetter[key] = float64(value) / float64(totalLetter)
	}
	return mapFreqLetter
}

func main() {
	/*content, err := os.ReadFile("frankenstein.txt")
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error reading file:", err)
		return
	}
	mapCounterLetterText, totalLetterText, err := CountLetters(content)

	mapFreqLetterText := GetFrequencyMapLetterCounter(mapCounterLetterText, totalLetterText)
	*/
	hexOriginal, _ := hex.DecodeString(hexEncoded)

	fmt.Println(hexOriginal)

	mapDecypherResp := make(map[byte][]byte)

	for cypherKey := range 256 {

		var decypherHex []byte
		for _, hexOriginalValue := range hexOriginal {
			decypherHex = append(decypherHex, hexOriginalValue^byte(cypherKey))
		}

		mapDecypherResp[byte(cypherKey)] = decypherHex
	}

	for key, value := range mapDecypherResp {
		fmt.Println("Key: ", key, " is translated as: ", string(value))
	}
}
