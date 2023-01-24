package dns

import (
	"fmt"
	"net"
	"strings"

	"github.com/miekg/dns"
)

const (
	dnsServer = "1.1.1.1:53"
)

/*
Leggenda variabile ex:
1 = tipo A, sito .somet
2 = tipo A, sito non .somet
3 = tipo non A
*/

func HandleRequest(conn net.PacketConn) (hostname string, ex int8) {
	// Buffer per leggere la richiesta DNS
	buf := make([]byte, 1024)

	// Legge la richiesta dal conn
	n, addr, err := conn.ReadFrom(buf)
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}

	// Crea un messaggio DNS dal buffer
	msg := &dns.Msg{}
	if err := msg.Unpack(buf[:n]); err != nil {
		fmt.Println("Error unpackaging DNS message:", err)
		return hostname, 3
	}

	// Verifica se la richiesta è di tipo A
	if len(msg.Question) == 0 || (msg.Question[0].Qtype != dns.TypeA && msg.Question[0].Qtype != dns.TypeAAAA) {
		// inoltra la richiesta al server DNS cloudflare
		fwd, err := dns.Exchange(msg, dnsServer)
		if err != nil {
			fmt.Println("Error forwarding DNS request:", err)
			return
		}
		// Scrive la risposta sul conn
		packageBytes, _ := fwd.Pack()
		_, err = conn.WriteTo(packageBytes, addr)
		if err != nil {
			fmt.Println("Error writing DNS response:", err)
			return
		}
		return
	}

	// Estrae il nome dell'host dalla richiesta
	hostnameFromDNS := msg.Question[0].Name

	hostnameFromDNS = strings.TrimSuffix(hostnameFromDNS, ".")

	//se la richiesta NON ha un suffisso .bit allora inoltra la richiesta al server DNS cloudflare
	if !(strings.HasSuffix(hostnameFromDNS, ".somet")) {
		fmt.Println("DNS esterno:", hostnameFromDNS)

		// inoltra la richiesta al server DNS cloudflare
		fwd, err := dns.Exchange(msg, dnsServer)
		if err != nil {
			fmt.Println("Error forwarding DNS request:", err)
			return
		}

		// Send the response back to the browser
		packageBytes, _ := fwd.Pack()
		_, err = conn.WriteTo(packageBytes, addr)
		if err != nil {
			fmt.Println("Error writing DNS response:", err)
			return
		}
		return hostnameFromDNS, 2
	}
	fmt.Println("DNS interno:", hostnameFromDNS)

	//hostname con il prefisso .somet
	return hostnameFromDNS, 1
}
