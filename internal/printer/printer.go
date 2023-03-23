package printer

import "fmt"

type Printer struct {}

func (p Printer) Print(s string) {
    fmt.Println(s)
}

func (p Printer) Printf(format string, a ...interface{}) {
    fmt.Printf(format, a...)
}