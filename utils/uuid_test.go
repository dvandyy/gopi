package utils

import (
	"reflect"
	"testing"
)

func TestGenerateUniqueID(t *testing.T) {
	if reflect.TypeOf(GenerateUniqueID("test-")) != reflect.TypeOf("") {
		t.Error("Unique ID should be a string.")
	}
}
