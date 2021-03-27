package test

import (
	"context"
	"testing"
)

func Context(t *testing.T) context.Context {
	if deadline, ok := t.Deadline(); ok {
		ctx, cancel := context.WithDeadline(context.Background(), deadline)
		t.Cleanup(cancel)
		return ctx
	}
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	return ctx
}
