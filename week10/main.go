package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isPrime(n int) bool {
	// var isPrime bool = true
	if n <= 1 {
		//isPrime = false
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	} else { // odd number
		j := 3 // start number
		for j*j <= n {
			if n%j == 0 {
				//isPrime = false
				//break
				return false
			}
			fmt.Printf("%d ", j)
			j = j + 2

		}
	}
	return true
}

func main() {

	fmt.Print("Input number : ")
	in := bufio.NewReader(os.Stdin)
	i, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	i = strings.TrimSpace(i)
	n, err := strconv.Atoi(i)
	if err != nil {
		log.Fatal(err)
	}

	if isPrime(n) { // == 비교 연산 제거
		fmt.Printf("%d is prime number.", n)
	} else {
		fmt.Printf("%d is NOT prime number.", n)
	}
}
