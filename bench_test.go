package main

import (
	"context"
	"testing"

	"rpcBench/args"
	gorpc2 "rpcBench/gorpc"
	"rpcBench/mango"
	"rpcBench/rpcx"
	"rpcBench/standard"

	"github.com/smallnest/rpcx/log"
	"github.com/valyala/gorpc"
)

const (
	sPort = 8569
	gPort = 8570
	xPort = 8571
	mPort = 8572
)

func init() {
	log.SetDummyLogger()
	gorpc2.RegisterTypes()
	go standard.ListenAndServe(sPort)
	go gorpc2.ListenAndServe(gPort)
	go rpcx.ListenAndServe(xPort)
	go mango.ListenAndServe(mPort)
}

func BenchmarkStandard(b *testing.B) {
	c := standard.CreateNewClient("localhost", sPort)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var reply int
		if err := c.Call(args.MyServerSum, args.Args{A: 10, B: 1}, &reply); err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
	c.Close()
}

func BenchmarkGoRPC(b *testing.B) {
	c, dc := gorpc2.NewClient("localhost", gPort)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := dc.Call("Sum", &args.Args{A: 10, B: 1})
		if err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
	c.Stop()
}

func BenchmarkGoRPCBatch(b *testing.B) {
	c, dc := gorpc2.NewClient("localhost", gPort)
	bc := dc.NewBatch()
	b.ResetTimer()
	var results []*gorpc.BatchResult
	for i := 0; i < b.N; i++ {
		results = append(results, bc.Add("Sum", &args.Args{A: 10, B: 1}))
		if i%10 == 0 {
			bc.Call()
			for _, r := range results {
				<-r.Done
				if r.Error != nil {
					b.Error(r.Error)
				}
			}
			results = nil
		}
	}
	if b.N%10 != 0 {
		bc.Call()
	}
	b.StopTimer()
	c.Stop()
}

func BenchmarkRPCX(b *testing.B) {
	c := rpcx.NewClient("localhost", xPort)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var reply int
		if err := c.Call(context.Background(), "Sum", args.Args{A: 10, B: 1}, &reply); err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
	c.Close()
}

func BenchmarkNanoMsg(b *testing.B) {
	c := mango.NewClient("localhost", mPort)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := c.Send([]byte("DATE")); err != nil {
			b.Errorf("can't send message on push socket: %s", err)
		}
		if _, err := c.Recv(); err != nil {
			b.Errorf("can't receive date: %s", err)
		}
	}
	b.StopTimer()
	c.Close()
}
