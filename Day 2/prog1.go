package main

import (
	"fmt"
	"sync"
)

func main() {
	var strings []string

	strings = append(strings, "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog")

	char := make(chan byte)
	done := make(chan bool)

	var wg sync.WaitGroup

	for i := 0; i < len(strings); i++ {
		wg.Add(1)
		s := strings[i]
		go func() {
			for j := 0; j < len(s); j++ {
				char <- s[j]
			}
			wg.Done()
		}()
	}

	freq := make(map[byte]int)

	go func() {
		for {
			select {
				case <- done:
					break
				case key := <- char:
					freq[key] += 1
			}
		}
	}()

	wg.Wait()
	done <- true

	for key, val := range freq {
		fmt.Printf("%c : %v\n", key, val)
	}

	fmt.Println("Length of map is:", len(freq))
}
