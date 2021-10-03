package v1

import "fmt"

type item struct {
	observerList []Observer
	name         string
	inStock      bool
}

func newItem(name string) *item {
	return &item{
		name: name,
	}
}

func (item *item) register(observer Observer) {
	item.observerList = append(item.observerList, observer)
}

func (item *item) deregister(observer Observer) {
	removeFromObserverList(item.observerList, observer)
}

func removeFromObserverList(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

func (i *item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}

func (i *item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}
