package helpers

import (
	"fmt"
	"testing"
)

func TestNewId(t *testing.T) {
	if NewId() == "" {
		t.Error("NewId() should not return empty string")
	}
}

func TestNewIdLengths(t *testing.T) {
	for i := 1; i < 64; i++ {
		t.Setenv("LNKSHRT_IDSIZE", fmt.Sprint(i))
		if len(NewId()) != i {
			t.Errorf("NewId() should return %d characters", i)
		}
	}
}

func TestNewIdRandom(t *testing.T) {
	id1 := NewId()
	id2 := NewId()
	if id1 == id2 {
		t.Error("NewId() should return different ids")
	}
}

func TestNewIdSymbols(t *testing.T) {
	t.Setenv("LNKSHRT_IDSIZE", "1024")
	id := NewId()
	for _, r := range id {
		if !(r >= 'a' && r <= 'z') && !(r >= 'A' && r <= 'Z') && !(r >= '0' && r <= '9') {
			t.Errorf("NewId() should contain only letters and numbers ([a-zA-Z0-9]), got %c", r)
		}
	}
}
