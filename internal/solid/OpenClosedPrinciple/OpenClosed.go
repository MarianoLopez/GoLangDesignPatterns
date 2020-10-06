package main

import "fmt"

type Color int //Alias

const (
	green Color = iota //enum
	blue
)

type Size int

const (
	small Size = iota
	large
)

type Product struct {
	name string
	color Color
	size Size
}

//Break the principle - for each new filter we need to implement it from scratch
/*type Filter struct {}

func (f *Filter) filterByColor(
	products []Product, color Color)[]*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}
func (f *Filter) filterBySize(
	products []Product, size Size) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}

	return result
}*/

//following the principle, for each new filter we only need to create the new specification (extends the base code)
type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

type SizeSpecification struct {
	size Size
}

type AndSpecification struct {
	first, second Specification
}

func (spec ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == spec.color
}

func (spec SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == spec.size
}

func (spec AndSpecification) IsSatisfied(p *Product) bool {
	return spec.first.IsSatisfied(p) && spec.second.IsSatisfied(p)
}

type ProductFilter struct {}

func (f *ProductFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{ "House", blue, large}

	products := []Product{apple, tree, house}

	greenSpec := ColorSpecification{green}
	largeSpec := SizeSpecification{large}
	largeGreenSpec := AndSpecification{largeSpec, greenSpec}
	productFilter := ProductFilter{}

	fmt.Print("Green products:\n")
	for _, v := range productFilter.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	fmt.Print("Large green items:\n")
	for _, v := range productFilter.Filter(products, largeGreenSpec) {
		fmt.Printf(" - %s is large and green\n", v.name)
	}
}
