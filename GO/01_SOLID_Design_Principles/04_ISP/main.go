package main

import "fmt"

type Document struct {
	text string
}

type MachineOperations int

const (
	PRINT MachineOperations = iota
	SCAN
	FAX
)

type Printer interface {
	print(d Document)
}

type Scanner interface {
	scan(d Document)
}

type MultiFunctinalMachine interface {
	Printer
	Scanner
	fax(d Document)
}

type ModernPrinterMachine struct{}

func (mpm *ModernPrinterMachine) print(d Document) {
	fmt.Printf("Printing... %s\n", d.text)
}
func (mpm *ModernPrinterMachine) scan(d Document) {
	fmt.Printf("Scanning... %s\n", d.text)
}
func (mpm *ModernPrinterMachine) fax(d Document) {
	fmt.Printf("Faxing... %s\n", d.text)
}

type OldPrinterMachine struct{}

func (opm *OldPrinterMachine) print(d Document) {
	fmt.Printf("Printing in black and white... %s\n", d.text)
}

func Operate(mp MultiFunctinalMachine, operation MachineOperations, d Document) {
	switch operation {
	case PRINT:
		mp.print(d)
	case SCAN:
		mp.scan(d)
	case FAX:
		mp.fax(d)
	}
}

func OperatePrinter(p Printer, d Document) {
	p.print(d)
}

func main() {
	doc := Document{text: "Hello world"}

	mpm := &ModernPrinterMachine{}

	Operate(mpm, PRINT, doc)
	Operate(mpm, SCAN, doc)
	Operate(mpm, FAX, doc)

	opm := &OldPrinterMachine{}

	OperatePrinter(mpm, doc)
	OperatePrinter(opm, doc)

}
