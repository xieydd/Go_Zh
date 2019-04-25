package main

import (
    "fmt"
)

type ISpeaker interface {
    Speak()
}

type SimpleSpeaker struct {
    Message string
}

func (s *SimpleSpeaker) Speak() {
    fmt.Println("I am a speaker?", s.Message)
}

func main() {
    var speaker ISpeaker
    speaker = &SimpleSpeaker{"Hello, Interface"}
    speaker.Speak()
}
