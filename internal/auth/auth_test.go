package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		header      http.Header
		expectedKey string
		expectedErr error
	}{
		{
			name:        "Valid API Key",
			header:      http.Header{"Authorization": []string{"ApiKey abc123"}},
			expectedKey: "abc123",
			expectedErr: nil,
		},
		{
			name:        "No Authorization Header",
			header:      http.Header{},
			expectedKey: "",
			expectedErr: ErrNoAuthHeaderIncluded,
		},
		{
			name:        "Malformed Authorization Header",
			header:      http.Header{"Authorization": []string{"Bearer abc123"}},
			expectedKey: "",
			expectedErr: errors.New("malformed authorization header"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.header)
			if key != tt.expectedKey || (err != nil && err.Error() != tt.expectedErr.Error()) {
				t.Errorf("Test %s failed: got (%v, %v), expected (%v, %v)", tt.name, key, err, tt.expectedKey, tt.expectedErr)
			}
		})
	}
}
