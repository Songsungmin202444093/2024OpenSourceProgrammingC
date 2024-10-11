package main

import (
	"bufio"
	"fmt"
	"os"
)

// c언어의 #include <stdio.h> 역할

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Print("Input your name : ")
	name, err := in.ReadString('\n')
	// fmt.Println(i, err)
	fmt.Println(name)
	fmt.Println(err)
}
