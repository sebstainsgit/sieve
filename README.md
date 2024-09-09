# Prime Number Sieve Tool

## Overview

This tool generates prime numbers up to a specified upper bound using the **Sieve of Eratosthenes** algorithm. It provides options to save the prime numbers in a single file or split them into multiple files. The tool also offers basic performance statistics, such as execution time and the total number of primes found.

## Features

- **Efficient Prime Calculation**: Uses the Sieve of Eratosthenes algorithm to calculate primes up to a user-specified limit.
- **File Output Options**: Save primes either in a single file or distribute them across multiple files.
- **Performance Statistics**: Optionally outputs statistics, including the time taken for computation and file I/O.
- **Concurrency**: Supports concurrent file writing for better performance when generating multiple files.

## Installation

Ensure that you have Go installed on your machine. To compile and run the program:

1. Clone the repository or download the code.
2. Navigate to the project directory.
3. Build the program:
   ```bash
   go build -o prime-sieve

## Usage

After building the binary, you can run it from the command line with the following flags:

```bash
./prime-sieve [flags]
```

# Flags:

- **-n**: of type int, specifies the upper bound for the sieve, i.e. until what number you want to search for primes.  
 Default = 1000000000 (1 billion)
- **-inaf**: of type int, specifies the number of primes in each file (assuming you don't use -onef). \
Default = 500000 (500 thousand)
- **-nostats**: of type bool, (but usage of the flag is the same as true, and not using it is false). States whether you want to see basic statistics on the running of the program. \
Default = false
- **-onef**: of type bool. States whether you want the primes in one file or not (not concurrent). \
Default = false