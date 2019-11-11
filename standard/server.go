package standard

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"rpcBench/args"
)

// ListenAndServe starts a standard rpc server
func ListenAndServe(port int) {
	arith := new(args.MyServer)
	server := rpc.NewServer()
	server.Register(arith)

	sport := fmt.Sprintf(":%d", port)
	l, e := net.Listen("tcp", sport)
	if e != nil {
		log.Fatal("ListenError:", e)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("AcceptError", err)
		}

		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		defer conn.Close()
	}
}

//CreateNewClient creates new client
func CreateNewClient(ip string, port int) *rpc.Client {
	is := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.Dial("tcp", is)

	if err != nil {
		log.Println("CreateNewClientError:", err)
		return nil
	}
	// defer conn.Close()
	return jsonrpc.NewClient(conn)
}
