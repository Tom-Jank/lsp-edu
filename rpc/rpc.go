package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// Encode message into a desired format 
func EncodeMessage(message any) string {
    content, err := json.Marshal(message)
    if err != nil {
        panic(err)
    }

    return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

type BaseMessage struct {
    Method string `json:"method"`
}

// Decode the message and fetch method name, and content len 
func DecodeMessage(message []byte) (string, []byte, error) {
    header, content, found := bytes.Cut(message, []byte{'\r', '\n', '\r', '\n'})
    if !found {
        return "", nil, errors.New("Separator not found")    
    }
    
    contentLengthBytes := header[len("Content-Length: "):]
    contentLength, err := strconv.Atoi(string(contentLengthBytes))
    if err != nil {
        return "", nil, err
    }

    var baseMessage BaseMessage
    if err := json.Unmarshal(content[:contentLength], &baseMessage); err != nil {
        return "", nil, err
    }

    return baseMessage.Method, content[:contentLength], nil
}
