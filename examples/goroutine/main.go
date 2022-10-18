package main

import (
	"context"
	"fmt"
	"github.com/orccn/ho"
	"time"
)

func main() {
	// Go
	ho.Go(func() {
		panic("this is ho.Go")
	})
	time.Sleep(time.Millisecond * 100)

	// GoCtx
	ctx := context.WithValue(context.Background(), "traceID", "1234567")
	ho.GoCtx(ctx, func(ctx context.Context) {
		panic(fmt.Sprintf("this is ho.GoCtx, traceID: %s", ctx.Value("traceID")))
	})
	time.Sleep(time.Millisecond * 100)

	// GoCtxRecover
	ho.GoCtxRecover(ctx, func(ctx context.Context) {
		panic(fmt.Sprintf("this is ho.GoCtxRecover, traceID: %s", ctx.Value("traceID")))
	}, func(ctx context.Context) {
		if err := recover(); err != nil {
			fmt.Println("you can implements your recover function in here")
			fmt.Println(err)
		}
	})
	time.Sleep(time.Millisecond * 100)
}
