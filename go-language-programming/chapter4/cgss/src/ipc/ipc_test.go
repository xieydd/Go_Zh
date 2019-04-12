package ipc

import (
	"testing"
)

type EchoServer struct{}

func (server *EchoServer) Name() string {
	return "ECHOServer"
}

func (server *EchoServer) Handle(request string) string {
	return "ECHO:" + request
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)
	resq1 := client1.Call("From Client1")
	resq2 := client2.Call("From Client2")

	if resq1 != "ECHO:From Client1" || resq2 != "ECHO:From Client2" {
		t.Error("IpcClient.Call failed. resp1: ", resq1, "resq2:", resq2)
	}

	client1.Close()
	client2.Close()
}
