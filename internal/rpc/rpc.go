package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg)

	if err != nil {
		// TODO: hier könnte ihre exception stehen
		panic(err)
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

type BaseMessage struct {
	Method string `json:"method"`
}

func DecodeMessage(msg []byte) (int, error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return 0, errors.New("Did not find separator")
	}

	// Content-Length: <number>
	// move to the end of this slice after Content of length
	contentLenghtBytes := header[len("Content-Length: "):]
	contentLenth, err := strconv.Atoi(string(contentLenghtBytes))
	if err != nil {
		return 0, err
	}

	// TODO: remove me after test
	_ = content

	var baseMessage BaseMessage
	if err := json.Unmarshal(&baseMessage); err != nil {
		return 0, err
	}

	// https://www.youtube.com/watch?v=YsdlcQoHqPY 15:27

	return contentLenth, nil
}
