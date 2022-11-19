package main

import (
	"fmt"

	"github.com/virb30/curso-go/05-packaging/3/math"
)

func main() {
	m := math.NewMath(1, 3)
	fmt.Println(m.Add())
}
