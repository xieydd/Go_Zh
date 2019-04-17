package main

import (
    "fmt"
    "net/http"
    "crypto/tls"
    "crypto/rand"
    "crypto/pem"
    "crypto/x509"
    "crypto/rsa"
    "time"
    "net"
    "io/ioutil"
    "errors"
)

const (
    SERVER_DOMAIN = "localhost"
    SERVER_PORT = 8080
    RESPONSE = "hello"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    w.Header().Set("Content-Length", fmt.Sprint(len(RESPONSE)))
    w.Write([]byte(RESPONSE))
}

func YourListenAndServeTLS(addr string, certFile string, keyFile string, hendler http.Handler) error {
    config := &tls.Config{
        Rand: rand.Reader,
        Time: time.Now(),
        NextProto: []string{"http/1.1"}
    }
    config.Certificates = make([]tls.Certificate,1)
    config.Certificates[0], err = YourLoadX509KeyPair(certFile, keyFile)
    if err != nil {
        return err
    }

    conn, err := net.Listen("tcp", addr)
    if err != nil {
        return err
    }
    tlsListener := tls.NewListener(conn, config)
    return http.Serve(tlsListener, handler)
}

func YourLoadX509Pair(certFile, keyFile string) (cert tls.Certificate, err error) {
    certPEMBlock, err := ioutil.ReadFile(certFile)
    if err != nil {
        return
    }
    certDERBlock,restPEMBlock := pem.Decode(certPemBlock)
    if certDERBlock == nil {
        err = errors.New("crypto/tls: failed to parse certificate PEM data")
        return
    }
    certDERBlockChain, _ := pem.Decode(restPEMBlock)
    if certDERBlockChain == nil {
        cert.Certificate = [][]byte{certDERBlock.Bytes}
    }else {
        cert.Certificate = [][]byte{certDERBlock.Bytes, certDERBlockChain.Bytes}
    }

    keyPEMBlock, err := ioutil.ReadFile(keyFile)
    if err != nil {
        return
    }
    keyDERBlock, _ := pem.Decode(keyPEMBlock)
    if keyDERBlock == nil {
        err = errors.New("crypto/tls: failed to parse key PEM data")
        return
    }
    key, err := x509.ParsePKCS1PrivateKey(keyDERBlock.Bytes)
    if err != nil {
        err = errors.New("crypto/tls: failed to parse key")
        return
    }

    cert.PrivateKey = key
    x509Cert, err := x509.ParseCertificate(certDERBlock.Bytes)
    if err != nil {
        return
    }
    if x509Cert.PublicKeyAlgorithm != x509.RSA || x509Cert.PublicKey.(*rsa.PublicKey).N.Cmp(key.PublicKey.N) != 0 {
        err = errors.New("crypto/tls: private key does not match public key")
        return
    }
    return
}

func main() {
    http.HandleFunc(fmt.Sprintf("%s:%d/", SERVER_DOMAIN, SERVER_PORT), rootHandler)
    YourListenAndServeTLS(fmt.Sprintf(":%d", SERVER_PORT), "rui.crt", "rui.key", nil)
}
