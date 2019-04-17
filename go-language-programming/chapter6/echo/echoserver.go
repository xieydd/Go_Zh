package main

import (
    "crypto/tls"
    "time"
    "crypto/rand"
    "log"
    "net"
    "io"
)

func handleClient(conn net.Conn) {
    defer conn.Close()

    buf := make([]byte, 512)
    for {
        log.Print("server: conn: wating")
        n, err := conn.Read(buf)
        if err != nil {
            if err != io.EOF {
                log.Printf("server: conn: read: %s", err)
            }
            break
        }
        log.Printf("server: conn: echo %q\n", string(buf[:n]))
        n, err = conn.Write(buf[:n])

        log.Printf("server: conn: wrote %d bytes", n)
        if err != nil {
            log.Printf("server: write: %s", err)
            break
        }
    }
    log.Printf("server: conn closed")
}

func main() {
    cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
    if err != nil {
        log.Fatalf("server: loadkeys: %s", err)
    }
    config := tls.Config{Certificates:[]tls.Certificate{cert}}
    config.Time = time.Now
    config.Rand = rand.Reader

    service := "127.0.0.1:8000"

    listen, err := tls.Listen("tcp", service, &config)
    if err != nil {
        log.Fatalf("server: listen: %s", err)
    }
    log.Printf("server: listening")
    for {
        conn, err := listen.Accept()
        if err != nil {
            log.Printf("server: accept: %s", err)
            break
        }
    log.Printf("server: accpted from %s", conn.RemoteAddr())
    go handleClient(conn)
    }
}
