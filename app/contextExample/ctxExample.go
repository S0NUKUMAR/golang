package main

import (
	"fmt"
	"time"
	"context"
)

func main(){
	ctx,cancel:= context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()
	t := clock(ctx)
	fmt.Println("time::", time.Now().Format("15:04:03"))
	for {
		select {
		case <-ctx.Done():
			return
		case tick:= <-t :
			fmt.Println(tick)
		}
	}
}

func clock(ctx context.Context) <-chan time.Time {
	ticker:= time.NewTicker(time.Second * 1)
	t := make(chan time.Time)
	go func() {
		defer close(t)
		for {
			select{
			case <-ctx.Done():
				ticker.Stop()
				return
			case tm:= <-ticker.C:
				t <-tm
			}
		}}() 
		return t
	}
