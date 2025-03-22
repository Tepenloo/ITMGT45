package main

import (
	"fmt"
	"strings"
)

// main
func main() {
	// Testing some examples from template
	fmt.Println(shiftLetter("A", 2))                          // "C"
	fmt.Println(caesarCipher("HELLO WORLD", 3))               // "KHOOR ZRUOG"
	fmt.Println(shiftByLetter("B", "K"))                      // "L"
	fmt.Println(vigenereCipher("A C", "KEY"))                 // "K A"
	fmt.Println(scytaleCipher("INFORMATION_AGE", 3))          // "IMNNA_FTAOIGROE"
	fmt.Println(scytaleDecipher("IMNNA_FTAOIGROE", 3))        // "INFORMATION_AGE"
	fmt.Println(scytaleCipher("ALGORITHMS_ARE_IMPORTANT", 8)) // "AOTSRIOALRH_EMRNGIMA_PTT"
}

// shiftLetter
func shiftLetter(letter string, shift int) string {
	if letter == " " {
		return " "
	}
	return string('A' + (letter[0]-'A'+byte(shift))%26)
}

// caesarCipher
func caesarCipher(message string, shift int) string {
	var result strings.Builder
	for _, char := range message {
		if char == ' ' {
			result.WriteRune(' ')
		} else {
			shifted := 'A' + (char-'A'+rune(shift))%26
			result.WriteRune(shifted)
		}
	}
	return result.String()
}

// shiftByLetter
func shiftByLetter(letter string, letterShift string) string {
	if letter == " " {
		return " "
	}
	shift := letterShift[0] - 'A'
	return shiftLetter(letter, int(shift))
}

// vigenereCipher
func vigenereCipher(message string, key string) string {
	var result strings.Builder
	keyIndex := 0
	keyLen := len(key)
	for _, char := range message {
		if char == ' ' {
			result.WriteRune(' ')
		} else {
			shift := key[keyIndex%keyLen] - 'A'
			shifted := 'A' + (char-'A'+rune(shift))%26
			result.WriteRune(shifted)
			keyIndex++
		}
	}
	return result.String()
}

// scytaleCipher
func scytaleCipher(message string, shift int) string {
	length := len(message)
	if length%shift != 0 {
		message += strings.Repeat("_", shift-(length%shift))
	}
	length = len(message)
	columns := length / shift
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		newIndex := (i / columns) + (shift * (i % columns))
		result[newIndex] = message[i]
	}
	return string(result)
}

// scytaleDecipher
func scytaleDecipher(message string, shift int) string {
	length := len(message)
	columns := length / shift
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		originalIndex := (i / shift) + (columns * (i % shift))
		result[originalIndex] = message[i]
	}
	return string(result)
}
