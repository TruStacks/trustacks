package engine

type Fact int

const NilFact Fact = -1

var factInc = 0

func NewFact() Fact {
	factInc++
	return Fact(factInc)
}
