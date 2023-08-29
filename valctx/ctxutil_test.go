package valctx_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/tuannguyenandpadcojp/utils/valctx"
)

func TestValueOnly(t *testing.T) {
	// t.Run("keep values of the parent context", func(t *testing.T) {
	// 	var key struct{}
	// 	ctx := context.WithValue(context.Background(), key, "bar")
	// 	ctx = valctx.ValueOnly(ctx)

	// 	require.Equal(t, "bar", ctx.Value(key))
	// })

	t.Run("remove cancellation from parent context", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		ctx = valctx.ValueOnly(ctx)
		cancel()
		require.NoError(t, ctx.Err())
	})

	t.Run("remove deadline from parent context", func(t *testing.T) {
		ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Minute))
		t.Cleanup(cancel)

		ctx = valctx.ValueOnly(ctx)
		_, ok := ctx.Deadline()
		require.False(t, ok)
	})

	t.Run("remove done from parent context", func(t *testing.T) {
		ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Minute))
		ctx = valctx.ValueOnly(ctx)
		cancel()
		require.Nil(t, ctx.Done())
	})
}
