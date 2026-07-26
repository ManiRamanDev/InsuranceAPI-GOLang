package models

import (
	"reflect"
	"testing"
)

func BenchmarkModelStructFieldInspection(b *testing.B) {
	types := []any{Customer{}, Policy{}, CustomerPolicy{}, Claim{}}

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, value := range types {
			typeOfValue := reflect.TypeOf(value)
			if typeOfValue.NumField() == 0 {
				b.Fatalf("expected fields on %s", typeOfValue.Name())
			}
			if field, ok := typeOfValue.FieldByName("ID"); !ok || field.Tag.Get("gorm") == "" {
				b.Fatalf("expected gorm tag on %s.ID", typeOfValue.Name())
			}
		}
	}
}
