package main

import (
    "bufio"
    "encoding/json"
    "log"
    "os"

    "github.com/Tom-Jank/lsp-edu/lsp"
    "github.com/Tom-Jank/lsp-edu/rpc"
)

func main() {
    logger := getLogger("/home/tomek/Dokumenty/Projects/lsp/log.txt")
    logger.Println("Started...")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(rpc.Split)

    for scanner.Scan() {
        message := scanner.Bytes()
        method, contents, err := rpc.DecodeMessage(message)
        if err != nil {
            logger.Printf("Gon an error: %s", err)
        }
        handleMessage(logger, method, contents)
    }
 }

func handleMessage(logger *log.Logger, method string, contents []byte) {
    logger.Printf("Received msg with method %s: ", method)

    switch method {
    case "initizialize":
        var request lsp.InitializeRequest
        if err := json.Unmarshal(contents, &request); err != nil {
            logger.Printf("Could not parse this %s: ", method)
        }
        logger.Printf("Connected to: %s, %s",
            request.Params.ClientInfo.Name,
            request.Params.ClientInfo.Version)
    }
}

func getLogger(filename string) *log.Logger {
    logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
    if err != nil {
        panic("Wrong file passed")
    }
    return log.New(logfile, "[lsp-edu]", log.Ldate|log.Ltime|log.Lshortfile)
}
