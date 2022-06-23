package auth

import (
	"context"
)

type authKey struct{}

func NewContext(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, authKey{}, id)
}

func FromContext(ctx context.Context) (id string, ok bool) {
	id, ok = ctx.Value(authKey{}).(string)
	return
}
