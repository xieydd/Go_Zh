package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"log"
)

func main() {
	pair, e := tls.LoadX509KeyPair("./client.crt", "./client.key")
	if e != nil {
		log.Fatal("LoadX509KeyPair Error:", e)
	}

	p := x509.NewCertPool()
	ca, e := ioutil.ReadFile("./ca.crt")
	if e != nil {
		log.Fatal("read ca: %s", e)
	}
	p.AppendCertsFromPEM(ca)
	config := tls.Config{RootCAs: p, Certificates: []tls.Certificate{pair}, InsecureSkipVerify: true}

	conn, err := tls.Dial("tcp", "127.0.0.1:8000", &config)
	if err != nil {
		log.Fatalf("client: Dial: %s", err)
	}
	log.Printf("client connect to ", conn.RemoteAddr())

	state := conn.ConnectionState()
	log.Println("client: handshake: ", state.HandshakeComplete)
	log.Println("client: mutual: ", state.NegotiatedProtocolIsMutual)

	message := "Hello, xieydd!\n"
	n, err := io.WriteString(conn, message)
	if err != nil {
		log.Fatalf("client: write: %s", err)
	}
	log.Printf("client: wrote %q (%d bytes)", message, n)

	reply := make([]byte, 256)
	n, err = conn.Read(reply)
	log.Printf("client: read %q (%d bytes)", string(reply[:n]), n)
	log.Print("client: exiting")
}
