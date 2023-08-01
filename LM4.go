package main

import (
	"fmt"
	"golang.org/x/crypto/chacha20"
	"log"
)

// cipher_text = plain_text XOR chacha_stream(key, nonce)
// plain_text = cipher_text XOR chacha_stream(key, nonce)

func main() {
	//check "make"
	key := make([]byte, 32)
	nonce := make([]byte, 24)
	plainText := []byte("Hello World")
	cipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		log.Fatal(err)
	}

	ciphertext := make([]byte, len(plainText))
	cipher.XORKeyStream(ciphertext, plainText)

	fmt.Printf("Cipher text: %x\n", ciphertext)

	cipher, err = chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		log.Fatal(err)
	}

	decrypted := make([]byte, len(ciphertext))
	cipher.XORKeyStream(decrypted, ciphertext)

	fmt.Printf("Plaintext: %s\n", decrypted)
}
