package main

import (
	"crypto/sha256"
	"fmt"
)

func SHA() {

	data := []byte("Hello")
	fmt.Println("Data:", data)
	hash := sha256.Sum256(data)
	fmt.Println("Hash after SHA-256:", hash)

}
