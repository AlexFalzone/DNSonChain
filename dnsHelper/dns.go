package dnsHelper

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

// forwardRequest forwards the DNS request to a specified DNS server and sends the response back to the client.
//
// Parameters:
//   - conn: net.PacketConn representing the connection to send and receive data.
//   - msg: *dns.Msg representing the DNS message to be forwarded.
//   - addr: net.Addr representing the address of the client.
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

// extractHostnameFromDNS extracts the hostname from a DNS message.
//
// Parameters:
//   - msg: *dns.Msg representing the DNS message.
//
// Returns:
//   - string: The extracted hostname.
func extractHostnameFromDNS(msg *dns.Msg) string {
	hostname := msg.Question[0].Name
	return strings.TrimSuffix(hostname, ".")
}

// isTypeAQuery checks if the DNS request is of type A or AAAA.
//
// Parameters:
//   - msg: *dns.Msg representing the DNS message.
//
// Returns:
//   - bool: true if the DNS request is of type A or AAAA, false otherwise.
func isTypeAQuery(msg *dns.Msg) bool {
	return len(msg.Question) > 0 && (msg.Question[0].Qtype == dns.TypeA || msg.Question[0].Qtype == dns.TypeAAAA)
}

// HandleRequest processes a DNS request and determines the hostname and request type.
// If the request is not a type A or AAAA query, it forwards the request to the specified DNS server.
//
// Parameters:
//   - conn: net.PacketConn representing the connection to send and receive data.
//
// Returns:
//   - hostname: string representing the extracted hostname.
//   - ex: int8 representing the request type (TypeASomet, TypeANonSomet, or TypeNonA).
func HandleRequest(conn net.PacketConn) (hostname string, ex int8, msg *dns.Msg, addr net.Addr) {
	buf := make([]byte, 1024)

	// Read the request from the browser
	n, addr, err := conn.ReadFrom(buf)
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}

	// Create a new DNS message from the request
	msg = &dns.Msg{}
	if err := msg.Unpack(buf[:n]); err != nil {
		fmt.Println("Error unpackaging DNS message:", err)
		return hostname, TypeNonA, msg, addr
	}

	if !isTypeAQuery(msg) {
		forwardRequest(conn, msg, addr)
		return hostname, TypeNonA, msg, addr
	}

	hostnameFromDNS := extractHostnameFromDNS(msg)

	// if the request does NOT have a .bit suffix then forward the request to the cloudflare DNS server
	if !(strings.HasSuffix(hostnameFromDNS, ".chain")) {
		fmt.Println("DNS esterno:", hostnameFromDNS)

		forwardRequest(conn, msg, addr)
		return hostnameFromDNS, TypeASomet, msg, addr

	}

	fmt.Println("DNS interno:", hostnameFromDNS)
	return hostnameFromDNS, TypeASomet, msg, addr
}
