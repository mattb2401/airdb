package installer

import (
	"airdb/helpers"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func serverInstaller() {
	fmt.Println("Configuring airdb server using supervisor.d ....")
	fmt.Print("Enter your desired server port where airdb is going to run: ")
	var port string
	_, err := fmt.Scanf("%s\n", &port)
	if err != nil {
		fmt.Println("Error " + err.Error())
		os.Exit(103)
	}
	err = helpers.Setenv("port", port)
	if err != nil {
		fmt.Println("Exception occured while setting server port error " + err.Error())
		os.Exit(103)
	}
	if port == "" {
		fmt.Println("Enter a valid server port ")
		os.Exit(103)
	}
	fmt.Print("Enter supervisorctl conf path:")
	var confPath string
	_, err = fmt.Scanf("%s\n", &confPath)
	if err != nil {
		fmt.Println("Error capturing conf path" + err.Error())
		os.Exit(103)
	}
	if confPath == "" {
		fmt.Println("Enter supervisorctl conf path to continue.")
		os.Exit(103)
	}
	c := confPath[len(confPath)-1:]
	if c == "/" {
		confPath = confPath[:len(confPath)-1]
	}
	path := confPath + "/airdb.conf"
	err = createConfFile(path)
	if err != nil {
		fmt.Println("Failed to create and write supervisorctl conf file error: " + err.Error())
		os.Exit(103)
	}
	fmt.Println("Reloading supervisor with new airdb conf settings...")
	cmd1 := exec.Command("/bin/bash", "-c", "supervisorctl restart")
	var outb, errb bytes.Buffer
	cmd1.Stdout = &outb
	cmd1.Stderr = &errb
	err = cmd1.Run()
	if err != nil {
		if err.Error() == "exit status 127" {
			fmt.Println("Error: Install supervisor.d to continue with this installation")
			os.Exit(103)
		} else {
			fmt.Println("Exception occured while reloading supervisor conf: " + err.Error())
			os.Exit(103)
		}
	}
	fmt.Println("Starting airdb web service...")
	cmd2 := exec.Command("/bin/bash", "-c", "supervisorctl restart airdb")
	cmd2.Stdout = &outb
	cmd2.Stderr = &errb
	err = cmd2.Run()
	if err != nil {
		if err.Error() == "exit status 127" {
			fmt.Println("Error: Install supervisor.d to continue with this installation")
			os.Exit(103)
		} else {
			fmt.Println("Exception occured while reloading supervisor conf: " + err.Error())
			os.Exit(103)
		}
	}
	if strings.Contains(outb.String(), "(no such process)") {
		fmt.Println("Failed to start airdb application. Check your supervisorctl settings")
		os.Exit(103)
	} else {
		fmt.Println("Airdb configured successfully")
		os.Exit(103)
	}

}

func createConfFile(path string) error {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			return err
		}
		err = writeConfiguration(path)
		if err != nil {
			fmt.Println("Failed to write configuration in conf file error: " + err.Error())
			os.Exit(103)
		}
		defer file.Close()
	}
	return nil
}

func writeConfiguration(path string) error {
	appPath, err := os.Getwd()
	if err != nil {
		return err
	}
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString("[program:airdb] \ncommand=" + appPath + "/airdb \ndirectory=" + appPath + "\nuser=root \nautostart=true \nautorestart=true \nredirect_stderr=true\n")
	if err != nil {
		return err
	}
	err = file.Sync()
	if err != nil {
		return err
	}
	return nil
}
