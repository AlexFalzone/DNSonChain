package dns

import (
	"fmt"
	"net"
	"strings"

	"github.com/ethereum/go-ethereum/log"
	"github.com/miekg/dns"
)

const (
	dnsServer = "1.1.1.1:53"

	TypeASomet    = 1
	TypeANonSomet = 2
	TypeNonA      = 3
)

// Forward the request to the DNS server and send the response back to the browser
func forwardRequest(conn net.PacketConn, msg *dns.Msg, addr net.Addr) {
	fwd, err := dns.Exchange(msg, dnsServer)
	if err != nil {
		log.Error("Error forwarding DNS request", "err", err)
		return
	}

	// Send the response back to the browser
	packageBytes, _ := fwd.Pack()
	_, err = conn.WriteTo(packageBytes, addr)
	if err != nil {
		log.Error("Error writing DNS response", "err", err)
		return
	}
}

// Extract the hostname from the DNS request
func extractHostnameFromDNS(msg *dns.Msg) string {
	hostname := msg.Question[0].Name
	return strings.TrimSuffix(hostname, ".")
}

// Check if the DNS request is of type A or AAAA
func isTypeAQuery(msg *dns.Msg) bool {
	return len(msg.Question) > 0 && (msg.Question[0].Qtype == dns.TypeA || msg.Question[0].Qtype == dns.TypeAAAA)
}

func HandleRequest(conn net.PacketConn) (hostname string, ex int8) {
	buf := make([]byte, 1024)

	// Read the request from the browser
	n, addr, err := conn.ReadFrom(buf)
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}

	// Create a new DNS message from the request
	msg := &dns.Msg{}
	if err := msg.Unpack(buf[:n]); err != nil {
		fmt.Println("Error unpackaging DNS message:", err)
		return hostname, TypeNonA
	}

	if !isTypeAQuery(msg) {
		forwardRequest(conn, msg, addr)
		return hostname, TypeNonA
	}

	hostnameFromDNS := extractHostnameFromDNS(msg)

	// if the request does NOT have a .bit suffix then forward the request to the cloudflare DNS server
	if !(strings.HasSuffix(hostnameFromDNS, ".somet")) {
		fmt.Println("DNS esterno:", hostnameFromDNS)

		forwardRequest(conn, msg, addr)
		return hostnameFromDNS, TypeANonSomet
	}

	fmt.Println("DNS interno:", hostnameFromDNS)
	return hostnameFromDNS, TypeASomet
}
