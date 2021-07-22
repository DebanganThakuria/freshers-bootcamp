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
	var totalRating int

	numOfStudents = 200
	totalRating = 0

	for i := 0; i < numOfStudents; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
			totalRating += rand.Intn(10)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Average Rating:", totalRating/numOfStudents)
}
