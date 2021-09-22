package main

import "fmt"

type Creature struct {
	Name            string
	Attach, Defence int
}

func NewCreature(name string, attack int, defence int) *Creature {
	return &Creature{
		Name:    name,
		Attach:  attack,
		Defence: defence,
	}
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attach, c.Defence)
}

// defining what a chain of responsibilty is
type Modifier interface {
	Add(m Modifier) // chaininf func
	Handle()
}

// This is what a Chain of Resposibilty for a Type looks like
type CreatureModifier struct {
	creature *Creature
	next     Modifier
}

func NewCreatureModifier(creature *Creature) *CreatureModifier {
	return &CreatureModifier{
		creature: creature,
	}
}

func (cm *CreatureModifier) Add(m Modifier) {
	if cm.next != nil {
		cm.next.Add(m)
	} else {
		cm.next = m
	}
}

// This function doesn't really have a purpose, its just for perfoming aggregating. It simply moves and execute the next handle
func (cm *CreatureModifier) Handle() {
	if cm.next != nil {
		cm.next.Handle()
	}
}

// Specific types of the base. They use base type's method to chain on modifiers but have their own way to handle it.
type DoubleAttackModifier struct {
	CreatureModifier
}

func NewDoubleAttackModifier(c *Creature) *DoubleAttackModifier {
	return &DoubleAttackModifier{
		CreatureModifier{
			creature: c,
		},
	}
}

func (dam *DoubleAttackModifier) Handle() {
	fmt.Println("Doubling", dam.creature.Name, "\b's attack")
	dam.creature.Attach *= 2
	dam.CreatureModifier.Handle()
}

type IncreaseDefenseModifier struct {
	CreatureModifier
}

func NewIncreaseDefenseModifier(creature *Creature) *IncreaseDefenseModifier {
	return &IncreaseDefenseModifier{
		CreatureModifier: CreatureModifier{
			creature: creature,
		},
	}
}

func (idm *IncreaseDefenseModifier) Handle() {
	if idm.creature.Attach <= 50 {
		fmt.Println("Increasing", idm.creature.Name, "\b's defense")
		idm.creature.Defence += 5
	}
	idm.CreatureModifier.Handle()
}

type NoPowerUpsModifier struct {
	CreatureModifier
}

func NewNoPowerUpsModifier(creature *Creature) *NoPowerUpsModifier {
	return &NoPowerUpsModifier{
		CreatureModifier: CreatureModifier{
			creature: creature,
		},
	}
}

func (npm *NoPowerUpsModifier) Handle() {
	fmt.Println("Disabling any more power ups for", npm.creature.Name)
	//  Since this type's doesn't implement/call the logic to pass the control to the next modifier. The chain stops here.
}

func main() {
	goblin := NewCreature("Goblin", 10, 10)

	fmt.Println(goblin.String())

	root := NewCreatureModifier(goblin) // This creature is added in the base root modifier

	root.Add(NewDoubleAttackModifier(goblin)) // chaining a particular type of modifier in the root modifier

	root.Add(NewNoPowerUpsModifier(goblin)) // This stops the chain processing

	root.Add(NewDoubleAttackModifier(goblin)) // chaining a particular type of modifier in the root modifier

	root.Add(NewIncreaseDefenseModifier(goblin))

	root.Handle() // the oroot executes the base type's Handle, which in turn call the chained modifiers and they inturn call their next modifier's handle

	fmt.Println(goblin.String())
}
