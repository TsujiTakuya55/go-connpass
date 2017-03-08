package connpass

import "testing"

func testNewClient(t *testing.T) {
	client := NewClient()

	if got, want := client.BaseURL.String(), defaultBaseURL; got != want {
		t.Errorf("NewClient BaseURL is %v, want %v", got, want)
	}
}