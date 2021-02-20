package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/2.2/conv"
	"gopl.io/ch2/tempconv"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			printValues(arg)
		}
	} else {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			printValues(input.Text())
		}
	}
}

func printValues(arg string) {
	t, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}

	var (
		// temperature
		f = tempconv.Fahrenheit(t)
		c = tempconv.Celsius(t)
		// length
		m  = conv.Meter(t)
		ft = conv.Feet(t)
		// weight
		k = conv.Kilogram(t)
		p = conv.Pond(t)
	)
	print(f, tempconv.FToC(f), c, tempconv.CToF(c))
	print(m, conv.MToF(m), ft, conv.FToM(ft))
	print(k, conv.KToP(k), p, conv.PToK(p))
}

func print(a, b, c, d interface{}) {
	fmt.Printf("%s = %s, %s = %s\n", a, b, c, d)
}
