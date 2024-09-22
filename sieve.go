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
	searchLim := math.Sqrt(float64(n)) + 1

	primeBool := make([]bool, n+1)

	var primes = []int{}

	for i := 0; i <= n; i++ {
		primeBool[i] = true
	}

	for p := 2; p < int(searchLim)+1; p++ {
		if primeBool[p] {
			for i := p * p; i <= n; i += p {
				primeBool[i] = false
			}
		}
	}

	for i := 2; i < len(primeBool); i++ {
		if primeBool[i] {
			primes = append(primes, i)
		}
	}

	return primes
}

func writeToFile(primes []int) {
	f, err := os.OpenFile("primes/all-primes", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.FileMode(0644))

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	writ := bufio.NewWriter(f)

	fmt.Fprintln(f, "Primes in this file:", len(primes))

	fmt.Fprintln(f, fmt.Sprint("Largest: ", primes[len(primes)-1], "\n"))

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
	os.RemoveAll("./primes")
	os.Mkdir("./primes", os.FileMode(0775))
}

func setAndParseFlags() (int, int, bool, bool) {
	checkUpTo := flag.Int("n", 1000000000, "Assigns the upper bound of numbers you want to search")
	primesInAFile := flag.Int("inaf", 5000000, "Sets how many primes are in a file")
	noStats := flag.Bool("nostats", false, "Determines whether you want stats to be shown at the end of the search")
	oneFile := flag.Bool("onef", false, "States whether you want the primes in one file or in many files")
	flag.Parse()

	return *checkUpTo, *primesInAFile, *noStats, *oneFile
}

func main() {
	n, primesInAFile, noStats, oneFile := setAndParseFlags()

	start := time.Now()

	if n < 2 {
		fmt.Println("Invalid upper bound, no primes can be found in that range")
		return
	} else if n == 2 {
		fmt.Println("2 is not valid upper bound, there are no integers between 2 and 2")
	}

	primes := sieve(n)

	sieveTime := time.Since(start)

	createPath()

	var filesCreated int

	l := len(primes) - 1

	if oneFile {
		writeToFile(primes)
		filesCreated = 1
	} else {
		var rangePrimes []int

		var wg sync.WaitGroup

		if len(primes)%primesInAFile == 0 {
			filesCreated = len(primes) / primesInAFile //Edge case in which there are exactly (primes in a file) primes in the last file
		} else {
			filesCreated = len(primes)/primesInAFile + 1
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
		fmt.Println("Time taken for sieve: ", sieveTime.Seconds(), "seconds")
		fmt.Println("Time taken for file I/O: ", end-sieveTime.Seconds(), "seconds")
		fmt.Println("Primes found: ", len(primes))
		fmt.Println("Numbers searched: ", n)
		fmt.Println("Files created: ", filesCreated)
		fmt.Println("Primes in a file: ", primesInAFile)
		fmt.Println("Largest prime found: ", primes[l])
	}
}
