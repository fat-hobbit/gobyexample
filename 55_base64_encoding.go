package gobyexample

import b64 "encoding/base64"
import "fmt"

// Base64EncodingDemo - demonstrates base64 encoding/decoding
func Base64EncodingDemo() {
	data := "abc123!?$*&()'-=@~"

	// Go supports both standard and URL-compatible base64.
	// Here's an example of encoding using the standard encoder.
	// The encoder requires a []byte so we cast our string to that.
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// Decoding may return an error, which you can check if you don't
	// already know the error to be well-formed.
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// This encodes / decodes using a URL-compatible base64 format.
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
