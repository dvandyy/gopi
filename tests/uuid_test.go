package tests

import (
	"reflect"
	"testing"

	"github.com/dvandyy/gopi/utils"
)

func TestGenerateUniqueID(t *testing.T) {
	if reflect.TypeOf(utils.GenerateUniqueID("test-")) != reflect.TypeOf("") {
		t.Error("Unique ID should be a string.")
	}
}
