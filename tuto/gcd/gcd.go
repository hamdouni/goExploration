// Copyright 2014 Brahim HAMDOUNI. All rights reserved.
// Use of this source code is governed by SIT license
// that can be found in the SIT LICENSE file

// gcd gives the greater common divisor of two integers (euclid algorithm)
package main

import (
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 3 {
		println("usage: gcd integer integer")
		os.Exit(0)
	}

	A, err := strconv.Atoi(os.Args[1])
	if err != nil {
		println("first argument is not an integer")
		os.Exit(1)
	}
	B, err := strconv.Atoi(os.Args[2])
	if err != nil {
		println("second argument is not an integer")
		os.Exit(1)
	}

	R := 1
	P := A
	Q := B

	for R != 0 {
		R = P % Q
		P = Q
		Q = R
	}
	println(P)
}
