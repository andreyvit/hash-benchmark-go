package benchmark_test

import (
	"crypto/rand"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
	"io"
	"testing"
)

var ( // ensure results are not optimized away
	total32 uint32
	total64 uint64
)

func BenchmarkHash_CRC32(t *testing.B) {
	buf := make([]byte, 1024*1024)
	must(io.ReadFull(rand.Reader, buf))
	c := crc32.NewIEEE()

	t.ResetTimer()
	total32 = 0
	for range t.N {
		c.Reset()
		c.Write(buf)
		sum := c.Sum32()
		total32 += sum
	}
}

func BenchmarkHash_CRC64(t *testing.B) {
	buf := make([]byte, 1024*1024)
	must(io.ReadFull(rand.Reader, buf))
	c := crc64.New(crc64.MakeTable(crc64.ISO))

	t.ResetTimer()
	total64 = 0
	for range t.N {
		c.Reset()
		c.Write(buf)
		sum := c.Sum64()
		total64 += sum
	}

	// t.Logf("N=%d total=%x", t.N, total)
}

func BenchmarkHash_FNV1a32(t *testing.B) {
	buf := make([]byte, 1024*1024)
	must(io.ReadFull(rand.Reader, buf))
	c := fnv.New32a()

	t.ResetTimer()
	total32 = 0
	for range t.N {
		c.Reset()
		c.Write(buf)
		sum := c.Sum32()
		total32 += sum
	}
}

func BenchmarkHash_FNV1a64_stdlib(t *testing.B) {
	buf := make([]byte, 1024*1024)
	must(io.ReadFull(rand.Reader, buf))
	c := fnv.New64a()

	t.ResetTimer()
	total64 = 0
	for range t.N {
		c.Reset()
		c.Write(buf)
		sum := c.Sum64()
		total64 += sum
	}
}

func BenchmarkHash_FNV1a64_inline(t *testing.B) {
	buf := make([]byte, 1024*1024)
	must(io.ReadFull(rand.Reader, buf))

	t.ResetTimer()
	total64 = 0
	for range t.N {
		c := uint64(14695981039346656037)
		for _, b := range buf {
			c = (c ^ uint64(b)) * 1099511628211
		}
		total64 += uint64(c)
	}
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
