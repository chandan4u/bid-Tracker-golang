package lib

import (
	"testing"
)

func TestGetDefaultConfigDir(t *testing.T) {
	got, _ := GetDefaultConfigDir()
	if got != "/Users/config" {
		t.Errorf("Status should be true")
	}
}
