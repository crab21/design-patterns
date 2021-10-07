package v1

import "context"

type button struct {
	Command
}

func (b *button) Press() {
	b.Command.execute(context.Background())

}
