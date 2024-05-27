package words

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	testcases := []string{"Eduardo Raider", " ", "!1234$"}
	for _, orig := range testcases {
		rev, err := Reverse(orig)
		if err != nil {
			return
		}

		rev2, err := Reverse(rev)
		if err != nil {
			return
		}

		if orig != rev2 {
			t.Errorf("Reverse(%q) == %q, want %q", orig, rev2, rev)
		}
	}
}

func FuzzReverse(f *testing.F) {
	testcases := []string{"Eduardo Raider", " ", "!1234$"}
	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, orig string) {
		rev, err := Reverse(orig)
		if err != nil {
			return
		}

		rev2, err := Reverse(rev)
		if err != nil {
			return
		}

		if orig != rev2 {
			t.Errorf("Reverse(%q) == %q, want %q", orig, rev2, rev)
		}

		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse(%q) == %q, want valid", orig, rev)
		}
	})
}
