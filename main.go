package main

import (
	"fmt"
	"strings"
)

const originalLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func hashLetterFn(key int, letter string) (result string) {
	runes := []rune(letter)
	lastLetterKey := string(runes[len(letter)-key : len(letter)])
	leftOverLetters := string(runes[0 : len(letter)-key])
	return fmt.Sprintf("%s%s", lastLetterKey, leftOverLetters)
}

func encrypt(key int, plainText string) (result string) {
	hashLetter := hashLetterFn(key, originalLetters)
	hashedString := ""
	findOne := func(r rune) rune {
		pos := strings.IndexRune(originalLetters, r)
		if pos != -1 {
			letterPosition := (pos + len(originalLetters)) % len(originalLetters)
			hashedString += string(hashLetter[letterPosition])
		}
		return r
	}

	strings.Map(findOne, plainText)
	return hashedString
}

func decrypt(key int, encryptedText string) (result string) {
	hashLetter := hashLetterFn(key, originalLetters)
	hashedString := ""
	findOne := func(r rune) rune {
		pos := strings.IndexRune(hashLetter, r)
		if pos != -1 {
			letterPosition := (pos + len(originalLetters)) % len(originalLetters)
			hashedString += string(originalLetters[letterPosition])
		}
		return r
	}

	strings.Map(findOne, encryptedText)
	return hashedString
}

func main() {
	plainText := "HELLOWORLD"
	fmt.Println("Plain Text: ", plainText)
	encryptedText := encrypt(5, plainText)
	fmt.Println("Encrypted Text: ", encryptedText)
	decryptedText := decrypt(5, encryptedText)
	fmt.Println("Decrypted Text: ", decryptedText)
}
