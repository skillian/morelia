package morelia

import (
	"io"
)


type NotImplementedType struct {}


var NotImplemented = NotImplementedType{}


// Object protocol:

func (ni NotImplementedType) Print(w io.Writer) error {
	panic("not implemented")
}

func (ni NotImplementedType) HasAttr(name Object) bool {
	panic("not implemented")
}

func (ni NotImplementedType) HasAttrString(name string) bool {
	panic("not implemented")
}

func (ni NotImplementedType) GetAttr(name Object) (Object, error) {
	panic("not implemented")
}

func (ni NotImplementedType) GetAttrString(name string) (Object, error) {
	panic("not implemented")
}

func (ni NotImplementedType) SetAttr(name Object, value Object) error {
	panic("not implemented")
}

func (ni NotImplementedType) SetAttrString(name string, value Object) error {
	panic("not implemented")
}

func (ni NotImplementedType) DelAttr(name Object) error {
	panic("not implemented")
}

func (ni NotImplementedType) DelAttrString(name string) error {
	panic("not implemented")
}

func (ni NotImplementedType) RichCompare(o2 Object, op Op) (Object, error) {
	panic("not implemented")
}

func (ni NotImplementedType) RichCompareBool(o2 Object, op Op) (bool, error) {
	panic("not implemented")
}

func (ni NotImplementedType) Repr() (Object, error) {
	panic("not implemented")
}

func (ni NotImplementedType) ASCII() (Object, error) {
	panic("not implemented")
}

func (ni NotImplementedType) Str() (Object, error) {
	panic("not implemented")
}

func (ni NotImplementedType) Bytes() (Object, error) {
	panic("not implemented")
}

func (ni NotImplementedType) Hash() (int, error) {
	panic("not implemented")
}

func (ni NotImplementedType) IsTrue() (bool, error) {
	panic("not implemented")
}

func (ni NotImplementedType) Not() (bool, error) {
	panic("not implemented")
}

func (ni NotImplementedType) Type() (Object, error) {
	panic("not implemented")
}

func (ni NotImplementedType) Length() (int, error) {
	panic("not implemented")
}

func (ni NotImplementedType) Size() (int, error) {
	panic("not implemented")
}

func (ni NotImplementedType) LengthHint() (int, error) {
	panic("not implemented")
}

func (ni NotImplementedType) GetItem(key Object) (Object, error) {
	panic("not implemented")
}

func (ni NotImplementedType) SetItem(key Object, value Object) error {
	panic("not implemented")
}

func (ni NotImplementedType) DelItem(key Object) error {
	panic("not implemented")
}

func (ni NotImplementedType) Dir() (Object, error) {
	panic("not implemented")
}

func (ni NotImplementedType) GetIter() (Object, error) {
	panic("not implemented")
}


// Exception protocol:

func (ni NotImplementedType) Error() string {
	return "NotImplemented"
}

func (ni NotImplementedType) GetTraceback() (Object, error) {
	panic("not implemented")
}

func (ni NotImplementedType) SetTraceback(tb Object) error {
	panic("not implemented")
}

func (ni NotImplementedType) GetContext() (Object, error) {
	panic("not implemented")
}

func (ni NotImplementedType) SetContext(ctx Object) error {
	panic("not implemented")
}

func (ni NotImplementedType) GetCause() (Object, error) {
	panic("not implemented")
}

func (ni NotImplementedType) SetCause(cause Object) error {
	panic("not implemented")
}

