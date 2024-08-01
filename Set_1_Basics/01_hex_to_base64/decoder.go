package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

const encodedHexData = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
const resp = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

func main() {

	decodedData, err := hex.DecodeString(encodedHexData)
	if err != nil {
		fmt.Println("Error decoding data: ", err)
	}
	fmt.Println("Decoded data: ", string(decodedData))

	encodedBase64Data := base64.StdEncoding.EncodeToString(decodedData)

	fmt.Println("Encoded base64 data: ", string(encodedBase64Data))

	if encodedBase64Data == resp {
		fmt.Println("Nicely Done!")
	} else {
		fmt.Println("Try Again, Pal.")
	}

}
