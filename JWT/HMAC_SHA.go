package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func HMAC_SHA() {

	secrect := []byte("my-secret")
	message := []byte("Hello World")
	h := hmac.New(sha256.New, secrect)
	h.Write(message)
	text := h.Sum(nil)
	fmt.Println("HMAC-SHA:", text)
}
