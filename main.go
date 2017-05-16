package main

import (
	"fmt"
	"mygin/mgin"
)

func main() {
	r := mgin.Default()

	r.Get("/", func(context *mgin.Context) {
		context.End("hello world")
	})
	str := "hello"
	fmt.Printf("%b", str[0])
	fmt.Println(string(str[:3]))
	fmt.Println(str[0] == 'h')

	r.Run("8080")
}
