package main

import (
	"fmt"
	"net"
	"runtime"
	"test/dns"
	"test/inject"
	"test/request"
	"test/util"
)

/*
TO-DO:
-sistemare la funzione HandleRequest in maniera tale che non sia necessario stoppare il servizio DNS
*/

func main() {
	// Create a new connection UDP to port 53
	conn, err := net.ListenPacket("udp", ":53")
	if err != nil {
		fmt.Println("Error listening on DNS port:", err)
		return
	}
	defer conn.Close()

	var url string

	for {
		/*
			1. Infura
			2. Localhost
			3. Certificates Management
			4. Exit
		*/
		choiceInt := util.DisplayMainMenu()

		if choiceInt == 4 {
			break
		}

		hostname, _ := dns.HandleRequest(conn)

		if choiceInt == 3 { //certificates management
			for {
				/*
					1. Generate Certificate
					2. Revoke Certificate
					3. Renew Certificate
					4. Back to main menu
				*/
				choiceInt := util.DisplayCertificatesMenu()

				if choiceInt == 4 { // Back to main menu
					break
				}

				url := util.GetInfuraOrLocalhost()
				request.CertificatesManagement(choiceInt, url)
			}
		} else if choiceInt == 1 || choiceInt == 2 { //infura or localhost from the main menu
			if choiceInt == 1 { //Infura
				url = "https://goerli.infura.io/v3/d777809793694d9dacf5e1f94bfec65a"
			} else if choiceInt == 2 { //Localhost
				url = "http://localhost:8545"
			}
			client, err := request.DialClient(url)
			if err != nil {
				fmt.Println(err)
			}

			_, _, err = request.Request(client, 4, hostname, nil)
			if err != nil {
				fmt.Println(err)
			}
			//INJECT CERTIFICATE
			switch runtime.GOOS {
			case "windows":
				fmt.Println("Windows")
			case "linux":
				fmt.Println("Linux")
				inject.InjectPKI(hostname)
				inject.InjectMozilla(hostname)
			}
		}
	}
}
