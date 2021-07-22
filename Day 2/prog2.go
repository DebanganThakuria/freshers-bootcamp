package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var numOfStudents int
	var wg sync.WaitGroup
	var m sync.Mutex
	var totalRating int

	numOfStudents = 200
	totalRating = 0

	for i := 0; i < numOfStudents; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
			m.Lock()
			totalRating += rand.Intn(10)
			m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Average Rating:", totalRating/numOfStudents)
}
