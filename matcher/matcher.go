package matcher

import (
	"context"
	"strings"
)

type Matcher struct{}

func NewMatcher() *Matcher {
	return &Matcher{}
}

type Params map[string]string

func (m *Matcher) Match(pattern, path string) (Params, bool) {
	patternParts := strings.Split(pattern, "/")
	pathParts := strings.Split(path, "/")

	if len(patternParts) != len(pathParts) {
		return nil, false
	}

	params := make(Params)
	for i, part := range patternParts {
		if len(part) > 0 && part[0] == '{' && part[len(part)-1] == '}' {
			paramName := part[1 : len(part)-1]
			params[paramName] = pathParts[i]
		} else if part != pathParts[i] {
			return nil, false
		}
	}

	return params, true
}

type contextKey string

const paramsKey contextKey = "params"

func WithParams(ctx context.Context, params Params) context.Context {
	return context.WithValue(ctx, paramsKey, params)
}

func GetParams(ctx context.Context) Params {
	if params, ok := ctx.Value(paramsKey).(Params); ok {
		return params
	}
	return nil
}
