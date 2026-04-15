package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := &http.Header{}

	headers.Add("Authorization", "ApiKey 1234")

	expected := "1234"

	got, _ := GetAPIKey(*headers)

	if expected != got {
		t.Fatalf("Exppected %v, got %v", expected, got)
	}
}

func TestGetAPIKeyNoHeader(t *testing.T) {
	headers := &http.Header{}

	got, err := GetAPIKey(*headers)

	if got != "" || err == nil {
		t.Fatalf("Expected error to be thrown")
	}
}
