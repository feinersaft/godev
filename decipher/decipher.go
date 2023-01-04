package main

import "fmt"

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	message := ""
	keyIndex := 0

	for j := 0; j < len(cipherText); j++ {
		// the single quotes convert the A character to a `rune` type rather than a string, which prevents a compilation error
		message += string(((cipherText[j] - keyword[keyIndex] + 26) % 26) + 'A')
		keyIndex++
		keyIndex %= len(keyword) // prevents address overflow of the keyword index
	}
	fmt.Println(message)
}
