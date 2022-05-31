package main

import (
	"flag"
	"fmt"
)

var (
	from, to      string
	limit, offset int64
)

func init() {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to write to")
	flag.Int64Var(&limit, "limit", 0, "limit of bytes to copy")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
}

func main() {
	flag.Parse()

	if len(from) == 0 {
		fmt.Println("Необходимо указать аргумент from")
		return
	}

	if len(to) == 0 {
		fmt.Println("Необходимо указать аргумент to")
		return
	}

	if offset < 0 {
		fmt.Println("offset должен быть положительным числом")
		return
	}

	if limit < 0 {
		fmt.Println("limit должен быть положительным числом")
		return
	}

	errorResult := Copy(from, to, offset, limit, 1)
	if errorResult != nil {
		fmt.Println(errorResult)
	}
}
