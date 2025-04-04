package matcher

import "testing"

func TestMatcher_MatchWithParams(t *testing.T) {
	matcher := NewMatcher()

	tests := []struct {
		pattern string
		path    string
		want    bool
		params  Params
	}{
		{"/users/{id}", "/users/123", true, Params{"id": "123"}},
		{"/users/{id}/edit", "/users/456/edit", true, Params{"id": "456"}},
		{"/test", "/other", false, nil},
		{"/", "/", true, Params{}},
	}

	for _, tt := range tests {
		params, matched := matcher.Match(tt.pattern, tt.path)
		if matched != tt.want {
			t.Errorf("Match(%q, %q) matched = %v, want %v", tt.pattern, tt.path, matched, tt.want)
		}
		if matched && !paramsEqual(params, tt.params) {
			t.Errorf("Match(%q, %q) params = %v, want %v", tt.pattern, tt.path, params, tt.params)
		}
	}
}

func paramsEqual(a, b Params) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
