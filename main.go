package main

import (
	"os"
)

func main() {
	init := Init{}
	init.Initialize()
	init.Run(":" + os.Getenv("port"))
}
