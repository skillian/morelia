package morelia


import (
	"fmt"
)


type Op int

const (
	LT Op = iota
	LE
	EQ
	NE
	GT
	GE
)

var (
	SwappedOp = [...]Op{GT, GE, EQ, NE, LT, LE}
)

func (op Op) String() string {
	switch op {
	case LT:
		return "<"
	case LE:
		return "<="
	case EQ:
		return "=="
	case NE:
		return "!="
	case GT:
		return ">"
	case GE:
		return ">="
	default:
		panic(fmt.Sprintf("Undefined operator: %v", op))
	}
}
