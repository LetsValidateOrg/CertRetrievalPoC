package main

import (
	"crypto/tls"
	"fmt"
	"log"
)

func main() {
	conf := tls.Config{}
	conn, err := tls.Dial("tcp", "www.publicntp.org:443", &conf)
	if err != nil {
		log.Fatalln("Error in TLS dial: ", err)
		return
	}
	defer conn.Close()
	certs := conn.ConnectionState().PeerCertificates
	for _, currCert := range certs {
		fmt.Println("\nFound TLS cert:")
		fmt.Println("\tIssuer Name: ", currCert.Issuer)
	}
}
