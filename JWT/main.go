package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	// 🔹 Original string
	s := "Nafiz"
	fmt.Println("Original:", s)

	// 🔹 Convert string → byte
	byteArr := []byte(s)
	fmt.Println("Bytes:", byteArr)

	// 🔹 Base64 Encoder (URL safe, no padding)
	enc := base64.URLEncoding.WithPadding(base64.NoPadding)

	// 🔐 Encode
	b64Str := enc.EncodeToString(byteArr)
	fmt.Println("Encoded (Base64):", b64Str)

	dnc, err := enc.DecodeString(b64Str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Decode:", dnc)

	// 🔹 Call SHA function (from another file or same file)
	SHA()
	HMAC_SHA()
}
