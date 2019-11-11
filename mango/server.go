package mango

import (
	"fmt"
	"log"
	"time"

	mangos "nanomsg.org/go/mangos/v2"
	"nanomsg.org/go/mangos/v2/protocol/rep"
	"nanomsg.org/go/mangos/v2/protocol/req"

	// register transports
	_ "nanomsg.org/go/mangos/v2/transport/all"
)

func date() string {
	return time.Now().Format(time.ANSIC)
}

// ListenAndServe starts a nanomsg server
func ListenAndServe(port int) {
	var sock mangos.Socket
	var err error
	var msg []byte
	if sock, err = rep.NewSocket(); err != nil {
		log.Fatalf("can't get new rep socket: %s", err)
	}
	if err = sock.Listen(fmt.Sprintf("tcp://127.0.0.1:%d", port)); err != nil {
		log.Fatalf("can't listen on rep socket: %s", err)
	}
	for {
		// Could also use sock.RecvMsg to get header
		msg, err = sock.Recv()
		if string(msg) == "DATE" { // no need to terminate
			// fmt.Println("NODE0: RECEIVED DATE REQUEST")
			d := date()
			// fmt.Printf("NODE0: SENDING DATE %s\n", d)
			err = sock.Send([]byte(d))
			if err != nil {
				log.Fatalf("can't send reply: %s", err)
			}
		}
	}
}

// NewClient returns a mangos Socket
func NewClient(ip string, port int) (sock mangos.Socket) {
	var err error
	if sock, err = req.NewSocket(); err != nil {
		log.Fatalf("can't get new req socket: %s", err.Error())
	}
	if err = sock.Dial(fmt.Sprintf("tcp://%s:%d", ip, port)); err != nil {
		log.Fatalf("can't dial on req socket: %s", err.Error())
	}
	return
}

// func node1(url string) {
// 	var sock mangos.Socket
// 	var err error
// 	var msg []byte

// 	if sock, err = req.NewSocket(); err != nil {
// 		log.Fatalf("can't get new req socket: %s", err.Error())
// 	}
// 	if err = sock.Dial(url); err != nil {
// 		log.Fatalf("can't dial on req socket: %s", err.Error())
// 	}
// 	fmt.Printf("NODE1: SENDING DATE REQUEST %s\n", "DATE")
// 	if err = sock.Send([]byte("DATE")); err != nil {
// 		log.Fatalf("can't send message on push socket: %s", err.Error())
// 	}
// 	if msg, err = sock.Recv(); err != nil {
// 		log.Fatalf("can't receive date: %s", err.Error())
// 	}
// 	fmt.Printf("NODE1: RECEIVED DATE %s\n", string(msg))
// 	sock.Close()
// }
