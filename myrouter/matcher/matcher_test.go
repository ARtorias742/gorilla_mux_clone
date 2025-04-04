package matcher

import "testing"

func TestMatcher_Match(t *testing.T) {
	matcher := NewMatcher()

	tests := []struct {
		pattern string
		path    string
		want    bool
	}{
		{"/test", "/test", true},
		{"/test", "/other", false},
		{"/", "/", true},
	}

	for _, tt := range tests {
		got := matcher.Match(tt.pattern, tt.path)
		if got != tt.want {
			t.Errorf("Match(%q, %q) = %v, want %v", tt.pattern, tt.path, got, tt.want)
		}
	}
}
