package main

import (
	"mygin/mgin"
)

func main() {
	r := mgin.Default()

	r.Get("/", func(context *mgin.Context) {
		context.End("hello world")
	})

	r.Run("8080")
}
