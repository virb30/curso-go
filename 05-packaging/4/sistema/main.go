package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/virb30/curso-go/05-packaging/4/math"
)

func main() {
	m := math.NewMath(1, 3)
	fmt.Println(m.Add())
	fmt.Println(uuid.New().String())
}
