# RPCBenchmarks
simple benchmark for some rpc libraries

Results:
```
goos: linux
goarch: amd64
pkg: rpcBench
BenchmarkStandard-8     	   34070	     31116 ns/op	    1016 B/op	      28 allocs/op
BenchmarkGoRPC-8        	   23294	     51287 ns/op	     727 B/op	      20 allocs/op
BenchmarkGoRPCBatch-8   	   63562	     18179 ns/op	    1174 B/op	      26 allocs/op
BenchmarkRPCX-8         	   29703	     37783 ns/op	    1755 B/op	      38 allocs/op
BenchmarkNanoMsg-8      	   52345	     21934 ns/op	     536 B/op	      23 allocs/op
PASS
ok  	rpcBench	8.361s

```
