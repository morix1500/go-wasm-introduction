package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Cli struct {
	pattern  string
	fileName string
}

func start(cli Cli) error {
	file, err := os.Open(cli.fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, cli.pattern) {
			fmt.Println(line)
		}
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <pattern> <file_name>")
		return
	}

	cli := Cli{
		pattern:  os.Args[1],
		fileName: os.Args[2],
	}

	if err := start(cli); err != nil {
		panic(err)
	}
}
