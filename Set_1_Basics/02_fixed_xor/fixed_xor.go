package main

import (
	"encoding/hex"
	"fmt"
)

const dataOne = "1c0111001f010100061a024b53535009181c"
const dataTwo = "686974207468652062756c6c277320657965"
const dataResponse = "746865206b696420646f6e277420706c6179"

func main() {
	hexOne, _ := hex.DecodeString(dataOne)
	hexTwo, _ := hex.DecodeString(dataTwo)

	response := make([]byte, len(hexOne))

	for i := range hexOne {
		response[i] = hexOne[i] ^ hexTwo[i]
	}

	hexResponse := hex.EncodeToString([]byte(response))

	if hexResponse == dataResponse {
		fmt.Println("Nicely Done!")
	} else {
		fmt.Println("BAD!, your response is incorrect: ", hexResponse)
	}
}
