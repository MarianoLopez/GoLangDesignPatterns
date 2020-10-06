package main

import "fmt"

type Document struct {

}

//what if my machine can't do everything?
type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type OldFashionedPrinter struct {}

func (o OldFashionedPrinter) Print(d Document) {
	fmt.Println("OldFashionedPrinter",d)
}

func (o OldFashionedPrinter) Fax(d Document) {
	fmt.Println("OldFashionedPrinter", "operation not supported") //panic
}

// Deprecated: ...
func (o OldFashionedPrinter) Scan(d Document) {
	fmt.Println("OldFashionedPrinter", "operation not supported") //panic
}

// better approach: split into several interfaces
type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

// printer only
type MyPrinter struct {}

func (m MyPrinter) Print(d Document) {
	fmt.Println("Printer", d)
}

// combine interfaces
type Photocopier struct {}

func (p Photocopier) Scan(d Document) {
	fmt.Println("Photocopier scan: ", d)
}

func (p Photocopier) Print(d Document) {
	fmt.Println("Photocopier print: ", d)
}

type MultiFunctionDevice interface {
	Printer
	Scanner
}
// interface combination + decorator
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}

func main() {
	doc := Document{}
	old := OldFashionedPrinter{}
	old.Print(doc)
	old.Fax(doc)

	printer := MyPrinter{}
	printer.Print(doc)

	var photocopier MultiFunctionDevice = &Photocopier{}

	photocopier.Print(doc)
	photocopier.Scan(doc)

	multi := MultiFunctionMachine{photocopier, photocopier}
	multi.Print(doc)
	multi.Scan(doc)
}