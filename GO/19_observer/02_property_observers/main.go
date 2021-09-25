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
	p.age = age
	p.NotifySubscribers(
		PropertyChange{
			Name:  "Age",
			Value: p.age,
		},
	)
}

// traffic management observer's a person's age till they are old enough to drive and then they congratulate them and then unsubscribe that person
type TrafficManagement struct {
	ob Observable
}

func (tm *TrafficManagement) Notify(data interface{}) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Value.(int) >= 18 {
			fmt.Println("Congrats! You can drive now.")
			tm.ob.Unsubscribe(tm)
		}
	}
}

func main() {
	p := NewPerson(15)
	tm := &TrafficManagement{
		ob: p.Observable,
	}

	p.Subscribe(tm)

	for i := 16; i <= 20; i++ {
		fmt.Println("Setting the age to:", i)
		p.SetAge(i)
	}
}
