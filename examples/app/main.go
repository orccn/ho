package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/orccn/ho"
	"time"
)

func main() {
	a := ho.NewApp()
	a.Go(func(ctx context.Context) {
		var i int
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine A done")
				return
			default:
				i++
				fmt.Println("goroutine A: ", i)
				time.Sleep(time.Second)
			}
		}
	})
	a.Go(func(ctx context.Context) {
		var i int
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine B done")
				return
			default:
				i++
				fmt.Println("goroutine B: ", i)
				time.Sleep(time.Second)
			}
		}
	})
	fmt.Println(a.Wait(func() (error error) {
		return errors.New("this is close function")
	}))
}
