package Pair1_IST402

import (
	"bufio"
	"fmt"
	"os"
)

package main

import (
"bufio"
"fmt"
"os"

"golang.org/x/crypto/chacha20"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the string to encrypt: ")
	plainText, _ := reader.ReadString('\n')

	key := []byte("supersecretpassword") // Replace this with your actual secret key

	// Encrypt the input string using ChaCha20
	encrypted, nonce := encrypt([]byte(plainText), key)

	// Decrypt the encrypted data back to the original form
	decrypted := decrypt(encrypted, key, nonce)

	fmt.Println("Encrypted:", string(encrypted))
	fmt.Println("Decrypted:", string(decrypted))
}

func encrypt(plainText, key []byte) ([]byte, []byte) {
	cipher, _ := chacha20.NewUnauthenticatedCipher(key, make([]byte, chacha20.NonceSize))
	nonce := make([]byte, chacha20.NonceSize)

	// Generate a random nonce for encryption
	if _, err := cipher.XORKeyStream(nonce, nonce); err != nil {
		panic(err)
	}

	encrypted := make([]byte, len(plainText))
	cipher.XORKeyStream(encrypted, plainText)

	return encrypted, nonce
}

func decrypt(encrypted, key, nonce []byte) []byte {
	cipher, _ := chacha20.NewUnauthenticatedCipher(key, nonce)
	decrypted := make([]byte, len(encrypted))
	cipher.XORKeyStream(decrypted, encrypted)
	return decrypted
}
