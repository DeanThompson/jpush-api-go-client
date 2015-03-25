package common

import "testing"

func Test_UniqString(t *testing.T) {
	s := []string{"a", "a", "a", "b"}
	uniq := UniqString(s)
	if !EqualStringSlice(uniq, []string{"a", "b"}) {
		t.Error("UniqString does not work correctly")
	}
}
