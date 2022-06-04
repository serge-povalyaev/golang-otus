package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Не указана env-директория")
		return
	}

	envPath, commandWithArgs := os.Args[1], os.Args[2:]

	if len(commandWithArgs) == 0 {
		fmt.Println("Не указана команда для выполнения")
		return
	}

	stat, err := os.Stat(envPath)
	if err != nil || !stat.IsDir() {
		fmt.Println("Указанная директория не существует")
		return
	}

	envs, err := ReadDir(envPath)
	fmt.Println(envs)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	fmt.Println(envPath, commandWithArgs)

	//programName := os.Args
	//fmt.Println(programName)
}
