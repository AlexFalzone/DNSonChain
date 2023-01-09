package dns

import (
	"fmt"
	"net"

	"github.com/miekg/dns"
)

func HandleRequest(conn net.PacketConn) {
	// Alloca un buffer di dimensione massima per i pacchetti DNS
	buf := make([]byte, 65536)

	// Legge i dati dalla connessione
	n, _, err := conn.ReadFrom(buf)
	if err != nil {
		fmt.Println("Errore durante la lettura dei dati:", err)
		return
	}

	// Esegue il parsing del pacchetto DNS
	msg := &dns.Msg{}
	err = msg.Unpack(buf[:n])
	if err != nil {
		fmt.Println("Errore durante il parsing del pacchetto DNS:", err)
		return
	}

	// Verifica se la richiesta è di tipo A
	if msg.Opcode != dns.OpcodeQuery || msg.Rcode != dns.RcodeSuccess || len(msg.Question) != 1 {
		fmt.Println("Richiesta DNS non valida")
		return
	}
	question := msg.Question[0]
	if question.Qtype != dns.TypeA {
		fmt.Println("Richiesta DNS non di tipo A")
		return
	}

	// Stampa il nome dell'URL
	fmt.Println("Nome dell'URL:", question.Name)
}
