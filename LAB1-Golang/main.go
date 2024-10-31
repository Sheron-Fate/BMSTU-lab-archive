package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func getValues() (x, y, z float64) {
	fmt.Println("Main: Trying to get from console...")
	consoleArguments := os.Args[1:]
	if len(consoleArguments) == 0 {
		fmt.Println("Main: Failed to get values\nMain:  Please enter arguments: ")
		fmt.Scanln(&x, &y, &z)
	} else {
		x, _ = strconv.ParseFloat(consoleArguments[0], 64)
		y, _ = strconv.ParseFloat(consoleArguments[1], 64)
		z, _ = strconv.ParseFloat(consoleArguments[2], 64)
	}
	return
}

func getDiscriminant(a, b, c float64) float64 {
	return b*b - 4*a*c
}

func getQuadraticRoots(a, b, d float64) (qRoots []float64) {
	if d == 0 {
		qRoots = append(qRoots, (-1)*b/2*a)
	} else if d > 0 {
		qRoots = append(qRoots, ((-1)*b+math.Sqrt(d))/2*a, ((-1)*b-math.Sqrt(d))/2*a)
	}
	return
}

func getRoots(qRoots []float64) (roots []float64) {
	for _, root := range qRoots {
		if root > 0 {
			roots = append(roots, math.Sqrt(root), (-1)*math.Sqrt(root))
		}
	}
	return
}

func main() {
	a, b, c := getValues()
	fmt.Println(a, b, c)

	qRoots := getQuadraticRoots(a, b, getDiscriminant(a, b, c))
	fmt.Println(qRoots)

	roots := getRoots(qRoots)
	fmt.Println(roots)

}
