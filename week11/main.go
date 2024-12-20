package keyborad

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"week11/keyboard"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	} else {
		j := 3
		for j*j <= n {
			if n%j == 0 {
				return false
			}
			// fmt.Printf("%d ", j)
			j = j + 2

		}
	}
	return true
}

func getIntager() (int, error) {
	in := bufio.NewReader(os.Stdin)
	a, err := in.ReadString('\n')

	if err != nil {
		//log.Fatal(err)
		return 0, err
	}
	a = strings.TrimSpace(a)
	n, err := strconv.Atoi(a)
	if err != nil {
		return 0, err
	}
	return n, err
}

func main() {
	fmt.Print("Input strat number : ")
	n1, err := keyboard.GetIntager()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Input end number : ")
	n2, err := keyboard.GetIntager()
	if err != nil {
		log.Fatal(err)
	}
	for j := n1; j <= n2; j++ {
		if isPrime(j) {
			fmt.Printf("%d ", j)
		}
	}
}
