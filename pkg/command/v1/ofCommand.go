package v1

import "context"

type OffCommand struct {
	Device
}

func (c *OffCommand) execute(ctx context.Context) {
	c.Device.off(ctx)
}
