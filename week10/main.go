package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	//fmt.Printf("%f\n", math.Sqrt(19.0))
	fmt.Print("Input number : ")
	in := bufio.NewReader(os.Stdin)
	i, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	i = strings.TrimSpace(i) // 줄바꿈등 제거. 파이썬의 strip 함수와 비슷
	n, err := strconv.Atoi(i)
	if err != nil {
		log.Fatal(err)
	}

	var isPrime bool = true // 가독성 up, 메모리 down

	if n <= 1 {
		isPrime = false
	} else {
		j := 2
		for j <= int(math.Sqrt(float64(n))) {
			if n%j == 0 {
				isPrime = false // 더하기 연산 제거
				break           // performance up
			}
			fmt.Printf("%d ", j) //check j loop
			j++

		}
	}

	if isPrime { // == 비교 연산 제거
		fmt.Printf("%d is prime number.", n)
	} else {
		fmt.Printf("%d is NOT prime number.", n)
	}
}
