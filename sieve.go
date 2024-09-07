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
	fmt.Println("Deleting ./primes directory if exists")
	os.RemoveAll("./primes")
	fmt.Print("Creating empty ./primes directory \n\n")
	os.Mkdir("./primes", os.FileMode(0775))
}

func setAndParseFlags() (int, int, bool, bool) {
	checkUpTo := flag.Int("n", 1000000000, "Assigns the upper bound of numbers you want to search")
	primesInAFile := flag.Int("inaf", 5000000, "Dictates how many primes are in a file")
	noStats := flag.Bool("nostats", false, "Determines whether you want stats to be shown at the end of the search")
	oneFile := flag.Bool("onef", false, "States whether you want the primes in one file or in many files")
	flag.Parse()

	return *checkUpTo, *primesInAFile, *noStats, *oneFile
}

func main() {
	n, primesInAFile, noStats, oneFile := setAndParseFlags()

	start := time.Now()

	if n < 2 {
		fmt.Println("Invalid upper bound, too low")
		return
	}

	fmt.Println("Sieve started")

	primes := sieve(n)

	sieveTime := time.Since(start)

	fmt.Println("Sieve ended")

	createPath()

	var filesCreated int

	if oneFile {
		writeToFile(primes)
		filesCreated = 1
	} else {
		var rangePrimes []int

		var wg sync.WaitGroup

		if len(primes)%primesInAFile == 0 {
			filesCreated = len(primes) / primesInAFile
		} else {
			filesCreated = len(primes)/primesInAFile + 1
		}

		wg.Add(filesCreated)

		l := len(primes) - 1

		for i := 0; i < l; i += primesInAFile {
			if i+primesInAFile > l {
				rangePrimes = primes[i:]
			} else {
				rangePrimes = primes[i : i+primesInAFile]
			}
			//i+1 is done so there is no such file as primes-0, which looks bad and that
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
		fmt.Println("Largest prime found: ", primes[len(primes)-1])
	}
}