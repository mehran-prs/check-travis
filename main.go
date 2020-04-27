package main

import (
	"fmt"
	"runtime"
)

const a = 1 << 580
//const aa = a << 112
//const b = a >> 121

func main() {
	// var a=1<<1
	// fmt.Println(a)
	// fmt.Printf("%b\n", a)
	// fmt.Println(b==2)
	fmt.Printf("Hello from go:%s \n",runtime.Version())
}

