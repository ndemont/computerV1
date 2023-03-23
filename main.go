package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"github.com/davecgh/go-spew/spew"
)

type Equation struct {
	orders				[]float64
	discriminant		float64
	solutions			interface{}
}

func findRealRoots(equ *Equation) {
	
	var solutions [2]float64

	solutions[0] = ((-equ.orders[1]) - (math.Sqrt(equ.discriminant))) / (2 * equ.orders[2])
	solutions[1] = ((-equ.orders[1]) + (math.Sqrt(equ.discriminant))) / (2 * equ.orders[2])

	equ.solutions = solutions
}

func findComplexRoots(equ *Equation) {
	equ.complexSolutions[0] = complex(
		(-equ.orders[1])/(2*equ.),
		(math.Sqrt(-equ.discriminant))/(2*equ.))
	equ.complexSolutions[1] = complex(
		(-equ.orders[1])/(2*equ.),
		-((math.Sqrt(-equ.discriminant)) / (2 * equ.)))
}

func solveEquation(equ *Equation) {

	equ.discriminant = math.Pow(equ.b, 2) - (4 * equ.a * equ.c)

	if equ.discriminant >= 0 {
		findRealRoots(equ)
	} else if equ.discriminant < 0 {
		findComplexRoots(equ)
	}
}

func extractEquation(equ *Equation, elem []string, sign float64, side float64) (err error) {

	var orders string = "012"

	spew.Dump(sign)
	spew.Dump(side)
	spew.Dump(elem)

	// Check Full Format
	if len(elem) != 3 {
		return fmt.Errorf("%v: wrong number of elements", elem)
	}

	// Check first elem
	nb, err := strconv.ParseFloat(elem[0], 64)

	if err != nil {
		return fmt.Errorf("%v: not a float", elem[0])
	}

	// Check second elem
	if elem[1] != "*" {
		return fmt.Errorf("%v: wrong operation, not *", elem[1])
	}

	// Check third elem
	if len(elem[2]) != 3 || elem[2][:2] != "X^" || !strings.Contains(orders, string(elem[2][2])) {
		return fmt.Errorf("%v: wrong format of element, not X^", elem[2])
	}

	// Increase power argumnent value
	if elem[2][2] == '0' {
		equ.c = equ.c + (side * sign * nb)
	}
	if elem[2][2] == '1' {
		equ.b = equ.b + (side * sign * nb)
	}
	if elem[2][2] == '2' {
		equ.a = equ.a + (side * sign * nb)
	}

	if (equ.a != 0) {
		equ.degree = 2
	} else if (equ.b != 0) {
		equ.degree = 1
	} else if (equ.c != 0) {
		equ.degree = 0
	} else {
		equ.degree = -1
	}

	return nil
}

func checkFormat(split []string) (equ Equation, err error) {

	var sep string = "+-="
	var side float64 = 1
	var sign float64 = 1
	var start int = 0
	var end int = 1

	for index, elem := range split {
		if strings.Contains(sep, elem) {

			end = index

			spew.Dump(start)
			spew.Dump(end)

			if end > start {
				err = extractEquation(&equ, split[start:end], sign, side)
				if err != nil {
					return equ, err
				}
			}

			if elem == "+" {
				sign = 1
			} else if elem == "-" {
				sign = -1
			} else {
				side = -1
			}
			start = index + 1
		}
	}

	err = extractEquation(&equ, split[start:], sign, side)

	return equ, err
}

func printReducedForm(equ Equation) {
	println("Reduced form: ")
	println(equ.c, " * X^0 ")
	println(equ.b, " * X^1 ")
	println(equ.a, " * X^2 ")
	println("= 0")

}

func main() {
	equation := flag.String("equation", "", "2nd degrees equation to solve.")
	flag.Parse()

	if *equation == "" {
		flag.Usage()
		os.Exit(1)
	}

	split := strings.Fields(*equation)
	equ, err := checkFormat(split)

	if err != nil {
		os.Exit(1)
	}

	printReducedForm(equ)
	solveEquation(&equ)

	spew.Dump(equ)

	os.Exit(0)
}
