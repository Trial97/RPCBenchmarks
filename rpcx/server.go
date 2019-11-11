package rpcx

import (
	"fmt"
	"log"

	"rpcBench/args"

	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/server"
)

// ListenAndServe starts a rpcx server
func ListenAndServe(port int) {
	s := server.NewServer()
	s.Register(new(args.MyServer3), "")
	// s.RegisterName("MyServer", new(args.MyServer), "")
	err := s.Serve("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
}

// NewClient returns a rpcx client
func NewClient(ip string, port int) client.XClient {
	d := client.NewPeer2PeerDiscovery(fmt.Sprintf("tcp@%s:%d", ip, port), "")
	opt := client.DefaultOption
	opt.SerializeType = protocol.JSON

	return client.NewXClient("MyServer3", client.Failtry, client.RandomSelect, d, opt)
	// defer xclient.Close()
}
