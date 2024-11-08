# Hash & Checksum Functions Benchmark (Golang)

Run:

```bash
go test -count=10 -cpu=1 -bench=.
```

(`-cpu=1` is for undecorated test names to make running benchstat easier)

For computers that tend to throttle quickly (like the MacBook Air), use the `-throttle` flag to introduce a delay between tests:

```bash
go test -throttle -count=10 -cpu=1 -bench=.
```

Then build the new results table:

```bash
go run summarize.go
```


## Results

All FNV1 results are basically the same, so collapsed into a single column.


### MB/sec

| CPU                              |  CRC32 |  CRC64 |   FNV1 | xxHash |
| :------------------------------- | -----: | -----: | -----: | -----: |
| M2 MB Air                        |  4 120 |    870 |    410 |  7 940 |
| Xeon W-2145 3.70GHz              | 23 390 |  2 280 |  1 070 | 16 810 |
| Xeon W-2295 3.00GHz              | 22 630 |  2 190 |  1 030 | 16 280 |

(rounded to 10 MB/sec)


### us/MB

| CPU                              |  CRC32 |  CRC64 |   FNV1 | xxHash |
| :------------------------------- | -----: | -----: | -----: | -----: |
| M2 MB Air                        |    242 |  1 153 |  2 422 |    126 |
| Xeon W-2145 3.70GHz              |     43 |    439 |    938 |     60 |
| Xeon W-2295 3.00GHz              |     44 |    458 |    974 |     61 |
