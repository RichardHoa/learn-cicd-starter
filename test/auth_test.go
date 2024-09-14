package test

import (
	"errors"
	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
	"net/http"
	"testing"
)

func TestAuth(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		want        string
		expectedErr error
	}{
		{name: "Correct syntax", input: "ApiKey 1234", want: "1234", expectedErr: nil},
		{name: "Missing ApiKey", input: "1234", want: "",
		 expectedErr: errors.New("malformed authorization header")},
	}

	for _, testCase := range tests {
		got, err := auth.GetAPIKey(http.Header{"Authorization": {testCase.input}})
		if got != testCase.want  {
			t.Errorf("GetAPIKey(%s) = %s, want %s ",
				testCase.input, got, testCase.want)
		}
		if err != nil && err.Error() != testCase.expectedErr.Error() {
			t.Errorf("GetAPIKey(%s) =  want error '%v' but get error '%v'",
				testCase.input, testCase.expectedErr, err)
		}
	}

}
