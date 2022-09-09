package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	// get command line arguments
	args := os.Args
	x, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}

	// generate x random numbers and store in slice
	randslice := make([]int, x)
	for i := 0; i < x; i++ {
		randslice[i] = rand.Int() % 1000
	}
	fmt.Println(randslice)

	// find sum, print sum and time took to find sum
	start := time.Now()
	total := sum(randslice)
	elapsed := time.Since(start)
	fmt.Println(total, elapsed)

	// find sum using two goroutines (NOTE: main is considered a goroutine)
	ch := make(chan int)
	start = time.Now()
	go func() {
		ch <- sum(randslice[:len(randslice)/2])
	}()
	v := <-ch
	total = (sum(randslice[len(randslice)/2:]) + v)
	elapsed = time.Since(start)
	fmt.Println(total, elapsed)

}

func sum(slice []int) int {
	var total int
	for i := 0; i < len(slice); i++ {
		total += slice[i]
	}
	return total
}
