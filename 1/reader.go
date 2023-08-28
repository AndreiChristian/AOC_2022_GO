package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func ReadFile(fileName string) (string, error) {

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file : ", err)
		return "", err
	}

	defer file.Close()

	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, file)
	if err != nil {
		fmt.Println("Error reading file : ", err)
		return "", err
	}

	return buffer.String(), nil

}
