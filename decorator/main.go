package main

import (
	"decorator/cipher"
	"decorator/lzw"
	"fmt"
)

type Component interface {
	Operator(string)
}

type SendComponent struct {
}

var sentData string
var receiveData string

func (s *SendComponent) Operator(data string) {
	// Send data
	sentData = data
}

type ZipComponent struct {
	com Component
}

func (s *ZipComponent) Operator(d string) {
	data, err := lzw.Write([]byte(d))
	if err != nil {
		panic(err)
	}

	s.com.Operator(string(data))
}

type EncryptComponent struct {
	key string
	com Component
}

func (s *EncryptComponent) Operator(d string) {
	data, err := cipher.Encrypt([]byte(d), s.key)
	if err != nil {
		panic(err)
	}
	s.com.Operator(string(data))
}

type DecryptComponent struct {
	key string
	com Component
}

func (s *DecryptComponent) Operator(d string) {
	data, err := cipher.Decrypt([]byte(d), s.key)
	if err != nil {
		panic(err)
	}
	s.com.Operator(string(data))
}

type UnzipComponent struct {
	com Component
}

func (s *UnzipComponent) Operator(d string) {
	data, err := lzw.Read([]byte(d))
	if err != nil {
		panic(err)
	}
	s.com.Operator(string(data))
}

type ReadComponent struct{}

func (s *ReadComponent) Operator(d string) {
	receiveData = d
}

func main() {
	// 암호화 + 압축
	sender := &EncryptComponent{
		key: "abcde",
		com: &ZipComponent{
			com: &SendComponent{},
		},
	}

	sender.Operator("Hello World")
	fmt.Println(sentData)

	receiver := &UnzipComponent{
		com: &DecryptComponent{
			key: "abcde",
			com: &ReadComponent{},
		},
	}

	receiver.Operator(sentData)

	fmt.Println(receiveData)
	// 압축만
	sender2 := &ZipComponent{
		com: &SendComponent{},
	}

	sender2.Operator("Hello World")
	fmt.Println(sentData)

	receiver2 := &UnzipComponent{
		com: &ReadComponent{},
	}

	receiver2.Operator(sentData)

	fmt.Println(receiveData)
}
