package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		headers http.Header
		want    string
		wantErr bool
	}{
		"simple":      {headers: http.Header{"Authorization": []string{"ApiKey hunter2"}}, want: "hunter2", wantErr: false},
		"empty":       {headers: http.Header{"Authorization": []string{"ApiKey"}}, want: "", wantErr: true},
		"wrongHeader": {headers: http.Header{"AC": []string{"ApiKey"}}, want: "", wantErr: true},
	}

	for name, tc := range tests {
		got, err := GetAPIKey(tc.headers)
		if !reflect.DeepEqual(tc.want, got) && !(tc.wantErr && err != nil) {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
		}
	}
}
