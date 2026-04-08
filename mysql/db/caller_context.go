package db

import "context"

type callerContextKey struct{}

func WithCaller(ctx context.Context, caller string) context.Context {
	return context.WithValue(ctx, callerContextKey{}, caller)
}

func CallerFromContext(ctx context.Context) (string, bool) {
	if ctx == nil {
		return "", false
	}
	caller, ok := ctx.Value(callerContextKey{}).(string)
	return caller, ok
}
