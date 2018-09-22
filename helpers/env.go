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
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if len(lines) > 0 {
		for _, ln := range lines {
			l := strings.Split(string(ln), "=")
			if l[0] != key && l[1] != value {
				if _, err = f.WriteString(text); err != nil {
					return err
				}
			}
		}
	} else {
		if _, err = f.WriteString(text); err != nil {
			return err
		}
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

func Flushenv() error {
	var err = os.Remove(env_file)
	if err != nil {
		return err
	}
	_, err = os.Stat(env_file)
	if os.IsNotExist(err) {
		var file, err = os.Create(env_file)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	return nil
}
