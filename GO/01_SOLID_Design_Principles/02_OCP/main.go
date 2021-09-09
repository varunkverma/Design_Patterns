package main

import "fmt"

// OCP
// Enterprise pattern - Specification

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

// Using Specification Enterprise pattern

// Specification Interface
type Specification interface {
	isSatisfied(p *Product) bool
}

// specific implementations of the Specification Interface
type ColorSpecification struct {
	color Color
}

func (cs *ColorSpecification) isSatisfied(p *Product) bool {
	return p.color == cs.color
}

type SizeSpecification struct {
	size Size
}

func (ss *SizeSpecification) isSatisfied(p *Product) bool {
	return p.size == ss.size
}

// Combinator Specification - Combines 2 Specifications
type AndSpecification struct {
	spec1 Specification
	spec2 Specification
}

func (andSpec *AndSpecification) isSatisfied(p *Product) bool {
	return andSpec.spec1.isSatisfied(p) && andSpec.spec2.isSatisfied(p)
}

// Combinator Specification - Combines multiple Specifications
type MultiAndSpecification struct {
	specs []Specification
}

func (mASpec *MultiAndSpecification) isSatisfied(p *Product) bool {
	isSatisfiedByAllSpecifications := true

	for _, spec := range mASpec.specs {
		isSatisfiedByAllSpecifications = isSatisfiedByAllSpecifications && spec.isSatisfied(p)
		if !isSatisfiedByAllSpecifications {
			break
		}
	}

	return isSatisfiedByAllSpecifications
}

// type that is not to be modified
type Filter struct{}

func (f *Filter) filterProducts(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)

	for i, p := range products {
		if spec.isSatisfied(&p) {
			result = append(result, &products[i])
		}
	}

	return result
}

func main() {
	apple := Product{
		name:  "Apple",
		color: green,
		size:  small,
	}
	tree := Product{
		name:  "Tree",
		color: green,
		size:  large,
	}
	house := Product{
		name:  "house",
		color: blue,
		size:  large,
	}

	products := []Product{
		apple,
		house,
		tree,
	}

	fil := &Filter{}

	greenColorSpec := &ColorSpecification{color: green}
	fmt.Println("Filtered green color products are:")
	for _, p := range fil.filterProducts(products, greenColorSpec) {
		fmt.Printf("- %s\n", p.name)
	}

	largeSizeSpec := &SizeSpecification{size: large}

	greenAndLageSpec := &AndSpecification{greenColorSpec, largeSizeSpec}
	fmt.Println("Filtered green colored and large sized products are:")
	for _, p := range fil.filterProducts(products, greenAndLageSpec) {
		fmt.Printf("- %s\n", p.name)
	}

	greenAndLageSpec2 := &MultiAndSpecification{
		specs: []Specification{greenColorSpec, largeSizeSpec},
	}
	fmt.Println("Filtered green colored and large sized products are:")
	for _, p := range fil.filterProducts(products, greenAndLageSpec2) {
		fmt.Printf("- %s\n", p.name)
	}
}
