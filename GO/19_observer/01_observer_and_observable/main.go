package main

import (
	"container/list"
	"fmt"
)

// observer
type Observer interface {
	Notify(data interface{})
}

// observable
type Observable struct {
	subscribers *list.List
}

func (ob *Observable) Subscribe(observer Observer) {
	ob.subscribers.PushBack(observer)
}

func (ob *Observable) Unsubscribe(observer Observer) {
	for ele := ob.subscribers.Front(); ele != nil; ele = ele.Next() {
		if ele.Value.(Observer) == observer {
			ob.subscribers.Remove(ele)
		}
	}
}

func (ob *Observable) NotifySubscribers(data interface{}) {
	for ele := ob.subscribers.Front(); ele != nil; ele = ele.Next() {
		ele.Value.(Observer).Notify(data)
	}
}

// acting as Observable; trigger events
type Person struct {
	Observable
	Name string
}

func NewPerson(name string) *Person {
	return &Person{
		Name: name,
		Observable: Observable{
			subscribers: new(list.List),
		},
	}
}

func (p *Person) CatchACold() {
	p.NotifySubscribers(p.Name)
}

// acting as Observer; listens to events
type DoctorService struct{}

func (d *DoctorService) Notify(data interface{}) {
	fmt.Printf("A doctor has been called for %s\n", data.(string))
}

func main() {
	p := NewPerson("Boris") // Observable
	ds := &DoctorService{}  //observer

	p.Subscribe(ds) // observer is subscrobed to this event

	// event
	p.CatchACold()
}
