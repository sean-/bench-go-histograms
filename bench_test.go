package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/openhistogram/circonusllhist"
	"github.com/prometheus/client_golang/prometheus"
	prometheusgo "github.com/prometheus/client_model/go"
)

const (
	minTime = 1 * time.Nanosecond
	maxTime = 10 * time.Second
)

var randSeed int64 = rand.Int63()

func randDuration(f *faker.Faker, min, max time.Duration) time.Duration {
	i := f.IntBetween(int(min), int(max))
	return time.Duration(i)
}

func Benchmark_CircLLHistDuration(b *testing.B) {
	h1 := circonusllhist.New()
	f := faker.NewWithSeed(rand.NewSource(randSeed))

	b.Run("CircLLHistDuration", func(pb *testing.B) {
		h1.RecordDuration(randDuration(&f, minTime, maxTime))
	})
}

func Benchmark_CircLLHistApproxSum(b *testing.B) {
	h1 := circonusllhist.New()
	src := rand.NewSource(randSeed)
	rnd := rand.New(src)
	f := faker.NewWithSeed(rnd)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			h1.RecordDuration(randDuration(&f, minTime, maxTime))
		}
	})

	b.Run("ApproxSum", func(b *testing.B) {
		_ = h1.ApproxSum()
	})
}

func Benchmark_CircLLHistBinCount(b *testing.B) {
	h1 := circonusllhist.New()
	src := rand.NewSource(randSeed)
	rnd := rand.New(src)
	f := faker.NewWithSeed(rnd)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			h1.RecordDuration(randDuration(&f, minTime, maxTime))
		}
	})

	b.Run("BinCount", func(b *testing.B) {
		_ = h1.BinCount()
	})
}

func Benchmark_CircLLHistBuckets(b *testing.B) {
	h1 := circonusllhist.New()
	src := rand.NewSource(randSeed)
	rnd := rand.New(src)
	f := faker.NewWithSeed(rnd)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			h1.RecordDuration(randDuration(&f, minTime, maxTime))
		}
	})

	b.Run("Count", func(b *testing.B) {
		_ = h1.Count()
	})
}

func Benchmark_CircLLHistMerge(b *testing.B) {
	h1 := circonusllhist.New()
	h2 := circonusllhist.New()
	f := faker.NewWithSeed(rand.NewSource(randSeed))

	b.Run("Merge", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				h2.RecordDuration(randDuration(&f, minTime, maxTime))
			}
		})

		h1.Merge(h2)
		h2.Reset()
	})
}

func Benchmark_CircLLHistQuantile(b *testing.B) {
	h1 := circonusllhist.New()
	src := rand.NewSource(randSeed)
	rnd := rand.New(src)
	f := faker.NewWithSeed(rnd)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			h1.RecordDuration(randDuration(&f, minTime, maxTime))
		}
	})

	b.Run("ValueAtQuantile", func(b *testing.B) {
		h1.ValueAtQuantile(rnd.Float64())
	})
}

func Benchmark_CircLLHistDurationNoLookup(b *testing.B) {
	h1 := circonusllhist.NewNoLocks()
	f := faker.NewWithSeed(rand.NewSource(randSeed))

	b.Run("CircLLHistDurationNoLookup", func(pb *testing.B) {
		d := randDuration(&f, minTime, maxTime)
		h1.RecordDuration(d)
	})
}

func Benchmark_PromDuration(b *testing.B) {
	h1 := prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "bench_prom_dur_hist",
	})
	f := faker.NewWithSeed(rand.NewSource(randSeed))

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			h1.Observe(randDuration(&f, minTime, maxTime).Seconds())
		}
	})
}

func Benchmark_PromApproxSum(b *testing.B) {
	src := rand.NewSource(randSeed)
	rnd := rand.New(src)
	f := faker.NewWithSeed(rnd)
	h1 := prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "bench_prom_dur_hist",
	})

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			h1.Observe(randDuration(&f, minTime, maxTime).Seconds())
		}
	})

	b.Run("SampleSum", func(b *testing.B) {
		m := &prometheusgo.Metric{}
		if err := h1.Write(m); err != nil {
			_ = err // b.Skip()
		}
		if m.Histogram.SampleSum != nil {
			_ = m.Histogram.SampleSum
		}
	})
}

func Benchmark_PromApproxCount(b *testing.B) {
	src := rand.NewSource(randSeed)
	rnd := rand.New(src)
	f := faker.NewWithSeed(rnd)
	h1 := prometheus.NewHistogram(prometheus.HistogramOpts{
		Name: "bench_prom_dur_hist",
	})

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			h1.Observe(randDuration(&f, minTime, maxTime).Seconds())
		}
	})

	b.Run("SampleCount", func(b *testing.B) {
		m := &prometheusgo.Metric{}
		if err := h1.Write(m); err != nil {
			_ = err // b.Skip()
		}
		if m.Histogram.SampleCount != nil {
			_ = m.Histogram.SampleCount
		}
	})
}
