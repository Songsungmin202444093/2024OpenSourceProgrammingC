package main

import (
	"fmt"
	"reflect"
) // c언어의 #include <stdio.h> 역할

func main() {
	var i int = 13
	f := 12.9
	c1 := 'Z' // 90
	c2 := '김' // 44608, Unicode

	fmt.Printf("value i : %d, vlue f : %f\n", i, f)
	//fmt.Printf("%d * %f = %f\n", i, f, i*f)
	//fmt.Printf("%d * %f = %f\n", i, f, float64(i)*f) // conversion
	fmt.Printf("%d * %f = %d\n", i, f, i*int(f)) // conversion
	fmt.Println(c1, c2)
	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f), reflect.TypeOf(c1), reflect.TypeOf(c2))

}
