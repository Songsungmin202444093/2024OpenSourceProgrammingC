package main

import (
	"fmt"
	"reflect"
) // c언어의 #include <stdio.h> 역할

func main() {
	var i int = 13
	f := 12.9
	fmt.Printf("value i : %d, vlue f : %f\n", i, f)
	//fmt.Printf("%d * %f = %f\n", i, f, i*f)
	fmt.Printf("%d * %f = %f\n", i, f, float64(i)*f) // conversion
	fmt.Println(reflect.TypeOf(i))

}
