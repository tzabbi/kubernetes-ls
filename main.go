package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/tzabbi/kubernetes-ls/internal/rpc"
)

func main() {
	logger := getLogger("/home/tom/Documents/github/tzabbi/kubernetes-ls/log.txt")
	logger.Println("Hey I started")
	fmt.Println("hi")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error: %s", err)
		}
		handleMessage(logger, method, contents)
	}
}

func handleMessage(logger *log.Logger, method string, contents []byte) {
	logger.Printf("Received msg with method: %s", method)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)

	if err != nil {
		panic("hey, you didnt give me a good file!")
	}

	return log.New(logfile, "[kubernetes-ls]", log.Ldate|log.Ltime|log.Lshortfile)
}
