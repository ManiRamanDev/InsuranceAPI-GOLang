package models

import (
	"reflect"
	"testing"
)

func TestModelStructFields(t *testing.T) {
	types := []any{Customer{}, Policy{}, CustomerPolicy{}, Claim{}}
	for _, value := range types {
		typeOfValue := reflect.TypeOf(value)
		if typeOfValue.NumField() == 0 {
			t.Fatalf("expected fields on %s", typeOfValue.Name())
		}
		if field, ok := typeOfValue.FieldByName("ID"); !ok || field.Tag.Get("gorm") == "" {
			t.Fatalf("expected gorm tag on %s.ID", typeOfValue.Name())
		}
	}
}
