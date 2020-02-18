package main

import "fmt"

func foo(options ...bool) {
	//fmt.Println(len(params))
	//if len(params) > 0 {
	//	fmt.Println(params[0])
	//}
	if len(options) > 0 && options[0] == true {
		fmt.Println(options[0])
	}

}

func main() {
	foo()
	foo(true)
	foo(false, true)
}
