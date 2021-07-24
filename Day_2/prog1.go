package main

import (
	"fmt"
)

func main() {
	var strings []string

	strings = append(strings, "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog")

	char := make(chan byte)
	done := make(chan bool)

	totalChars := 0

	for i := 0; i < len(strings); i++ {
		s := strings[i]
		totalChars += len(s)
		go func() {
			for j := 0; j < len(s); j++ {
				char <- s[j]
			}
		}()
	}

	freq := make(map[byte]int)

	go func() {
		for i := 0; i < totalChars; i++ {
			freq[<-char] += 1
		}
		done <- true
	}()

	<- done

	for key, val := range freq {
		fmt.Printf("%c : %v\n", key, val)
	}

	fmt.Println("Length of map is:", len(freq))
}
