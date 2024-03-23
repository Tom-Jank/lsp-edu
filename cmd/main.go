package main

import (
	"bufio"
	"log"
	"os"

	"github.com/Tom-Jank/lsp-edu/rpc"
)

func main() {
    logger := getLogger("/home/tomek/Dokumenty/Projects/lsp/log.txt")
    logger.Println("Started...")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(rpc.Split)
    
    for scanner.Scan() {
        message := scanner.Text()
        handleMessage(logger, message)
    }
}

func handleMessage(logger *log.Logger, message any) {
    logger.Println(message)
}

func getLogger(filename string) *log.Logger {
    logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
    if err != nil {
        panic("Wrong file passed")
    }
    return log.New(logfile, "[lsp-edu]", log.Ldate|log.Ltime|log.Lshortfile)
}
