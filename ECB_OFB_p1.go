package main

import (
	"fmt"
)

// electronic code book
// H = 0100 1000 ASCII BIN Value
var message = [4]int{0b01, 0b00, 0b10, 0b00}
var Codebook = [4][2]int{{0b00, 0b01}, {0b01, 0b10}, {0b10, 0b11}, {0b11, 0b00}}
var Cipher = [4]int{}

func codebookLookup(xor int) (lookupValue int) {
	var i, j int = 0, 0
	for i = 0; i < 4; i++ {
		if Codebook[i][j] == xor {
			j++
			lookupValue = Codebook[i][j]
			break
		}
	}
	return lookupValue
}
func codebookLookupByValue(xor int) (lookupValue int) {
	var i, j int = 0, 1
	for i = 0; i < 4; i++ {
		if Codebook[i][j] == xor {
			lookupValue = Codebook[i][j-1]
			break
		}
	}
	return lookupValue
}

// var lookupValue int = 0
func main() {

	for i := 0; i < 4; i++ {
		fmt.Printf("The plaintext value is %02b\n", message[i])
	}
	for i := 0; i < 4; i++ {
		lookupValue := codebookLookup(message[i])
		fmt.Printf("The ciphertext value is %02b\n", lookupValue)
		Cipher[i] += lookupValue
	}
	for i := 0; i < 4; i++ {
		lookupValue := codebookLookupByValue(Cipher[i])
		fmt.Printf("The plaintext value is %02b\n", lookupValue)
	}

}

//Password exchange: Output Feedback Mode
//block cipher is continuous stream

//if x > 25 (subtract 25)
