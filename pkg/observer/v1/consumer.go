package v1

import "fmt"

type Consumer struct {
	id string
}

func (c *Consumer) update(name string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, name)
}

func (c *Consumer) getID() string {
	return c.id
}
