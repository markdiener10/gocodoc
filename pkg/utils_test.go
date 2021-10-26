package gocodoc

import (
	"testing"
)

func TestIndex(t *testing.T) {

	//We are one directory under the root directory
	haystack := "lib/ar"
	needle := "b/"

	idx := 0

	idx = RevIat(haystack, needle, 0, 5)
	if idx != 2 {
		t.Error("IndexReverse Not Working:", idx)
	}

	idx = RevIat(haystack, needle, 1, 4)
	if idx != 2 {
		t.Error("IndexReverse with Offset 1 Not Working:", idx)
	}

	idx = Iat(haystack, needle, 0, -1)
	if idx != 2 {
		t.Error("Index with Offset 1 Not Working:", idx)
	}

	idx = Iat(haystack, needle, 1, 5)
	if idx != 2 {
		t.Error("Index with Offset 1 Not Working:", idx)
	}

}
