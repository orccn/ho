package ho

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

//go:generate gogen option -n App -s recover,ctx --with-prefix --with-init
type App struct {
	wg      sync.WaitGroup
	recover ContextFunc
	ctx     context.Context
	cancel  context.CancelFunc
}

// TODO Restart
// func Restart()

func (a *App) init() {
	if a.ctx == nil {
		a.ctx = context.Background()
	}
	a.ctx, a.cancel = context.WithCancel(a.ctx)
	if a.recover == nil {
		a.recover = DefaultRecover
	}
}

func (a *App) Wait(close func() error, signals ...os.Signal) (err error) {
	ch := make(chan os.Signal, 1)
	if len(signals) == 0 {
		signals = []os.Signal{syscall.SIGTERM, syscall.SIGINT}
	}
	signal.Notify(ch, signals...)
	<-ch
	a.cancel()
	a.wg.Wait()
	if close != nil {
		err = close()
	}
	return
}

func (a *App) Go(f ContextFunc) {
	a.wg.Add(1)
	GoCtx(a.ctx, func(ctx context.Context) {
		defer a.wg.Done()
		f(ctx)
	})
}
