package main

import(
    "net/rpc"
    "net"
    "log"
    "net/http"
    "errors"
    "fmt"
)

type Arith int

type Args struct {
    A, B int
}

type Quotient struct {
    Quo, Rem int
}

func (t *Arith) Multiply(args *Args, reply *int) error {
    *reply = args.A * args.B
    return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
    if args.B == 0 {
        return errors.New("divide by zero")
    }
    quo.Quo = args.A / args.B
    quo.Rem = args.A % args.B
    return nil
}


func main() {
    arith := new(Arith)
    rpc.Register(arith)
    rpc.HandleHTTP()
    l, e := net.Listen("tcp", ":1234")
    if e != nil {
        log.Fatal("listen error:", e)
    }
    go http.Serve(l, nil)

    serverAddress := "127.0.0.1"
    client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
    if err != nil {
        log.Fatal("dialing:", err)
    }

    args := &Args{7,8}
    var reply int
    err = client.Call("Arith.Multiply", args, &reply)
    if err != nil {
        log.Fatal("arith error:", err)
    }
    fmt.Println("Arith: %d*%d=%d", args.A, args.B, reply)

    //quotient := new(Quotient)
    //divCall := client.Go("Arith.Divide", args, &quotient, nil)
    //replyCall := <-divCall.Done
}
