package main

import (
	"testing"
	"time"
)

func TestHelloResponse(t *testing.T) {
	now := time.Date(2024, 1, 2, 15, 4, 5, 0, time.UTC)
	result := getResponse(now)
	expected := "Hello World! The time is 2024-01-02 15:04:05 +0000 UTC"

	if string(result) != expected {
		t.Errorf("expected %q, got: %q", expected, result)
	}
}
