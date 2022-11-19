package math

type math struct {
	a int
	b int
	C int
}

func NewMath(a, b int) math {
	return math{a: a, b: b}
}

func (m *math) Add() int {
	m.C = m.a + m.b
	return m.C
}
