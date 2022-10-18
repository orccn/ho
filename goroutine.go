package ho

import (
	"context"
	"log"
	"runtime"
)

// Go run a goroutine
func Go(f func()) {
	go func() {
		defer DefaultRecover(context.Background())
		f()
	}()
}

// GoCtx run a goroutine with context
func GoCtx(ctx context.Context, f ContextFunc) {
	GoCtxRecover(ctx, f, DefaultRecover)
}

// GoCtxRecover run a goroutine with context, specify recover function
func GoCtxRecover(ctx context.Context, f, recover ContextFunc) {
	go func() {
		defer recover(ctx)
		f(ctx)
	}()
}

// ContextFunc a function with context.context
type ContextFunc func(ctx context.Context)

// DefaultRecover default recover function
var DefaultRecover = func(ctx context.Context) {
	if err := recover(); err != nil {
		trace := make([]byte, 1<<16)
		runtime.Stack(trace, false)
		log.Printf("panic: %v\n", err)
		log.Printf("%s\n", trace)
	}
}
