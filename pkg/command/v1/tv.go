package v1

import (
	"context"
	"fmt"
)

type Tv struct {
	IsRunning bool
}

func (Tv *Tv) on(ctx context.Context) {
	Tv.IsRunning = true
	fmt.Println(Tv.IsRunning, "status: is on")
}
func (Tv *Tv) off(ctx context.Context) {
	Tv.IsRunning = false
	fmt.Println(Tv.IsRunning, " status: is off")
}
