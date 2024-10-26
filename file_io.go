package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func createPath(onef bool) {
	var dirName string
	if onef {
		dirName = "./onef-primes"
	} else {
		dirName = "./primes"
	}

	err := os.RemoveAll(dirName)

	if err != nil {
		log.Printf("Error removing directory '%s': %s", dirName, err)
	}

	err = os.Mkdir(dirName, os.FileMode(0775))

	if err != nil {
		log.Printf("Error creating directory '%s': %s", dirName, err)
	}
}

// Can be rewritten to remove lensub1, only done to reduce wasted len()
func writeToFile(primes []int, lensub1 int) {
	fileName := "./onef-primes/onef-primes.txt"

	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.FileMode(0644))

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	writ := bufio.NewWriter(f)

	fmt.Fprintln(f, "Primes in this file:", lensub1+1)

	fmt.Fprintln(f, fmt.Sprint("Largest: ", primes[lensub1], "\n"))

	for i := range primes {
		fmt.Fprint(writ, fmt.Sprint(primes[i], ", "))
	}

	writ.Flush()
}

func writeToFiles(order int, rangePrimes []int, wg *sync.WaitGroup) {
	defer wg.Done()

	fileName := fmt.Sprint("primes/primes-" + strconv.Itoa(order) + ".txt")

	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.FileMode(0644))

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	writ := bufio.NewWriter(f)

	fmt.Fprintln(f, "Primes in this file:", len(rangePrimes))

	fmt.Fprintln(writ, fmt.Sprint("Largest: ", rangePrimes[len(rangePrimes)-1], "\n"))

	for i := range rangePrimes {
		fmt.Fprint(writ, fmt.Sprint(rangePrimes[i], ", "))
	}

	writ.Flush()
}
