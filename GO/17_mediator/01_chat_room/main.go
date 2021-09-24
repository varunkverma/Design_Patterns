package main

import "fmt"

// mediator
type ChatRoom struct {
	people []*Person
}

func (cr *ChatRoom) Broadcast(source, message string) {
	for _, p := range cr.people {
		if p.Name != source {
			p.Receive(source, message)
		}
	}
}

func (cr *ChatRoom) Message(source, receiver, message string) {
	for _, p := range cr.people {
		if p.Name == receiver {
			p.Receive(source, message)
		}
	}
}

func (cr *ChatRoom) Join(p *Person) {
	joinMessage := p.Name + " has joined the chat\n"
	cr.Broadcast("Room", joinMessage)

	p.Room = cr
	cr.people = append(cr.people, p)
}

type Person struct {
	Name    string
	Room    *ChatRoom // reference to the mediator
	chatlog []string
}

func NewPerson(name string) *Person {
	return &Person{Name: name}
}

func (p *Person) Receive(sender, message string) {
	s := fmt.Sprintf("%s: %s\n", sender, message)
	fmt.Printf("[%s's chat session]: %s\n", p.Name, s)
	p.chatlog = append(p.chatlog, s)
}

// Broadcast
func (p *Person) Say(message string) {
	p.Room.Broadcast(p.Name, message)
}

func (p *Person) PrivateMessage(receiver, message string) {
	p.Room.Message(p.Name, receiver, message)
}

func main() {
	room := ChatRoom{}

	holly := NewPerson("Holly")
	mike := NewPerson("Mike")
	dwight := NewPerson("Dwight")

	room.Join(dwight)
	room.Join(mike)

	mike.Say("New HR is joining")
	dwight.PrivateMessage("Mike", "Shall we destroy her?")
	mike.PrivateMessage("Dwight", "NO DWIGHT!!!")

	room.Join(holly)
	holly.Say("Hey Guys")
	mike.Say("I love...HRs")
	dwight.Say("Meh")

	mike.PrivateMessage("Holly", "Ignore Dwight")
}
