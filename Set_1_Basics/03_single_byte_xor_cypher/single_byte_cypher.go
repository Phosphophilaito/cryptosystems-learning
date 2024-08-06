package main

import (
	"fmt"
	"os"
	"regexp"
)

func CountLettersFile(name string) {
	content, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	counter := make(map[byte]int)
	reg, _ := regexp.Compile("a-z")

	for _, r := range content {
		_, found := counter[r]
		match := reg.MatchString(string(r))
		if found && match {
			counter[r] += 1
		} else if !found && match {
			counter[r] = 1
		}
	}

	fmt.Println(counter)

}

func main() {
	CountLettersFile("frankenstein.txt")
}
