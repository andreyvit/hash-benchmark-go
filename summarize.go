//go:build ignore

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const inputDir = "results"

const modelColLabel = "CPU"

const (
	minModelColWidth  = 32
	minResultColWidth = len("12 345")
)

var renameAlgs = map[string]string{
	"FNV1a64_stdlib": "FNV1",
	"FNV1a32":        "",
	"FNV1a64_inline": "",
	"XXHash":         "xxHash",
}

type modelData struct {
	name string
	algs map[string]int64
}

var benchmarkLineRe = regexp.MustCompile(`^BenchmarkHash_(\w+)\s+\d+\s+(\d+) ns/op$`)

func main() {
	var algs []string
	var models []*modelData

	for _, ent := range must(os.ReadDir(inputDir)) {
		name := ent.Name()
		modelName, ok := strings.CutSuffix(name, ".txt")
		if !ok {
			continue
		}
		modelName = strings.ReplaceAll(modelName, "_", " ")
		fn := filepath.Join(inputDir, name)
		lines := strings.FieldsFunc(string(must(os.ReadFile(fn))), isNL)

		data := &modelData{name: modelName, algs: make(map[string]int64)}
		models = append(models, data)

		var alg string
		var sum, count int64

		flush := func() {
			if alg == "" || count == 0 {
				return
			}
			if !slices.Contains(algs, alg) {
				algs = append(algs, alg)
			}
			data.algs[alg] = (sum + count/2) / count
		}

		for _, line := range lines {
			m := benchmarkLineRe.FindStringSubmatch(line)
			if m == nil {
				continue
			}
			a, ns := m[1], must(strconv.ParseInt(m[2], 10, 64))
			if ren, ok := renameAlgs[a]; ok {
				if ren == "" {
					continue
				}
				a = ren
			}
			if alg != a {
				flush()
				alg = a
				sum = 0
				count = 0
			}
			sum += ns
			count++
		}
		flush()
	}

	var buf strings.Builder

	buf.WriteString("\n\n### MB/sec\n\n")

	printTable(&buf, algs, models, func(v int64) string {
		if v == 0 {
			return "n/a"
		}
		bandw := round10((1_000_000_000 + v/2) / v)
		return fint64(bandw)
	})

	buf.WriteString("\n(rounded to 10 MB/sec)\n")

	buf.WriteString("\n\n### us/MB\n\n")

	printTable(&buf, algs, models, func(v int64) string {
		if v == 0 {
			return "n/a"
		}
		us := (v + 500) / 1_000
		return fint64(us)
	})

	fmt.Fprintln(os.Stdout, buf.String())
}

func printTable(buf *strings.Builder, algs []string, models []*modelData, format func(v int64) string) {
	firstColWidth := max(minModelColWidth, len(modelColLabel))
	for _, data := range models {
		firstColWidth = max(firstColWidth, len(data.name))
	}

	// first row
	buf.WriteByte('|')
	buf.WriteByte(' ')
	buf.WriteString(rpad(modelColLabel, firstColWidth))
	buf.WriteByte(' ')
	buf.WriteByte('|')
	for _, alg := range algs {
		colWidth := max(minResultColWidth, len(alg))
		buf.WriteByte(' ')
		buf.WriteString(lpad(alg, colWidth))
		buf.WriteByte(' ')
		buf.WriteByte('|')
	}
	buf.WriteByte('\n')

	// separator row
	buf.WriteByte('|')
	buf.WriteByte(' ')
	buf.WriteByte(':')
	buf.WriteString(strings.Repeat("-", firstColWidth-1))
	buf.WriteByte(' ')
	buf.WriteByte('|')
	for _, alg := range algs {
		colWidth := max(minResultColWidth, len(alg))
		buf.WriteByte(' ')
		buf.WriteString(strings.Repeat("-", colWidth-1))
		buf.WriteByte(':')
		buf.WriteByte(' ')
		buf.WriteByte('|')
	}
	buf.WriteByte('\n')

	for _, data := range models {
		buf.WriteByte('|')
		buf.WriteByte(' ')
		buf.WriteString(rpad(data.name, firstColWidth))
		buf.WriteByte(' ')
		buf.WriteByte('|')
		for _, alg := range algs {
			colWidth := max(minResultColWidth, len(alg))
			buf.WriteByte(' ')
			value := format(data.algs[alg])
			buf.WriteString(lpad(value, colWidth))
			buf.WriteByte(' ')
			buf.WriteByte('|')
		}
		buf.WriteByte('\n')
	}

}

func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func isNL(r rune) bool {
	return r == '\n'
}

func lpad(s string, n int) string {
	if len(s) >= n {
		return s
	}
	return strings.Repeat(" ", n-len(s)) + s
}
func rpad(s string, n int) string {
	if len(s) >= n {
		return s
	}
	return s + strings.Repeat(" ", n-len(s))
}

func fint64(i int64) string {
	s := strconv.FormatInt(i, 10)
	n := len(s)
	if n > 3 {
		return s[:n-3] + " " + s[n-3:]
	}
	return s
}

func round10(v int64) int64 {
	return (v + 5) / 10 * 10
}
