package main

import (
  "fmt"
  "os"
	"github.com/ndemont/computerV1/internal/solver"

func main() {
    // ...
    // create a solver and solve the equation
    // ...

    p := printer.Printer{}

    if result.Err != nil {
        p.Printf("Error: %s\n", result.Err.Error())
        os.Exit(1)
    }

    p.Print("Results:")
    p.Printf("  - Degree: %d\n", result.Degree)
    p.Printf("  - Discriminant: %f\n", result.Discriminant)
    p.Printf("  - Roots: %v\n", result.Roots)
}
