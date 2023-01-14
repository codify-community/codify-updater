package utils

import (
	"errors"
	"os"
)

func Try(err error) {
	if err != nil {
		panic(err)
	}
}

func GetEnviromentVariable(item string) (string, error) {
	res := os.Getenv(item)
	if res != "" {
		return res, nil
	} else {
		return "", errors.New("no env var named: " + item)
	}
}
