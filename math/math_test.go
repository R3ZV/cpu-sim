package math

import (
	"testing"
)

func TestGCD(t *testing.T) {
	want := 5
	got := GCD(20, 5)
	if got != want {
		t.Fatalf(`Expected = %d, Got = %d`, want, got)
	}

	want = 1
	got = GCD(3, 5)
	if got != want {
		t.Fatalf(`Expected = %d, Got = %d`, want, got)
	}

	want = 6
	got = GCD(6, 6)
	if got != want {
		t.Fatalf(`Expected = %d, Got = %d`, want, got)
	}
}

func TestLCM(t *testing.T) {
	want := 20
	got := LCM(20, 5)
	if got != want {
		t.Fatalf(`Expected = %d, Got = %d`, want, got)
	}

	want = 15
	got = LCM(3, 5)
	if got != want {
		t.Fatalf(`Expected = %d, Got = %d`, want, got)
	}

	want = 6
	got = LCM(6, 6)
	if got != want {
		t.Fatalf(`Expected = %d, Got = %d`, want, got)
	}
}
