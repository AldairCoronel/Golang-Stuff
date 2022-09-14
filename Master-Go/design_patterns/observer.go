package main

import "fmt"

type Topic interface {
	register(observer Observer)
	broadcast()
}

type Observer interface {
	getId() string
	updateValue(string)
}

type EmailClient struct {
	id string
}

func (ec *EmailClient) getId() string {
	return ec.id
}

func (ec *EmailClient) updateValue(value string) {
	fmt.Printf("Sending email - %s available from client %s\n", value, ec.id)
}

// Item -> No stock
// Item -> Stock
type Item struct {
	observers []Observer
	name      string
	available bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) UpdateAvailable() {
	fmt.Printf("Item %s is available\n", i.name)
	i.available = true
	i.broadcast()
}

func (i *Item) broadcast() {
	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}
}

func (i *Item) register(observer Observer) {
	i.observers = append(i.observers, observer)
}

func main() {
	airbusItem := NewItem("A380")
	firstObserver := &EmailClient{
		id: "12a",
	}
	secondObserver := &EmailClient{
		id: "34dc",
	}

	airbusItem.register(firstObserver)
	airbusItem.register(secondObserver)
	airbusItem.UpdateAvailable()
}
