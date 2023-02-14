package util

import (
	"fmt"
	"strconv"
)

func MenuCertificate() (choiceKey int, choiceBitSize int, pathToCert string) {
	fmt.Println("Name:")
	var name string
	fmt.Scan(&name)

	fmt.Println("Choose the key:")
	fmt.Println("1. RSA")
	fmt.Println("2. ECDSA")
	fmt.Println("3. ED25519")

	var keyChoice string
	fmt.Scan(&keyChoice)
	keyChoiceInt, _ := strconv.Atoi(keyChoice)

	switch keyChoiceInt {
	case 1:
		fmt.Println("Choose the bit size:")
		fmt.Println("1. 2048")
		fmt.Println("2. 3072")
		fmt.Println("3. 4096 (discouraged)")
	case 2:
		fmt.Println("Choose the curve:")
		fmt.Println("1. P-256")
		fmt.Println("2. P-384")
		fmt.Println("3. P-521")
	case 3:
		fmt.Println("Choose the bit size:")
		fmt.Println("1. 256")
		fmt.Println("2. 384")
		fmt.Println("3. 512")
	default:
		fmt.Println("Invalid choice")
		return
	}
	var bitSizeChoice string
	fmt.Scan(&bitSizeChoice)
	bitSizeChoiceInt, _ := strconv.Atoi(bitSizeChoice)

	fmt.Println("Choose the path to the certificate:")
	var path string
	fmt.Scan(&path)

	return keyChoiceInt, bitSizeChoiceInt, path
}
