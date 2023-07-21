package variables

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// printvariables()
	// printVariablesType()
	constVariables()
}

var printvariables = func() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}

var printVariablesType = func() {
	println(reflect.TypeOf(42).Kind().String())
	var a = "initial"
	a = "hello"
	fmt.Println("a----->", reflect.TypeOf(a).Kind().String())
}

const s string = "constant"

var constVariables = func() {
	fmt.Println(s)
	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
