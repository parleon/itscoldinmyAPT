package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
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

	// time sort slice, print elapsed time
	start := time.Now()
	sort.Slice(randslice, func(i int, j int) bool {return randslice[i] < randslice[j]})
	elapsed := time.Since(start)
	fmt.Println(elapsed)

	// regenerate random numbers
	for i := 0; i < x; i++ {
		randslice[i] = rand.Int() % 1000
	}

	// time sort slicestable, print elapsed time
	start = time.Now()
	sort.SliceStable(randslice, func(i int, j int) bool {return randslice[i] < randslice[j]})
	elapsed = time.Since(start)
	fmt.Println(elapsed)


}
