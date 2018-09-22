package installer

import (
	"airdb/helpers"
	"fmt"
	"os"
)

func serverInstaller() {
	fmt.Println("Configuring airdb server....")
	fmt.Print("Enter your desired server port where airdb is going to run: ")
	var port string
	_, err := fmt.Scanf("%s\n", &port)
	if err != nil {
		fmt.Println("Enter your desired port to continue error " + err.Error())
		os.Exit(103)
	}
	err = helpers.Setenv("port", port)
	if err != nil {
		fmt.Println("Exception occured while setting server port error " + err.Error())
		os.Exit(103)
	}
}
