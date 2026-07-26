package database

import "testing"

func BenchmarkInitSchemaNil(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if err := InitSchema(nil); err == nil {
			b.Fatal("expected error for nil db")
		}
	}
}

func BenchmarkCloseNil(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if err := Close(nil); err != nil {
			b.Fatalf("expected nil error, got %v", err)
		}
	}
}
