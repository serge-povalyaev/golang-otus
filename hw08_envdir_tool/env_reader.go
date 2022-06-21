package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

var ErrNotDir = errors.New("указанный путь не является путем до директории")

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	stat, err := os.Stat(dir)
	if err != nil {
		return nil, err
	}

	if !stat.IsDir() {
		return nil, ErrNotDir
	}

	envList := make(Environment)
	for _, env := range os.Environ() {
		envData := strings.SplitN(env, "=", 2)

		envList[envData[0]] = EnvValue{envData[1], false}
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, fileInfo := range files {
		fileName := fileInfo.Name()
		filePath := fmt.Sprintf("%s/%s", dir, fileName)

		delete(envList, fileName)

		if fileInfo.Size() == 0 || strings.Contains(fileName, "=") {
			continue
		}

		envValue, err := readFirstLine(filePath)
		if err != nil {
			continue
		}

		envList[fileName] = EnvValue{stringProcess(envValue), false}
	}

	return envList, nil
}

func readFirstLine(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		return scanner.Text(), nil
	}

	return "", nil
}

func stringProcess(str string) string {
	str = strings.TrimRight(str, "\t\n")
	str = strings.TrimRight(str, " ")
	str = string(bytes.ReplaceAll([]byte(str), []byte("\x00"), []byte("\n")))

	return str
}
