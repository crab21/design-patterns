package v1

import "context"

type OnCommand struct {
	Device
}

func (c *OnCommand) execute(ctx context.Context) {
	c.Device.on(ctx)
}
