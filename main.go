package main

import (
	"fmt"
	"net"
	"runtime"
	"sync"
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

	var url string
	var cert string
	var errRequest error

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

			// Create a new connection UDP to port 53
			conn, err := net.ListenPacket("udp", ":5354")
			if err != nil {
				fmt.Println("Error listening on DNS port:", err)
				return
			}
			defer conn.Close()

			for {
				hostname, ex := dns.HandleRequest(conn)

				if ex == dns.TypeASomet {

					if choiceInt == 1 { //Infura
						url = "https://sepolia.infura.io/v3/d777809793694d9dacf5e1f94bfec65a"
					} else if choiceInt == 2 { //Localhost
						url = "http://localhost:8545"
					}

					client, err := request.DialClient(url)
					if err != nil {
						fmt.Println(err)
					}

					var wg sync.WaitGroup
					wg.Add(1)
					// Get certificate
					go func() {
						defer wg.Done()
						_, cert, errRequest = request.Request(client, 4, hostname, "")
						if errRequest != nil {
							fmt.Println(err)
						}
					}()

					wg.Wait()

					if errRequest == nil {
						// Remove the first character of the certificate (it's a "0")
						// because the blockchain returns the certificate with a "0" at the beginning
						cert = cert[1:]
						fmt.Println("Richiesta ricevuta:", cert)
					}

					//INJECT CERTIFICATE
					switch runtime.GOOS {
					case "windows":
						fmt.Println("Windows")
						certPathWindows, err := inject.SaveCertToFileWindows(cert)
						if err != nil {
							fmt.Println(err)
						}

						err = inject.InjectWindows(hostname, certPathWindows)
						if err != nil {
							fmt.Println(err)
						}
					case "linux":
						fmt.Println("Linux")

						certpathLinux, err := inject.SaveCertToFile(cert)
						if err != nil {
							fmt.Println(err)
						}

						err = inject.InjcetLinux(hostname, certpathLinux)
						if err != nil {
							fmt.Println(err)
						}
					}
				}
			}
		}
	}
}
