/**
 * @Author: lenovo
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2023/04/13 23:16
 */
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.

type X struct {
	a int
}

type Y struct {
	X
}

func (x *X) Set(val int) {
	x.a = val
}

func (x X) Get() int {
	return x.a
}
func main() {
	x := X{a: 1}
	y := Y{X: x}
	y.Set(4)
	fmt.Println(y.Get())

	(*Y).Set(&y, 3)

	Y.Set(y, 3)
}
