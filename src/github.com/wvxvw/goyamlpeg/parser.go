package goyamlpeg

import (
	"log"
)

type YamlMessage struct {

}

type Messages struct {
	messages []YamlMessage
}

func (m *Messages) Append(s string) {
	log.Printf("appending message: %s", s)
}
