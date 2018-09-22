package main

import (
	"airdb/helpers"
	"airdb/installer"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) >= 2 {
		if os.Args[1] == "-i" {
			installer.RunInstaller()
		} else {
			fmt.Println("Unknown argument")
			os.Exit(103)
		}
	} else {
		init := Init{}
		init.Initialize()
		port, err := helpers.Getenv("port")
		if err != nil {
			fmt.Println("Error occurred while determining application port. Running airdb -install first to configure the application first")
			os.Exit(103)
		}
		init.Run("0.0.0.0:" + port)
	}
}
