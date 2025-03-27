package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	testCases := []struct {
		name     string
		headers  http.Header
		expected string
		wantErr  bool
	}{
		{
			name: "valid API key",
			headers: http.Header{
				"Authorization": []string{"ApiKey valid-key-123"},
			},
			expected: "valid-key-123",
			wantErr:  false,
		},
		{
			name: "missing Authorization header",
			headers: http.Header{
				// No Authorization header
			},
			expected: "",
			wantErr:  true,
		},
		{
			name: "malformed authorization header - wrong prefix",
			headers: http.Header{
				"Authorization": []string{"Bearer some-key"},
			},
			expected: "",
			wantErr:  true,
		},
		{
			name: "malformed authorization header - no space",
			headers: http.Header{
				"Authorization": []string{"ApiKeysome-key"},
			},
			expected: "",
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := GetAPIKey(tc.headers)
			
			// Check if we got an error when we wanted one
			if (err != nil) != tc.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			
			// Check if the result matches what we expected
			if result != tc.expected {
				t.Errorf("GetAPIKey() = %v, want %v", result, tc.expected)
			}
		})
	}
}