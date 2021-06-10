package uuid

import (
	"testing"
	"regexp"
)

func TestNewV4(t *testing.T) {
	t.Run("creates [16]byte where all its values are nonzero", func(t *testing.T) {
		uuid, _ := NewV4()
		allZero := true
		for _, v := range uuid {
			if v != byte(0) {
				allZero = false
				break
			}
		}
		if allZero {
			t.Error("all the fields of the UUID are zero")
		}
	})
}

func TestUUID(t *testing.T) {
	t.Run("it returns a proper string representation", func(t *testing.T) {
		uuid, _ := NewV4()
		pattern := regexp.MustCompile(`[A-Za-z0-9]{8}\-[A-Za-z0-9]{4}\-[A-Za-z0-9]{4}\-[A-Za-z0-9]{4}\-[A-Za-z0-9]{12}`)
		got := uuid.String()
		t.Logf("got %s", got)
		if !pattern.MatchString(got) {
			t.Errorf("the string representation of the UUID doesn't match, got: %s", got)
		}
	})
}
