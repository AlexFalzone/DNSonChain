package util

import (
	"fmt"
	"strconv"
)

func DisplayMainMenu() int {
	fmt.Println("1. Infura")
	fmt.Println("2. Localhost")
	fmt.Println("3. Certificates Management")
	fmt.Println("4. Exit")

	var choice string
	fmt.Scan(&choice)
	choiceInt, _ := strconv.Atoi(choice)
	return choiceInt
}

func DisplayCertificatesMenu() int {
	fmt.Println("1. Generate Certificate")
	fmt.Println("2. Revoke Certificate")
	fmt.Println("3. Renew Certificate")
	fmt.Println("4. Back to main menu")

	var choice string
	fmt.Scan(&choice)
	choiceInt, _ := strconv.Atoi(choice)
	return choiceInt
}

func GetInfuraOrLocalhost() string {
	fmt.Println("Do you want to use Infura or your localhost?")
	fmt.Println("1. Infura")
	fmt.Println("2. Localhost")

	var choiceURL string
	fmt.Scan(&choiceURL)
	choiceIntURL, _ := strconv.Atoi(choiceURL)

	if choiceIntURL == 1 { //Infura
		return "https://goerli.infura.io/v3/d777809793694d9dacf5e1f94bfec65a"
	} else if choiceIntURL == 2 { //Localhost
		return "http://localhost:8545"
	}
	return ""
}
