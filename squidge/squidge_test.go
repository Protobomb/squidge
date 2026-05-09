package squidge

import (
	"strings"
	"testing"
)

func TestSquidge(t *testing.T) {
	got := Squidge()
	want := "Squidge!"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestSquidgeProperties(t *testing.T) {
	out := Squidge()

	t.Run("non-empty", func(t *testing.T) {
		if out == "" {
			t.Error("Squidge() returned empty string")
		}
	})

	t.Run("ends with bang", func(t *testing.T) {
		if !strings.HasSuffix(out, "!") {
			t.Errorf("Squidge() = %q, want suffix %q", out, "!")
		}
	})

	t.Run("starts capitalized", func(t *testing.T) {
		first := out[:1]
		if first != strings.ToUpper(first) {
			t.Errorf("Squidge() = %q, want first char to be uppercase", out)
		}
	})
}
