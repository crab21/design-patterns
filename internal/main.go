package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// ctxChild, cancelChild := context.WithTimeout(ctx, 12*time.Second)
	// defer cancelChild()

	// fmt.Println(time.Now().Unix())
	// go func() {
	// 	go func(ctx context.Context) {
	// 		for {
	// 			i := 0
	// 			i++
	// 			time.Sleep(time.Second)
	// 			fmt.Println(i)
	// 		}
	// 	}(ctx)
	// 	select {
	// 	case <-ctx.Done():
	// 		fmt.Println(ctx.Err())
	// 		fmt.Println("parent:", time.Now().Unix())
	// 	case <-time.After(30 * time.Second):
	// 		fmt.Println("time out")
	// 	}
	// }()

	// go func() {
	// 	select {
	// 	case <-ctxChild.Done():
	// 		fmt.Println(ctxChild.Err())
	// 		fmt.Println("child:", time.Now().Unix())
	// 	case <-time.After(30 * time.Second):
	// 		fmt.Println("time out")
	// 	}
	// }()

	// c, cc := context.WithCancel(ctx)
	// cc()
	// time.Sleep(time.Second)
	// cancel()
	// select {}
	AsyncCall()

}

func AsyncCall() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*8))
	defer cancel()

	cc, ce := context.WithCancel(ctx)
	defer ce()
	go AsyncCall2(cc)
	select {
	case <-ctx.Done():
		fmt.Println("call successfully!!!")
		return
	case <-time.After(time.Duration(time.Second * 900)):
		// default:
		fmt.Println("timeout!!!")
		time.Sleep(time.Second * 900)
		return
	}
}

func AsyncCall2(ctx context.Context) {

	go func(ctx context.Context) {
		// 发送HTTP请求
		i := 0

		for {
			i++
			if i == 10 {
				break
			}
			time.Sleep(time.Second)
			fmt.Println(i)
		}
	}(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("call successfully  2!!!")
		return
	case <-time.After(time.Duration(time.Second * 900)):
		// default:
		fmt.Println("timeout!!!")
		time.Sleep(time.Second * 900)
		return
	}
}
