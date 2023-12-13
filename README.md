# Benchmark for Go histogram implementations

* [`openhistogram/circonusllhist`](https://github.com/openhistogram/circonusllhist)
* [`prometheus/client_golang`](github.com/prometheus/client_golang)

```
$ go test -test.bench=. -benchmem
goos: darwin
goarch: arm64
pkg: github.com/sean-/bench-go-histograms
Benchmark_CircLLHistDuration-10    	 7262088	       166.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_PromDuration-10          	 4536674	       281.2 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/sean-/bench-go-histograms	3.382s
```
