package utils

import (
	"testing"
)

func TestMessage(t *testing.T) {
	got := Message(true, "Hello World")
	if got["status"] != true || got["message"] != "Hello World" {
		t.Errorf("Status should be true")
	}
}

func TestAddBidRequestValidation(t *testing.T) {
	got := AddBidRequestValidation("chandan", "docker", 10)
	if got["status"] != true || got["message"] != "User input successfully validate." {
		t.Errorf("Validation should be true")
	}
}

func TestUserRequestValidation(t *testing.T) {
	got := UserRequestValidation("chandan")
	if got["status"] != true || got["message"] != "User input successfully validate." {
		t.Errorf("Validation should be true")
	}
}

func TestUserRequestValidationFailCase(t *testing.T) {
	got := UserRequestValidation("c")
	if got["status"] != false || got["message"] != "Username should greater than 1 character." {
		t.Errorf("Validation should be true")
	}
}
