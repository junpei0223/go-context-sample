package session

import "context"

type ctxKey struct{}

var sequence int = 1

func SetSessionID(ctx context.Context) context.Context {
	idCtx := context.WithValue(ctx, ctxKey{}, sequence)
	sequence += 1
	return idCtx
}

func GetSessionID(ctx context.Context) int {
	id := ctx.Value(ctxKey{}).(int)
	return id
}
