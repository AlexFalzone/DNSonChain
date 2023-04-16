package main

import (
	"fmt"
	"net"
	"runtime"
	"sync"
	"test/dnsHelper"
	"test/inject"
	"test/request"
	"test/util"

	"github.com/ethereum/go-ethereum/log"
	"github.com/miekg/dns"
)

/*
TO-DO:
-sistemare la funzione HandleRequest in maniera tale che non sia necessario stoppare il servizio DNS
*/

func main() {

	var url string
	var cert string
	var ip string
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
				hostname, ex, msg, addr := dnsHelper.HandleRequest(conn)

				if ex == dnsHelper.TypeASomet {

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
						// In this case the certificate is not generated, but it is requested from the blockchain
						// and the expiration date is not specified
						ip, cert, errRequest = request.Request(client, 4, hostname, "", 0)
						if errRequest != nil {
							fmt.Println(err)
						}
						fmt.Println("IP address:", ip)
					}()

					// Wait for the request to be completed
					wg.Wait()

					if errRequest == nil {
						// Create a DNS response with the IP address obtained from the blockchain
						response := new(dns.Msg)
						response.SetReply(msg) // Use the original DNS request message `msg` from the HandleRequest function
						response.Authoritative = true
						response.Answer = append(response.Answer, &dns.A{
							Hdr: dns.RR_Header{Name: msg.Question[0].Name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 0},
							A:   net.ParseIP(ip),
						})

						fmt.Println("IP address PARSED:", ip)

						// Send the DNS response with the IP address to the client
						responseBytes, _ := response.Pack()
						_, err = conn.WriteTo(responseBytes, addr)
						if err != nil {
							log.Error("Error writing DNS response", "err", err)
						}
					}

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
