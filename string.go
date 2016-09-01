package morelia

import (
	"io"
)

type String string


func NewStringObject(s string) String {
	return String(s)
}


// Object protocol:

func (s String) Print(w io.Writer) error {
	_, err := w.Write([]byte(s.String()))
	return err
}

func (s String) HasAttr(name Object) bool {
	panic("not implemented")
}

func (s String) HasAttrString(name string) bool {
	panic("not implemented")
}

func (s String) GetAttr(name Object) (Object, error) {
	panic("not implemented")
}

func (s String) GetAttrString(name string) (Object, error) {
	return GetAttrString(s, name)
}

func (s String) SetAttr(name Object, value Object) error {
	panic("not implemented")
}

func (s String) SetAttrString(name string, value Object) error {
	return SetAttrString(s, name, value)
}

func (s String) DelAttr(name Object) error {
	panic("not implemented")
}

func (s String) DelAttrString(name string) error {
	return DelAttrString(s, name)
}

func (s String) RichCompare(o2 Object, op Op) (Object, error) {
	if s2, ok := o2.(String); ok {
		str1 := []byte(s.String())
		str2 := []byte(s2.String())
		l1 := len(str1)
		l2 := len(str2)
		if l1 != l2 {
			return False, nil
		}
		for i := 0; i < l1 && i < l2; i++ {
			if str1[i] != str2[i] {
				return False, nil
			}
		}
		return True, nil
	} else {
		return nil, NotImplemented
	}
}

func (s String) RichCompareBool(o2 Object, op Op) (bool, error) {
	panic("not implemented")
}

func (s String) Repr() (Object, error) {
	panic("not implemented")
}

func (s String) ASCII() (Object, error) {
	panic("not implemented")
}

func (s String) Str() (Object, error) {
	return s, nil
}

func (s String) Bytes() (Object, error) {
	panic("not implemented")
}

func (s String) Hash() (int, error) {
	panic("not implemented")
}

func (s String) IsTrue() (bool, error) {
	panic("not implemented")
}

func (s String) Not() (bool, error) {
	panic("not implemented")
}

func (s String) Type() *Type {
	panic("not implemented")
}

func (s String) Length() (int, error) {
	panic("not implemented")
}

func (s String) Size() (int, error) {
	panic("not implemented")
}

func (s String) LengthHint() (int, error) {
	panic("not implemented")
}

func (s String) GetItem(key Object) (Object, error) {
	panic("not implemented")
}

func (s String) SetItem(key Object, value Object) error {
	panic("not implemented")
}

func (s String) DelItem(key Object) error {
	panic("not implemented")
}

func (s String) Dir() (Object, error) {
	panic("not implemented")
}

func (s String) GetIter() (Object, error) {
	panic("not implemented")
}


// Sequence protocol:

/*
func (s String) Size() (int, error) {
	panic("not implemented")
}

func (s String) Length() (int, error) {
	panic("not implemented")
}
*/

func (s String) Concat(o2 Object) (Object, error) {
	panic("not implemented")
}

func (s String) Repeat(count int) (Object, error) {
	panic("not implemented")
}

func (s String) InPlaceConcat(o2 Object) (Object, error) {
	panic("not implemented")
}

func (s String) InPlaceRepeat(count int) (Object, error) {
	panic("not implemented")
}

func (s String) GetItemInt(i int) (Object, error) {
	panic("not implemented")
}

func (s String) GetSliceInt(i1 int, i2 int) (Object, error) {
	panic("not implemented")
}

func (s String) SetItemInt(i int, v Object) error {
	panic("not implemented")
}

func (s String) DelItemInt(i int) error {
	panic("not implemented")
}

func (s String) SetSliceInt(i1 int, i2 int, v Object) error {
	panic("not implemented")
}

func (s String) DelSliceInt(i1 int, i2 int) error {
	panic("not implemented")
}

func (s String) Count(value Object) (int, error) {
	panic("not implemented")
}

func (s String) Contains(value Object) (bool, error) {
	panic("not implemented")
}

func (s String) Index(value Object) (int, error) {
	panic("not implemented")
}

func (s String) List() (Object, error) {
	panic("not implemented")
}

func (s String) Tuple() (Object, error) {
	panic("not implemented")
}

func (s String) String() string {
	return string(s)
}