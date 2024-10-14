package config

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func parse() {
	r, err := os.Open(".env")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	s := bufio.NewScanner(r)

	for s.Scan() {
		name, value, err := splitStr(s.Text())
		if err != nil {
			panic(err)
		}

		setOsEnvIfNotExists(name, value)
	}
}

func splitStr(str string) (name string, value string, err error) {
	parsed := strings.Split(str, "=")

	if len(parsed) != 2 || parsed[0] == "" || parsed[1] == "" {
		return "", "", errors.New("config: invalid format of .env file")
	}

	return strings.TrimSpace(parsed[0]), strings.TrimSpace(parsed[1]), nil
}

func setOsEnvIfNotExists(name string, value string) {
	_, exists := os.LookupEnv(name)

	if !exists {
		os.Setenv(name, value)
	}
}
