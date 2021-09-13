package main

import "fmt"

type Relationship int

const (
	PARENT Relationship = iota
	CHILD
	SIBLING
)

type Person struct {
	name string
}

type RelativeInfo struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// Interface
type RelationshipBrowser interface {
	FindAllChildrenOfParent(name string) []*Person
}

// Low Level Module
type Relationships struct {
	relations []RelativeInfo
}

func (r *Relationships) AddParentAndChildRelation(parent, child *Person) {
	r.relations = append(r.relations, RelativeInfo{
		from:         parent,
		relationship: PARENT,
		to:           child,
	})
	r.relations = append(r.relations, RelativeInfo{
		from:         child,
		relationship: CHILD,
		to:           parent,
	})
}

// Now the implementation belongs to the low level module
func (r *Relationships) FindAllChildrenOfParent(pName string) []*Person {
	result := make([]*Person, 0)

	for i, rel := range r.relations {
		if rel.from.name == pName && rel.relationship == PARENT {
			result = append(result, r.relations[i].to)
		}
	}
	return result
}

// High Level Module
type Research struct {
	// fix DIP by using Abstraction called Relationship Browser (Interface)
	browser RelationshipBrowser
}

func (r *Research) GetChildrenOfParent(pName string) {
	// The HLM uses functions exposed by interface without knowing the internal implementation
	for _, p := range r.browser.FindAllChildrenOfParent(pName) {
		fmt.Printf("%s is a child of %s\n", p.name, pName)
	}
}

func main() {

	parent1 := &Person{name: "John"}
	child1 := &Person{name: "Ben"}
	child2 := &Person{name: "Matt"}

	relationships := &Relationships{}
	relationships.AddParentAndChildRelation(parent1, child1)
	relationships.AddParentAndChildRelation(parent1, child2)

	research := &Research{
		browser: relationships,
	}
	research.GetChildrenOfParent(parent1.name)
}
