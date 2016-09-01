package morelia

import (
	"io"
)


type Object interface {
	Print(w io.Writer) error
	HasAttr(name Object) bool
	HasAttrString(name string) bool
	GetAttr(name Object) (Object, error)
	GetAttrString(name string) (Object, error)
	SetAttr(name, value Object) error
	SetAttrString(name string, value Object) error
	DelAttr(name Object) error
	DelAttrString(name string) error
	RichCompare(o2 Object, op Op) (Object, error)
	RichCompareBool(o2 Object, op Op) (bool, error)
	Repr() (Object, error)
	ASCII() (Object, error)
	Str() (Object, error)
	Bytes() (Object, error)
	Hash() (int, error)
	IsTrue() (bool, error)
	Not() (bool, error)
	Type() *Type
	Length() (int, error)
	Size() (int, error)
	LengthHint() (int, error)
	GetItem(key Object) (Object, error)
	SetItem(key, value Object) error
	DelItem(key Object) error
	Dir() (Object, error)
	GetIter() (Object, error)
}

type Callable interface {
	Call(args, kwargs Object) (Object, error)
	CallObject(args Object) (Object, error)
	//CallFunction(fmt string, args... interface{})
	//CallMethod(method, fmt string, args... interface{})
	CallFunctionArgs(args... Object) (Object, error)
	CallMethodArgs(name Object, args... Object) (Object, error)
}

type Number interface {
	Add(o2 Object) (Object, error)
	Subtract(o2 Object) (Object, error)
	Multiply(o2 Object) (Object, error)
	MatrixMultiply(o2 Object) (Object, error)
	FloorDivide(o2 Object) (Object, error)
	TrueDivide(o2 Object) (Object, error)
	Remainder(o2 Object) (Object, error)
	Divmod(o2 Object) (Object, error)
	Power(o2, o3 Object) (Object, error)
	Negative() (Object, error)
	Positive() (Object, error)
	Absolute() (Object, error)
	Invert() (Object, error)
	Lshift(o2 Object) (Object, error)
	Rshift(o2 Object) (Object, error)
	And(o2 Object) (Object, error)
	Xor(o2 Object) (Object, error)
	Or(o2 Object) (Object, error)
	InPlaceAdd(o2 Object) (Object, error)
	InPlaceSubtract(o2 Object) (Object, error)
	InPlaceMultiply(o2 Object) (Object, error)
	InPlaceMatrixMultipl(o2 Object) (Object, error)
	InPlaceFloorDivide(o2 Object) (Object, error)
	InPlaceTrueDivide(o2 Object) (Object, error)
	InPlaceRemainder(o2 Object) (Object, error)
	InPlacePower(o2, o3 Object) (Object, error)
	InPlaceLshift(o2 Object) (Object, error)
	InPlaceRshift(o2 Object) (Object, error)
	InPlaceAnd(o2 Object) (Object, error)
	InPlaceXor(o2 Object) (Object, error)
	Long() (Object, error)
	Float() (Object, error)
	Index() (Object, error)
	ToBase(base int) (Object, error)
	AsInt() (int, error)
}

type Sequence interface {
	Size() (int, error)
	Length() (int, error)
	Concat(o2 Object) (Object, error)
	Repeat(count int) (Object, error)
	InPlaceConcat(o2 Object) (Object, error)
	InPlaceRepeat(count int) (Object, error)
	GetItemInt(i int) (Object, error)
	GetSliceInt(i1, i2 int) (Object, error)
	SetItemInt(i int, v Object) error
	DelItemInt(i int) error
	SetSliceInt(i1, i2 int, v Object) error
	DelSliceInt(i1, i2 int) error
	Count(value Object) (int, error)
	Contains(value Object) (bool, error)
	Index(value Object) (int, error)
	List() (Object, error)
	Tuple() (Object, error)
}

type Mapping interface {
	Size() (int, error)
	Length() (int, error)
	DelItemString(key string) error
	DelItem(key Object) error
	HasKeyString(key string) bool
	HasKey(key Object) bool
	Keys() (Object, error)
	Values() (Object, error)
	Items() (Object, error)
	GetItemString(key string) (Object, error)
	SetItemString(key string, v Object) error
}

type Iterator interface {
	Next() (Object, error)
}

type Exception interface {
	error
	GetTraceback() (Object, error)
	SetTraceback(tb Object) error
	GetContext() (Object, error)
	SetContext(ctx Object) error
	GetCause() (Object, error)
	SetCause(cause Object) error
}

/*
// Don't know if I'm going to need this in Go...
type Buffer interface {

}
*/
