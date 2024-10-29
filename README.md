# Benchmark for Go histogram implementations

* [`openhistogram/circonusllhist`](https://github.com/openhistogram/circonusllhist)
* [`prometheus/client_golang`](github.com/prometheus/client_golang)

```
$ go test -test.bench=. -benchmem
goos: darwin
goarch: arm64
pkg: github.com/sean-/bench-go-histograms
Benchmark_CircLLHistDuration-10     	                 7221150	             149.7 ns/op	       0 B/op	       0 allocs/op
Benchmark_CircLLHistApproxSum/ApproxSum-10         	1000000000	         0.0000002 ns/op	       0 B/op	       0 allocs/op
Benchmark_CircLLHistBinCount/BinCount-10           	1000000000	         0.0000001 ns/op	       0 B/op	       0 allocs/op
Benchmark_CircLLHistBuckets/Count-10               	1000000000	         0.0000000 ns/op	       0 B/op	       0 allocs/op
Benchmark_CircLLHistMerge/Merge-10                 	   8280682	             142.6 ns/op	       0 B/op	       0 allocs/op
Benchmark_CircLLHistQuantile/ValueAtQuantile-10    	1000000000	         0.0000005 ns/op	       0 B/op	       0 allocs/op
Benchmark_PromDuration-10                          	   2781200	             422.3 ns/op	       0 B/op	       0 allocs/op
Benchmark_PromApproxSum/SampleSum-10               	1000000000	         0.0000015 ns/op	       0 B/op	       0 allocs/op
Benchmark_PromApproxCount/SampleCount-10           	1000000000	         0.0000059 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/sean-/bench-go-histograms	4.399s

$ go version
go version go1.21.5 darwin/arm64
```
