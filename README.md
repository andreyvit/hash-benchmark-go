# Hash & Checksum Functions Benchmark (Golang)

Run:

```bash
go test -bench=.
```


## Results (in MB/sec)

| CPU                   | CRC32  | CRC64  | FNV1a-32 | FNV1a-64 |
| --------------------- |------: | -----: | -------: | -------: |
| M2 (MB Air)           |  4 680 |    960 |    467   |    468   |
| Xeon W-2295 3.00 GHz  | 21 560 |  2 070 |    977   |    983   |
| Xeon W-2145 3.70 GHz  | 23 290 |  2 240 |  1 050   |  1 050   |

## Raw Results

```
goos: darwin
goarch: arm64
pkg: github.com/andreyvit/hash-benchmark-go
cpu: Apple M2
BenchmarkHash_CRC32-8            	    5619	    213694 ns/op
BenchmarkHash_CRC64-8            	    1146	   1040201 ns/op
BenchmarkHash_FNV1a32-8          	     560	   2142679 ns/op
BenchmarkHash_FNV1a64_stdlib-8   	     561	   2135785 ns/op
BenchmarkHash_FNV1a64_inline-8   	     560	   2138464 ns/op

goos: linux
goarch: amd64
pkg: github.com/andreyvit/hash-benchmark-go
cpu: Intel(R) Xeon(R) W-2295 CPU @ 3.00GHz
BenchmarkHash_CRC32-36                     24901             46307 ns/op
BenchmarkHash_CRC64-36                      2284            482637 ns/op
BenchmarkHash_FNV1a32-36                    1216           1023380 ns/op
BenchmarkHash_FNV1a64_stdlib-36             1164           1017803 ns/op
BenchmarkHash_FNV1a64_inline-36             1196           1020778 ns/op

goos: linux
goarch: amd64
pkg: github.com/andreyvit/hash-benchmark-go
cpu: Intel(R) Xeon(R) W-2145 CPU @ 3.70GHz
BenchmarkHash_CRC32-16                     27374             42941 ns/op
BenchmarkHash_CRC64-16                      2720            445973 ns/op
BenchmarkHash_FNV1a32-16                    1112            956008 ns/op
BenchmarkHash_FNV1a64_stdlib-16             1168            956063 ns/op
BenchmarkHash_FNV1a64_inline-16             1098            956932 ns/op
```
