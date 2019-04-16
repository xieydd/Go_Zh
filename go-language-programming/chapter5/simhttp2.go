package main

import (
    "fmt"
    "os"
    "net"
    "io/ioutil"
)

func checkError(err error) {
    if err != nil {
        fmt.Println(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(0)
    }
}


func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: " ,os.Args[0], " host")
    }

    service := os.Args[1]
    //conn, err := net.Dial("tcp", service)
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)

    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    checkError(err)

    _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
    checkError(err)

    result, err := ioutil.ReadAll(conn)
    checkError(err)
    fmt.Println(string(result))
    os.Exit(0)
}
