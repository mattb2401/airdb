package installer

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func RunInstaller() {
	fmt.Println("Welcome to airDB")
	fmt.Println("For a fresh install follow the steps below.")
	fmt.Println("Initializing....... \n")
	fmt.Println("Choose a deployment method")
	fmt.Println("1. Docker")
	fmt.Println("2. Supervisor")
	fmt.Print("Enter your choose(1/2): ")
	var deploymentChoice string
	_, err := fmt.Scanf("%s\n", &deploymentChoice)
	if err != nil {
		fmt.Println("Error occured while choosing method " + err.Error())
		os.Exit(103)
	}
	if deploymentChoice == "2" {
		fmt.Println("Checking for supervisor .....")
		cmd := exec.Command("/bin/bash", "-c", "supervisorctl")
		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			if err.Error() == "exit status 127" {
				fmt.Println("Error: Install supervisor.d to continue with this installation")
				os.Exit(103)
			} else {
				fmt.Println("Exception occured while checking for supervisor: " + err.Error())
				os.Exit(103)
			}
		} else {
			fmt.Println("Supervisor installed.. \n")
			dbInstaller()
			serverInstaller()
		}
	} else if deploymentChoice == "1" {

	} else {
		fmt.Println("Choose a deployment method to continue.")
		os.Exit(103)
	}
}
