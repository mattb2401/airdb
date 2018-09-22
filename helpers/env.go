package helpers

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const env_file = "./.env"

func Setenv(key string, value string) error {
	var text string
	text = fmt.Sprintf("%s=%s\n", key, value)
	f, err := os.OpenFile(env_file, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.WriteString(text); err != nil {
		return err
	}
	return nil
}

func Getenv(key string) (string, error) {
	f, err := os.Open(env_file)
	if err != nil {
		return "", err
	}
	defer f.Close()
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for _, ln := range lines {
		l := strings.Split(string(ln), "=")
		if l[0] == key {
			return l[1], nil
		}
	}
	return "", errors.New("false")
}
