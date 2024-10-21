package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"sync"
	"time"
)

func sieve(n int) []int {
	//Returns the upper limit of prime numbers to search for
	searchLim := math.Sqrt(float64(n)) + 1

	primeBool := make([]bool, n+1)
	//2 is prime, easier to implement skipping even numbers later
	var primes = []int{2}

	for i := 3; i <= n; i += 2 {
		primeBool[i] = true
	}

	for p := 2; p < int(searchLim)+1; p++ {
		if primeBool[p] {
			for i := p * p; i <= n; i += p {
				primeBool[i] = false
			}
		}
	}
	//Implementation easier here
	for i := 3; i < n+1; i += 2 {
		if primeBool[i] {
			primes = append(primes, i)
		}
	}

	return primes
}

// Can be rewritten to remove lensub1, only done to reduce wasted len()
func writeToFile(primes []int, lensub1 int) {
	f, err := os.OpenFile("primes/all-primes", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.FileMode(0644))

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

func createPath() {
	err := os.RemoveAll("./primes")

	if err != nil {
		log.Printf("Error removing directory 'primes': %s", err)
	}

	err = os.Mkdir("./primes", os.FileMode(0775))

	if err != nil {
		log.Printf("Error creating directory 'primes': %s", err)
	}
}

func setAndParseFlags() (int, int, bool, bool) {
	checkUpTo := flag.Int("n", 1000000000, "Assigns the upper bound of numbers you want to search")
	primesInAFile := flag.Int("inaf", 5000000, "Sets how many primes are in a file")
	noStats := flag.Bool("nostats", false, "Determines whether you want stats to be shown at the end of the search")
	oneFile := flag.Bool("onef", false, "Determines whether you want the primes in one file or in many files")
	flag.Parse()

	return *checkUpTo, *primesInAFile, *noStats, *oneFile
}

func main() {
	n, primesInAFile, noStats, oneFile := setAndParseFlags()

	start := time.Now()

	if n < 2 {
		fmt.Println("Invalid upper bound, no positive primes can be found in that range")
		return
	} else if n == 2 {
		fmt.Println("2 is not valid upper bound, there are no integers between 2 and 2")
		return
	}

	primes := sieve(n)

	sieveTime := time.Since(start).Seconds()

	createPath()

	var filesCreated int

	//Reduces wasted computation (recalculating len is unecessary)
	l := len(primes) - 1

	if oneFile {
		writeToFile(primes, l)
		filesCreated = 1
	} else {
		var rangePrimes []int

		var wg sync.WaitGroup

		if (l+1)%primesInAFile == 0 {
			filesCreated = (l + 1) / primesInAFile //Edge case in which there are exactly (primes in a file) primes in the last file
		} else {
			filesCreated = (l + 1)/primesInAFile + 1
		}

		wg.Add(filesCreated)

		for i := 0; i < l; i += primesInAFile {
			if i+primesInAFile > l {
				rangePrimes = primes[i:]
			} else {
				rangePrimes = primes[i : i+primesInAFile]
			}
			// i/primesInAFile + 1 is done so there is no such file as primes-0, which looks bad
			go writeToFiles(i/primesInAFile+1, rangePrimes, &wg)
		}

		wg.Wait()
	}

	end := time.Since(start).Seconds()

	if !noStats {
		fmt.Println("Total time elapsed: ", end, "seconds")
		fmt.Println("Time taken for sieve: ", sieveTime, "seconds")
		fmt.Println("Time taken for file I/O: ", end-sieveTime, "seconds")
		fmt.Println("Primes found: ", l+1)
		fmt.Println("Numbers searched: ", n)
		fmt.Println("Files created: ", filesCreated)
		fmt.Println("Primes in a file: ", primesInAFile)
		fmt.Println("Largest prime found: ", primes[l])
	}
}
