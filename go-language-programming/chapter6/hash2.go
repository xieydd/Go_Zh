package main

import (
    "io"
    "crypto/sha1"
    "crypto/md5"
    "fmt"
    "os"
)

// We van found , when i change 1.txt file content, md5 will change, but sha1 not change
func main() {
    TestFile := "1.txt"
    infile, inerr := os.Open(TestFile)
    defer infile.Close()
    if inerr == nil {
        md5h := md5.New()
        io.Copy(md5h, infile)
        fmt.Printf("%x %s\n",md5h.Sum([]byte("")), TestFile)

        sha1h := sha1.New()
        io.Copy(sha1h, infile)
        fmt.Printf("%x %s\n", sha1h.Sum([]byte("")), TestFile)
    }else {
        fmt.Println(inerr)
        os.Exit(1)
    }
}

