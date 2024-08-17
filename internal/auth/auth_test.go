package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	expected := ""

	actual, _ := GetAPIKey(http.Header{})

	if expected != actual {
		t.Errorf("Expected %v but actual %v", expected, actual)
	}
}
