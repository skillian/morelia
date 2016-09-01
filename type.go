package morelia

import (
	"io"
	"reflect"
)


type Type struct {
	GoType reflect.Type
	Base *Type
	Mro Object
	Name String
}


func (t *Type) Print(w io.Writer) error {
	panic("not implemented")
}

func (t *Type) HasAttr(name Object) bool {
	panic("not implemented")
}

func (t *Type) HasAttrString(name string) bool {
	panic("not implemented")
}

func (t *Type) GetAttr(name Object) (Object, error) {
	panic("not implemented")
}

func (t *Type) GetAttrString(name string) (Object, error) {
	panic("not implemented")
}

func (t *Type) SetAttr(name Object, value Object) error {
	panic("not implemented")
}

func (t *Type) SetAttrString(name string, value Object) error {
	panic("not implemented")
}

func (t *Type) DelAttr(name Object) error {
	panic("not implemented")
}

func (t *Type) DelAttrString(name string) error {
	panic("not implemented")
}

func (t *Type) RichCompare(o2 Object, op Op) (Object, error) {
	panic("not implemented")
}

func (t *Type) RichCompareBool(o2 Object, op Op) (bool, error) {
	panic("not implemented")
}

func (t *Type) Repr() (Object, error) {
	panic("not implemented")
}

func (t *Type) ASCII() (Object, error) {
	panic("not implemented")
}

func (t *Type) Str() (Object, error) {
	panic("not implemented")
}

func (t *Type) Bytes() (Object, error) {
	panic("not implemented")
}

func (t *Type) Hash() (int, error) {
	panic("not implemented")
}

func (t *Type) IsTrue() (bool, error) {
	panic("not implemented")
}

func (t *Type) Not() (bool, error) {
	panic("not implemented")
}

func (t *Type) Type() *Type {
	panic("not implemented")
}

func (t *Type) Length() (int, error) {
	panic("not implemented")
}

func (t *Type) Size() (int, error) {
	panic("not implemented")
}

func (t *Type) LengthHint() (int, error) {
	panic("not implemented")
}

func (t *Type) GetItem(key Object) (Object, error) {
	panic("not implemented")
}

func (t *Type) SetItem(key Object, value Object) error {
	panic("not implemented")
}

func (t *Type) DelItem(key Object) error {
	panic("not implemented")
}

func (t *Type) Dir() (Object, error) {
	panic("not implemented")
}

func (t *Type) GetIter() (Object, error) {
	panic("not implemented")
}

