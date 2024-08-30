package main

import (
	"testing"
)

func TestConvertToCron(t *testing.T) {
	t.Setenv("CRON_API_URL", "https://natural-cron-api.fly.dev")
	t.Setenv("BOT_TOKEN", "test")
	initEnv()

	// Test the convertToCron function
	cron, ok := convertToCron("every day at 9am")
	if !ok {
		t.Error("Expected ok to be true")
	}
	if cron != "0 9 * * ? *" {
		t.Errorf("Expected cron to be 0 9 * * ? *, got %s", cron)
	}
}
