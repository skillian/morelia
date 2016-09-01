package morelia


import (
	"errors"
	"fmt"
	"io"
	"reflect"
)


const (
	ZeroInt = Int(0)
)

var (
	IntType = Type{GoType: reflect.TypeOf(ZeroInt), Base: nil, Mro: nil, Name: "int"}

	ZeroDivisionError = errors.New("division by zero")
)


type Int int


func NewIntObject(i int) Int {
	return Int(i)
}

func (i Int) Print(w io.Writer) error {
	panic("not implemented")
}

func (i Int) HasAttr(name Object) bool {
	panic("not implemented")
}

func (i Int) HasAttrString(name string) bool {
	return HasAttrString(i, name)
}

func (i Int) GetAttr(name Object) (Object, error) {
	panic("not implemented")
}

func (i Int) GetAttrString(name string) (Object, error) {
	return GetAttrString(i, name)
}

func (i Int) SetAttr(name Object, value Object) error {
	panic("not implemented")
}

func (i Int) SetAttrString(name string, value Object) error {
	return SetAttrString(i, name, value)
}

func (i Int) DelAttr(name Object) error {
	panic("not implemented")
}

func (i Int) DelAttrString(name string) error {
	return DelAttrString(i, name)
}

// Simple Go int comparison, returns Go bool and error if op is invalid
func poorCompare(i1, i2 int, op Op) (bool, error) {
	switch op {
	case EQ:
		return i1 == i2, nil
	case NE:
		return i1 != i2, nil
	case LE:
		return i1 <= i2, nil
	case GE:
		return i1 >= i2, nil
	case LT:
		return i1 < i2, nil
	case GT:
		return i1 > i2, nil
	default:
		return false, fmt.Errorf("Invalid operator:", op)
	}
}

func (i Int) RichCompare(o2 Object, op Op) (Object, error) {
	switch x := o2.(type) {
	case Int:
		i1, _ := i.AsInt()
		i2, _ := x.AsInt()
		if cmp, err := poorCompare(i1, i2, op); err == nil {
			return NewBoolObject(cmp), nil
		} else {
			return nil, err
		}

	default:
		return nil, NotImplemented
	}
}

func (i Int) RichCompareBool(o2 Object, op Op) (bool, error) {
	return RichCompareBool(i, o2, op)
}

func (i Int) Repr() (Object, error) {
	panic("not implemented")
}

func (i Int) ASCII() (Object, error) {
	panic("not implemented")
}

func (i Int) Str() (Object, error) {
	panic("not implemented")
}

func (i Int) Bytes() (Object, error) {
	panic("not implemented")
}

func (i Int) Hash() (int, error) {
	panic("not implemented")
}

func (i Int) IsTrue() (bool, error) {
	panic("not implemented")
}

func (i Int) Not() (bool, error) {
	panic("not implemented")
}

func (i Int) Type() *Type {
	return &IntType
}

func (i Int) Length() (int, error) {
	panic("not implemented")
}

func (i Int) Size() (int, error) {
	panic("not implemented")
}

func (i Int) LengthHint() (int, error) {
	panic("not implemented")
}

func (i Int) GetItem(key Object) (Object, error) {
	panic("not implemented")
}

func (i Int) SetItem(key Object, value Object) error {
	panic("not implemented")
}

func (i Int) DelItem(key Object) error {
	panic("not implemented")
}

func (i Int) Dir() (Object, error) {
	panic("not implemented")
}

func (i Int) GetIter() (Object, error) {
	panic("not implemented")
}


// Number protocol:

func (i Int) Add(o2 Object) (Object, error) {
	if i2, ok := o2.(Int); ok {
		return NewIntObject(int(i) + int(i2)), nil
	} else {
		return nil, NotImplemented
	}
}

func (i Int) Subtract(o2 Object) (Object, error) {
	if i2, ok := o2.(Int); ok {
		return NewIntObject(int(i) - int(i2)), nil
	} else {
		return nil, NotImplemented
	}
}

func (i Int) Multiply(o2 Object) (Object, error) {
	if i2, ok := o2.(Int); ok {
		return NewIntObject(int(i) * int(i2)), nil
	} else {
		return nil, NotImplemented
	}
}

func (i Int) MatrixMultiply(o2 Object) (Object, error) {
	panic("not implemented")
}

func (i Int) FloorDivide(o2 Object) (Object, error) {
	if i2, ok := o2.(Int); ok {
		int2 := int(i2)
		if int2 == 0 {
			return nil, ZeroDivisionError
		}
		return NewIntObject(int(i) / int2), nil
	} else {
		return nil, NotImplemented
	}
}

func (i Int) TrueDivide(o2 Object) (Object, error) {
	panic("not implemented")
}

func (i Int) Remainder(o2 Object) (Object, error) {
	panic("not implemented")
}

func (i Int) Divmod(o2 Object) (Object, error) {
	panic("not implemented")
}

func (i Int) Power(o2 Object, o3 Object) (Object, error) {
	panic("not implemented")
}

func (i Int) Negative() (Object, error) {
	panic("not implemented")
}

func (i Int) Positive() (Object, error) {
	panic("not implemented")
}

func (i Int) Absolute() (Object, error) {
	panic("not implemented")
}

func (i Int) Invert() (Object, error) {
	panic("not implemented")
}

func (i Int) Lshift(o2 Object) (Object, error) {
	panic("not implemented")
}

func (i Int) Rshift(o2 Object) (Object, error) {
	panic("not implemented")
}

func (i Int) And(o2 Object) (Object, error) {
	panic("not implemented")
}

func (i Int) Xor(o2 Object) (Object, error) {
	panic("not implemented")
}

func (i Int) Or(o2 Object) (Object, error) {
	panic("not implemented")
}

func (i Int) InPlaceAdd(o2 Object) (Object, error) {
	panic("not implemented")
}

func (i Int) InPlaceSubtract(o2 Object) (Object, error) {
	panic("not implemented")
}

func (i Int) InPlaceMultiply(o2 Object) (Object, error) {
	panic("not implemented")
}

func (i Int) InPlaceMatrixMultipl(o2 Object) (Object, error) {
	panic("not implemented")
}

func (i Int) InPlaceFloorDivide(o2 Object) (Object, error) {
	panic("not implemented")
}

func (i Int) InPlaceTrueDivide(o2 Object) (Object, error) {
	panic("not implemented")
}

func (i Int) InPlaceRemainder(o2 Object) (Object, error) {
	panic("not implemented")
}

func (i Int) InPlacePower(o2 Object, o3 Object) (Object, error) {
	panic("not implemented")
}

func (i Int) InPlaceLshift(o2 Object) (Object, error) {
	panic("not implemented")
}

func (i Int) InPlaceRshift(o2 Object) (Object, error) {
	panic("not implemented")
}

func (i Int) InPlaceAnd(o2 Object) (Object, error) {
	panic("not implemented")
}

func (i Int) InPlaceXor(o2 Object) (Object, error) {
	panic("not implemented")
}

func (i Int) Long() (Object, error) {
	panic("not implemented")
}

func (i Int) Float() (Object, error) {
	panic("not implemented")
}

func (i Int) Index() (Object, error) {
	panic("not implemented")
}

func (i Int) ToBase(base int) (Object, error) {
	panic("not implemented")
}

func (i Int) AsInt() (int, error) {
	return int(i), nil
}

