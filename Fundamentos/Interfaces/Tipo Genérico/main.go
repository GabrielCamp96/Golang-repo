package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(true)

	mapa := map[interface{}]interface{}{
		1:    "x",
		"X":  6.515,
		true: false,
	}

	fmt.Println(mapa)
}
