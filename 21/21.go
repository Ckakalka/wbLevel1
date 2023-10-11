package main

import "fmt"

// Пусть есть классы, функционал которых мы хотим использовать,
// но они не удовлетворяют требуемому интерфейсу.
// Изменять классы по каким-либо причинам мы не можем,
// например, потому что это классы сторонних библиотек.
type Engineer struct {
}

func (*Engineer) Construct() {
	fmt.Println("Construction")
}

type Scientist struct {
}

func (*Scientist) Prove() {
	fmt.Println("Proofs")
}

// Интерфейс, которому должны удовлетворять классы
type Solver interface {
	SolveProblem()
}

// Адаптер для класса Engineer
type EngineerAdapter struct {
	*Engineer
}

func (eA EngineerAdapter) SolveProblem() {
	eA.Construct()
}

// Адаптер для класса Scientist
type ScientistAdapter struct {
	*Scientist
}

func (sA ScientistAdapter) SolveProblem() {
	sA.Prove()
}

func main() {
	scientist := Scientist{}
	engineer := Engineer{}
	solvers := []Solver{ScientistAdapter{&scientist}, EngineerAdapter{&engineer}}
	for _, solver := range solvers {
		solver.SolveProblem()
	}
}
