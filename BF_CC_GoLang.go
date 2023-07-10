package main

import "fmt"

func main() {
	cipherText := "Ugew gnwj zwjw Oslkgf"

	for i := 0; i < 26; i++ {
		plainText := ""
		for _, char := range cipherText {
			if char >= 'a' && char <= 'z' {
				nChar := rune((int(char-'a')-i+26)%26 + 'a')
				plainText += string(nChar)
			} else if char >= 'A' && char <= 'Z' {
				nChar := rune((int(char-'A')-i+26)%26 + 'A')
				plainText += string(nChar)
			} else {
				plainText += string(char)
			}
		}
		fmt.Printf("Iteration %d: %s\n", i, plainText)
	}
}
