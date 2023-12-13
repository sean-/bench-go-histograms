package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/openhistogram/circonusllhist"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	minTime  = 0 * time.Nanosecond
	maxTime  = 10 * time.Second
	randSeed = 42
)

func randDuration(f *faker.Faker, min, max time.Duration) time.Duration {
	i := f.IntBetween(int(min), int(max))
	return time.Duration(i) * time.Nanosecond
}

func Benchmark_CircLLHistDuration(b *testing.B) {
	h := circonusllhist.New()
	f := faker.NewWithSeed(rand.NewSource(randSeed))

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			h.RecordDuration(randDuration(&f, minTime, maxTime))
		}
	})
}

// func Benchmark_CircLLHistDurationNoLookup(b *testing.B) {
// 	h := circonusllhist.NewNoLocks()
// 	f := faker.NewWithSeed(rand.NewSource(randSeed))

// 	b.RunParallel(func(pb *testing.PB) {
// 		for pb.Next() {
// 			d := randDuration(&f, minTime, maxTime)
// 			h.RecordDuration(d)
// 		}
// 	})
// }

func Benchmark_PromDuration(b *testing.B) {
	h := prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "bench_prom_dur_hist",
	})
	f := faker.NewWithSeed(rand.NewSource(randSeed))

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			h.Observe(randDuration(&f, minTime, maxTime).Seconds())
		}
	})
}
