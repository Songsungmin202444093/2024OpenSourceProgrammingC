package main

import (
	"fmt"
	"time"
)

// c언어의 #include <stdio.h> 역할

func main() {
	var now time.Time = time.Now()
	var month int = int(now.Month()) // var month time.Month = now.Month()
	fmt.Println(month)
}
