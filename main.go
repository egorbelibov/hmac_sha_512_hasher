package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
	"syscall"
)

func main() {
	printIntro()
	showInput := shouldShowInput()

	message := getInput("Message: ", showInput)
	key := getInput("Key: ", showInput)

	base16Encoded, base64Encoded := hmacSha512(message, key)
	printResult(base16Encoded, base64Encoded)
}

func printIntro() {
	fmt.Println("===========================")
	fmt.Println("HMAC Hasher to Base 16 & 64")
	fmt.Println("===========================")
}

func shouldShowInput() bool {
	if len(os.Args) == 2 {
		if os.Args[1] == "show-input" {
			return true
		}
	}
	return false
}

func getInput(message string, showInput bool) string {
	var input string
	if showInput {
		fmt.Print(message)
		fmt.Scanln(&input)
	} else {
		fmt.Println("(Your input is invisible)")
		fmt.Print(message)
		input = readSecretInput()
		fmt.Println()
	}
	return input
}

func readSecretInput() string {
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println("\nWrong input. Try again: ")
	}
	password := string(bytePassword)
	return strings.TrimSpace(password)
}

func hmacSha512(message, key string) (string, string) {
	hmacSha512 := hmac.New(sha512.New, []byte(key))
	hmacSha512.Write([]byte(message))

	base16Encoded := hex.EncodeToString(hmacSha512.Sum(nil))
	base64Encoded := base64.StdEncoding.EncodeToString(hmacSha512.Sum(nil))

	return base16Encoded, base64Encoded
}

func printResult(base16Encoded, base64Encoded string) {
	fmt.Println("===========================")
	fmt.Printf("==> Base 16:\n%s\n", base16Encoded)
	fmt.Printf("==> Base 64:\n%s\n", base64Encoded)
}
