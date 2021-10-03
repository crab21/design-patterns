package v1

import "testing"

func TestObserver(t *testing.T) {
	shirtItem := newItem("Nike Shirt")

	observerFirst := &Consumer{id: "abc@gmail.com"}
	observerSecond := &Consumer{id: "xyz@gmail.com"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.updateAvailability()
}
