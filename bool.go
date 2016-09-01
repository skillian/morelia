package morelia


import (
	"io"
)


type Bool bool


const (
	True = Bool(true)
	False = Bool(false)
)


func NewBoolObject(b bool) Bool {
	return Bool(b)
}


// Object protocol:

func (b Bool) Print(w io.Writer) error {
	panic("not implemented")
}

func (b Bool) HasAttr(name Object) bool {
	panic("not implemented")
}

func (b Bool) HasAttrString(name string) bool {
	panic("not implemented")
}

func (b Bool) GetAttr(name Object) (Object, error) {
	panic("not implemented")
}

func (b Bool) GetAttrString(name string) (Object, error) {
	panic("not implemented")
}

func (b Bool) SetAttr(name Object, value Object) error {
	panic("not implemented")
}

func (b Bool) SetAttrString(name string, value Object) error {
	panic("not implemented")
}

func (b Bool) DelAttr(name Object) error {
	panic("not implemented")
}

func (b Bool) DelAttrString(name string) error {
	panic("not implemented")
}

func (b Bool) RichCompare(o2 Object, op Op) (Object, error) {
	panic("not implemented")
}

func (b Bool) RichCompareBool(o2 Object, op Op) (bool, error) {
	panic("not implemented")
}

func (b Bool) Repr() (Object, error) {
	panic("not implemented")
}

func (b Bool) ASCII() (Object, error) {
	panic("not implemented")
}

func (b Bool) Str() (Object, error) {
	panic("not implemented")
}

func (b Bool) Bytes() (Object, error) {
	panic("not implemented")
}

func (b Bool) Hash() (int, error) {
	panic("not implemented")
}

func (b Bool) IsTrue() (bool, error) {
	return b == True, nil
}

func (b Bool) Not() (bool, error) {
	panic("not implemented")
}

func (b Bool) Type() *Type {
	panic("not implemented")
}

func (b Bool) Length() (int, error) {
	panic("not implemented")
}

func (b Bool) Size() (int, error) {
	panic("not implemented")
}

func (b Bool) LengthHint() (int, error) {
	panic("not implemented")
}

func (b Bool) GetItem(key Object) (Object, error) {
	panic("not implemented")
}

func (b Bool) SetItem(key Object, value Object) error {
	panic("not implemented")
}

func (b Bool) DelItem(key Object) error {
	panic("not implemented")
}

func (b Bool) Dir() (Object, error) {
	panic("not implemented")
}

func (b Bool) GetIter() (Object, error) {
	panic("not implemented")
}


// Number protocol:

func (b Bool) Add(o2 Object) (Object, error) {
	panic("not implemented")
}

func (b Bool) Subtract(o2 Object) (Object, error) {
	panic("not implemented")
}

func (b Bool) Multiply(o2 Object) (Object, error) {
	panic("not implemented")
}

func (b Bool) MatrixMultiply(o2 Object) (Object, error) {
	panic("not implemented")
}

func (b Bool) FloorDivide(o2 Object) (Object, error) {
	panic("not implemented")
}

func (b Bool) TrueDivide(o2 Object) (Object, error) {
	panic("not implemented")
}

func (b Bool) Remainder(o2 Object) (Object, error) {
	panic("not implemented")
}

func (b Bool) Divmod(o2 Object) (Object, error) {
	panic("not implemented")
}

func (b Bool) Power(o2 Object, o3 Object) (Object, error) {
	panic("not implemented")
}

func (b Bool) Negative() (Object, error) {
	panic("not implemented")
}

func (b Bool) Positive() (Object, error) {
	panic("not implemented")
}

func (b Bool) Absolute() (Object, error) {
	panic("not implemented")
}

func (b Bool) Invert() (Object, error) {
	panic("not implemented")
}

func (b Bool) Lshift(o2 Object) (Object, error) {
	panic("not implemented")
}

func (b Bool) Rshift(o2 Object) (Object, error) {
	panic("not implemented")
}

func (b Bool) And(o2 Object) (Object, error) {
	panic("not implemented")
}

func (b Bool) Xor(o2 Object) (Object, error) {
	panic("not implemented")
}

func (b Bool) Or(o2 Object) (Object, error) {
	panic("not implemented")
}

func (b Bool) InPlaceAdd(o2 Object) (Object, error) {
	panic("not implemented")
}

func (b Bool) InPlaceSubtract(o2 Object) (Object, error) {
	panic("not implemented")
}

func (b Bool) InPlaceMultiply(o2 Object) (Object, error) {
	panic("not implemented")
}

func (b Bool) InPlaceMatrixMultipl(o2 Object) (Object, error) {
	panic("not implemented")
}

func (b Bool) InPlaceFloorDivide(o2 Object) (Object, error) {
	panic("not implemented")
}

func (b Bool) InPlaceTrueDivide(o2 Object) (Object, error) {
	panic("not implemented")
}

func (b Bool) InPlaceRemainder(o2 Object) (Object, error) {
	panic("not implemented")
}

func (b Bool) InPlacePower(o2 Object, o3 Object) (Object, error) {
	panic("not implemented")
}

func (b Bool) InPlaceLshift(o2 Object) (Object, error) {
	panic("not implemented")
}

func (b Bool) InPlaceRshift(o2 Object) (Object, error) {
	panic("not implemented")
}

func (b Bool) InPlaceAnd(o2 Object) (Object, error) {
	panic("not implemented")
}

func (b Bool) InPlaceXor(o2 Object) (Object, error) {
	panic("not implemented")
}

func (b Bool) Long() (Object, error) {
	panic("not implemented")
}

func (b Bool) Float() (Object, error) {
	panic("not implemented")
}

func (b Bool) Index() (Object, error) {
	panic("not implemented")
}

func (b Bool) ToBase(base int) (Object, error) {
	panic("not implemented")
}

func (b Bool) AsInt() (int, error) {
	if b == True {
		return 1, nil
	} else {
		return 0, nil
	}
}

