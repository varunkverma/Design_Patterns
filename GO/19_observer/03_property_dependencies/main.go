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
	age int
}

func NewPerson(age int) *Person {
	return &Person{
		age: age,
		Observable: Observable{
			subscribers: new(list.List),
		},
	}
}

type PropertyChange struct {
	Name  string      // name of the property
	Value interface{} // new value of the property
}

// Person's age getter()
func (p *Person) Age() int {
	return p.age
}

// Person's age setter()
func (p *Person) SetAge(age int) {
	if age == p.age {
		return
	}

	oldCanVote := p.CanVote()

	p.age = age

	if oldCanVote != p.CanVote() {
		p.NotifySubscribers(
			PropertyChange{
				Name:  "CanVote",
				Value: p.CanVote(),
			},
		)
	}

}

// Now the issue is we also need a notification for when a person can vote. But the notification sending is not done in getters, but in setters
func (p *Person) CanVote() bool {
	return p.age >= 18
}

type ElectoralRoll struct{}

func (er *ElectoralRoll) Notify(data interface{}) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Name == "CanVote" && pc.Value.(bool) {
			fmt.Println("Congrats! You can vote")
		}
	}
}

func main() {
	p := NewPerson(15)

	er := &ElectoralRoll{}

	p.Subscribe(er)

	for i := 16; i < 20; i++ {
		fmt.Println("Setting the age to", i)
		p.SetAge(i)
	}
}
