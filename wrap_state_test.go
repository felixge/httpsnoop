package httpsnoop

import (
	"reflect"
	"testing"
)

func TestSharedStateLayout(t *testing.T) {
	fields := reflect.TypeOf(rwFields{})
	for _, value := range []any{rwState{}, rw0{}, rw511{}} {
		typ := reflect.TypeOf(value)
		if typ.NumField() != 1 {
			t.Fatalf("%s has %d fields, want one shared field", typ, typ.NumField())
		}
		field := typ.Field(0)
		if !field.Anonymous || field.Type != fields || field.Offset != 0 {
			t.Errorf("%s does not embed rwFields at offset zero", typ)
		}
		if typ.Size() != fields.Size() || typ.Align() != fields.Align() {
			t.Errorf("%s changes the shared state's size or alignment", typ)
		}
	}
}
