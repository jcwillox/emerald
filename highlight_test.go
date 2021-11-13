package emerald

import (
	"os"
	"path/filepath"
	"testing"
)

func BenchmarkHighlightFile(b *testing.B) {
	path := "my_file.class"
	b.Log(HighlightFile(path))
	for i := 0; i < b.N; i++ {
		HighlightFile(path)
	}
}

func BenchmarkHighlightFile2(b *testing.B) {
	path := "my_file.class"
	b.Log(HighlightFile(path, 0777))
	for i := 0; i < b.N; i++ {
		HighlightFile(path, 0777)
	}
}

func BenchmarkHighlightFileStat(b *testing.B) {
	path := "print.go"
	filename := filepath.Base(path)
	stat, _ := os.Stat(path)
	//mode := stat.Mode()
	b.Log(HighlightFileStat(filename, stat))
	for i := 0; i < b.N; i++ {
		HighlightFileStat(filename, stat)
	}
}

func BenchmarkHighlightPath(b *testing.B) {
	path := "/hello/world/print.go"
	b.Log(HighlightPath(path))
	for i := 0; i < b.N; i++ {
		HighlightPath(path)
	}
}

func BenchmarkHighlightPath2(b *testing.B) {
	path := "/hello/world/print.go"
	b.Log(HighlightPath(path, 0777))
	for i := 0; i < b.N; i++ {
		HighlightPath(path, 0777)
	}
}

func BenchmarkHighlightPath3(b *testing.B) {
	path := "/hello/world/README.md"
	b.Log(HighlightPath(path))
	for i := 0; i < b.N; i++ {
		HighlightPath(path)
	}
}

func BenchmarkHighlightPathStat(b *testing.B) {
	path := "print.go"
	stat, _ := os.Stat(path)
	b.Log(HighlightPathStat(path, stat))
	for i := 0; i < b.N; i++ {
		HighlightPathStat(path, stat)
	}
}

func BenchmarkHighlightPathStat2(b *testing.B) {
	path := "README.md"
	stat, _ := os.Stat(path)
	b.Log(HighlightPathStat(path, stat))
	for i := 0; i < b.N; i++ {
		HighlightPathStat(path, stat)
	}
}

func BenchmarkHighlightPathNonExistent(b *testing.B) {
	path := "/hello/world"
	stat, _ := os.Stat(path)
	b.Log(HighlightPathStat(path, stat))
	for i := 0; i < b.N; i++ {
		HighlightPathStat(path, stat)
	}
}
