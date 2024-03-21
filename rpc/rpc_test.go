package rpc

import (
	"testing"
)

type EncodingExample struct {
    TestCase bool
}

func TestEncode(t *testing.T) {
    expected := "Content-Length: 17\r\n\r\n{\"TestCase\":true}"
    actual := EncodeMessage(EncodingExample{TestCase: true})

    if expected != actual {
        t.Fatalf("Expected: %s, Actual: %s", expected, actual)
    }
}

func TestDecode(t *testing.T) {
    incomingMessage := "Content-Length: 15\r\n\r\n{\"Method\":\"Hi\"}"

    method, content, err := DecodeMessage([]byte(incomingMessage))
    contentLength := len(content)
    
    if err != nil {
        t.Fatal(err)
    }
    if contentLength != 15 {
        t.Fatalf("Expected: 15, Actual: %d", contentLength)
    }
    if method != "Hi" {
        t.Fatalf("Expected: 'hi'), Actual: %s", method)
    }
}

