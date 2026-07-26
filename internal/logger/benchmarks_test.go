package logger

import "testing"

func BenchmarkParseLevel(b *testing.B) {
	levels := []string{"DEBUG", "INFO", "WARN", "ERROR"}

	b.ReportAllocs()
	for _, level := range levels {
		b.Run(level, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if _, err := parseLevel(level); err != nil {
					b.Fatalf("parse level failed: %v", err)
				}
			}
		})
	}
}
