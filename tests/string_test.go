package utils_test

import (
	"reflect"
	"testing"

	"github.com/codify-community/internals/codify-updater/utils"
)

func TestRemoveEmptyStringsFromSlice(t *testing.T) {
	original := []string{"Hello", "", "World"}
	removed := utils.RemoveEmptyStringsFromSlice(original)

	if reflect.DeepEqual(original, removed) {
		t.Errorf("slices are equal! %v == %v", original, removed)
	}

	if !reflect.DeepEqual(removed, []string{"Hello", "World"}) {
		t.Errorf("unexpected value: %v != [Hello World]", removed)
	}
}
