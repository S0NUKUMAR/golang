package main

import (
	"fmt"
	"time"
	"context"
)

func main(){
	ctx,cancel:= context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()
	t := make(chan time.Time)
	ticker:= time.NewTicker(time.Second * 1)
	go clock(ctx,t,ticker)
	for i := range t {
			fmt.Println(i)
	}
}

func clock(ctx context.Context, tm chan<- time.Time, ticker *time.Ticker) {
		defer close(tm)
		for {
			select{
			case <-ctx.Done():
				ticker.Stop()
				return
			case t:= <-ticker.C:
				tm <-t
			}
	}
}
