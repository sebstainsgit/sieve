# Prime Number Sieve Tool

## Overview

This tool generates prime numbers up to a specified upper bound using the **Sieve of Eratosthenes** algorithm. It provides options to save the prime numbers in a single file or split them into multiple files. The tool also offers basic performance statistics, such as execution time and the total number of primes found.

## Features

- **Efficient Prime Calculation**: Uses the Sieve of Eratosthenes algorithm to calculate primes up to a user-specified limit.
- **File Output Options**: Saves primes either in a single file or distributes them across multiple files.
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

- **-n**: of type int, specifies the upper bound for the sieve, i.e. until what number you would like to search for primes.  
 Default = 1000000000 (1 billion)
- **-inaf**: of type int, specifies the number of primes in each file (assuming you don't use -onef). \
Default = 500000 (500 thousand)
- **-nostats**: of type bool, (but usage of the flag is the same as true, and not using it is false). States whether you want to see basic statistics on the running of the program. \
Default = false
- **-onef**: of type bool. States whether you want the primes in one file or not (not concurrent). \
Default = false

## Times to expect:
Number given being upper bound (-n), time taken being an average of 5 runs.

- **1 to 10 million** is consistently less than a second so I thought it would be unneccesary to include.

- **100 million**: 2.6 seconds.
- **1 billion**: 17.6 seconds. 
- **10 billion**: 200.2 seconds.

**Note**: for one of the previous 10 billion mean calculations, I omitted a run which took 233s in the final calculation becuase it was anomalous. Notably, that particular run was the first, so when you run the program and in what context (after doing what) likely affects the time taken quite a lot.

# My specs:

- **RAM**: 32 GB of DDR4 RAM  

- **CPU**: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz  
