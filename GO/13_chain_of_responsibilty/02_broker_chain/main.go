package main

import (
	"fmt"
	"sync"
)

// CoR, Mediator, Observer, CQS

// when you want to know a creature's attack or defence value, you make a query, which is a separate struct and tou send this struct to the creature as opposed to just calling a method on it

type Argument int

const (
	Attack Argument = iota
	Defense
)

// CQS
type Query struct {
	CreatureName string
	WhatToQuery  Argument
	Value        int
}

// Observer - listens on event and do something
type Observer interface {
	Handle(q *Query)
}

// implemented by whatever type which wants to notify other about something happening. The event in this case will be the Query
type Observable interface {
	Subscribe(o Observer)   //start listening to events(Queries)
	Unsubscribe(o Observer) // stop listening to events(Queries)
	Fire(q *Query)
}

// Game is something that every participant is going to subscribe
type Game struct {
	observers sync.Map // sync map is going to allow us to basically keep a map of every single subscriber and to iterate this map to go through every single subscriber and notify that subscriber
}

// Add the observer received to the list of observers
func (g *Game) Subscribe(o Observer) {
	g.observers.Store(o, struct{}{}) // we just care about the key
}

// Removes the observer received to the list of observers
func (g *Game) Unsubscribe(o Observer) {
	g.observers.Delete(o)
}

// execute action for every observer
func (g *Game) Fire(q *Query) {
	g.observers.Range(func(key interface{}, value interface{}) bool {
		if key == nil {
			return false
		}
		key.(Observer).Handle(q)
		return true
	})
}

type Creature struct {
	game            *Game // mediator
	Name            string
	attack, defense int // only store initail values
}

func NewCreature(game *Game, name string, attack, defense int) *Creature {
	return &Creature{
		game:    game,
		attack:  attack,
		defense: defense,
		Name:    name,
	}
}

func (c *Creature) Attack() int {
	// in order to get the value, we make a query object. We get the game to fire the query. The subscribers get to process the query and modify the final attack value and then we return that value.
	q := &Query{
		CreatureName: c.Name,
		WhatToQuery:  Attack,
		Value:        c.attack,
	}

	c.game.Fire(q)

	return q.Value
}

func (c *Creature) Defense() int {
	// in order to get the value, we make a query object. We get the game to fire the query. The subscribers get to process the query and modify the final defense value and then we return that value.
	q := &Query{
		CreatureName: c.Name,
		WhatToQuery:  Defense,
		Value:        c.defense,
	}

	c.game.Fire(q)

	return q.Value
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack(), c.Defense())
}

// CoR
// CreatureModifier is simply a template
// Modifier == Observer
type CreatureModifier struct {
	game     *Game
	creature *Creature
}

func (cm *CreatureModifier) Handle(q *Query) {
	//
}

// Modifier == Observer
type DoubleAttackModifier struct {
	CreatureModifier
}

func (dam *DoubleAttackModifier) Handle(q *Query) {
	if q.CreatureName == dam.creature.Name && q.WhatToQuery == Attack {
		q.Value *= 2
	}
}

// Can be used to unsubscribe this modifier from the game event
func (dam *DoubleAttackModifier) Close() {
	dam.game.Unsubscribe(dam)
}

func NewDoubleAttackModifier(g *Game, c *Creature) *DoubleAttackModifier {
	dam := &DoubleAttackModifier{
		CreatureModifier: CreatureModifier{
			game:     g,
			creature: c,
		},
	}
	g.Subscribe(dam)
	return dam
}

func main() {
	// central mediator/Observable - manages subsribe/unsubcribe and Firing
	game := &Game{
		observers: sync.Map{},
	}

	// creature
	goblin := NewCreature(game, "Green Goblin", 10, 10)
	fmt.Println(goblin.String())

	{
		dam := NewDoubleAttackModifier(game, goblin)
		fmt.Println(goblin.String())
		dam.Close()
	}

	fmt.Println(goblin.String())
}

// Create the observable and mediator - Game
//  			v
// Create a new creature - certain entity with is there for store data of name, attack, defense and which game is belongs to
//				v
// We call NewDoubleAttackModifier which creates a new DoubleAttackModifier which is an observer since it implements the Handle(q *Query) func and then subscrive it to the game(observable)'s list of observers
//				v
//	We execute a creatures's String(). This function triggers the queries/events of Attack() and Defense().
//				v
// These func form a query/event, fire it
//				v
// The game's fire function notifies all the observers by executing their Handle function
//				v
//	An oberver/modifer like DoubleAttackModifier, in their Handle func, check if the event/query is for their creature. If yes, then they modify the value of the query they received
//				v
// After all the observers are iterated over for an event/query, the fire func returns the value of the query
// 				v
// Then we call the DoubleAttackModifier(observer)'s Close() func, which calls the Game(Observable)'s unsubscribe func()
// 				v
// The unsubscribe func removes the DoubleAttackModifier(observer)from the map of obervers of the game. So in case another event/query is fired the unsubscribed observer won't receive this event/query
