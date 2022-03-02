package services_test

import (
	"testing"

	"github.com/arthuramorim04/go-activity-api.git/services"
)

func TestSHA256Encode(t *testing.T) {
	if services.SHA256Encode("123") == "123" {
		t.Errorf("SHA256Encode not work")
	}
}
