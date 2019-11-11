package gorpc

import (
	"fmt"
	"log"

	"rpcBench/args"

	"github.com/valyala/gorpc"
)

func newDispatcher() (d *gorpc.Dispatcher) {
	d = gorpc.NewDispatcher()
	// Register exported service functions
	d.AddService("MyServer", &args.MyServer2{})
	return
}

// ListenAndServe starts a gorpc server
func ListenAndServe(port int) {
	d := newDispatcher()
	s := gorpc.NewTCPServer(fmt.Sprintf(":%d", port), d.NewHandlerFunc())
	defer s.Stop()
	if err := s.Serve(); err != nil {
		log.Fatalf("Cannot start rpc server: %s", err)
	}
}

// NewClient returns a gorpc Service client
func NewClient(ip string, port int) (*gorpc.Client, *gorpc.DispatcherClient) {
	d := newDispatcher()
	// Start the client and connect it to the server
	is := fmt.Sprintf("%s:%d", ip, port)
	c := gorpc.NewTCPClient(is)
	c.Start()

	// Create a client wrapper for calling server functions.
	dc := d.NewServiceClient("MyServer", c)
	return c, dc
}

func RegisterTypes() {
	gorpc.RegisterType(&args.MyServer{})
	gorpc.RegisterType(&args.Args{})
}
