package morelia

import (
	"errors"
	"io"
)

var (
	indexOutOfRange = errors.New("tuple index out of range")
	indexTypeError = errors.New("tuple index must be an int")
	tupleSetItemError = errors.New("tuple does not support item assignment")
	tupleDelItemError = errors.New("tuple does not support item deletion")
	tupleConcatError = errors.New("can only concatenate tuple to tuple")
)


type Tuple struct {
	items []Object
}


func NewTupleObject(items... Object) Tuple {
	return Tuple{items: items}
}

func (t Tuple) Print(w io.Writer) error {
	panic("not implemented")
}

func (t Tuple) HasAttr(name Object) bool {
	panic("not implemented")
}

func (t Tuple) HasAttrString(name string) bool {
	panic("not implemented")
}

func (t Tuple) GetAttr(name Object) (Object, error) {
	panic("not implemented")
}

func (t Tuple) GetAttrString(name string) (Object, error) {
	panic("not implemented")
}

func (t Tuple) SetAttr(name Object, value Object) error {
	panic("not implemented")
}

func (t Tuple) SetAttrString(name string, value Object) error {
	panic("not implemented")
}

func (t Tuple) DelAttr(name Object) error {
	panic("not implemented")
}

func (t Tuple) DelAttrString(name string) error {
	panic("not implemented")
}

func (t Tuple) RichCompare(o2 Object, op Op) (Object, error) {
	panic("not implemented")
}

func (t Tuple) RichCompareBool(o2 Object, op Op) (bool, error) {
	panic("not implemented")
}

func (t Tuple) Repr() (Object, error) {
	panic("not implemented")
}

func (t Tuple) ASCII() (Object, error) {
	panic("not implemented")
}

func (t Tuple) Str() (Object, error) {
	panic("not implemented")
}

func (t Tuple) Bytes() (Object, error) {
	panic("not implemented")
}

func (t Tuple) Hash() (int, error) {
	panic("not implemented")
}

func (t Tuple) IsTrue() (bool, error) {
	return len(t.items) > 0, nil
}

func (t Tuple) Not() (bool, error) {
	tr, err := t.IsTrue()
	return !tr, err
}

func (t Tuple) Type() *Type {
	panic("not implemented")
}

func (t Tuple) Length() (int, error) {
	return len(t.items), nil
}

func (t Tuple) Size() (int, error) {
	return t.Length()
}

func (t Tuple) LengthHint() (int, error) {
	return t.Length()
}

func (t Tuple) GetItem(key Object) (Object, error) {
	l := len(t.items)
	switch x := key.(type) {
	case Int:
		i, _ := x.AsInt()
		if i < 0 {
			i = l - i
		}
		if i > l {
			return nil, indexOutOfRange
		} else {
			return t.items[i], nil
		}

	default:
		return nil, indexTypeError
	}
}

func (t Tuple) SetItem(key Object, value Object) error {
	return tupleSetItemError
}

func (t Tuple) DelItem(key Object) error {
	return tupleDelItemError
}

func (t Tuple) Dir() (Object, error) {
	panic("not implemented")
}

func (t Tuple) GetIter() (Object, error) {
	panic("not implemented")
}

/*
func (t Tuple) Size() (int, error) {
	panic("not implemented")
}

func (t Tuple) Length() (int, error) {
	panic("not implemented")
}
*/

func (t Tuple) Concat(o2 Object) (Object, error) {
	if t2, ok := o2.(Tuple); ok {
		return NewTupleObject(append(t.items, t2.items...)...), nil
	} else {
		return nil, tupleConcatError
	}
}

func (t Tuple) Repeat(count int) (Object, error) {
	panic("not implemented")
}

func (t Tuple) InPlaceConcat(o2 Object) (Object, error) {
	panic("not implemented")
}

func (t Tuple) InPlaceRepeat(count int) (Object, error) {
	panic("not implemented")
}

func (t Tuple) GetItemInt(i int) (Object, error) {
	return t.GetItem(NewIntObject(i))
}

func (t Tuple) GetSliceInt(i1 int, i2 int) (Object, error) {
	return NewTupleObject(t.items[i1:i2]...), nil
}

func (t Tuple) SetItemInt(i int, v Object) error {
	return tupleSetItemError
}

func (t Tuple) DelItemInt(i int) error {
	return tupleDelItemError
}

func (t Tuple) SetSliceInt(i1 int, i2 int, v Object) error {
	return tupleSetItemError
}

func (t Tuple) DelSliceInt(i1 int, i2 int) error {
	return tupleDelItemError
}

func (t Tuple) Count(value Object) (int, error) {
	count := 0
	for _, item := range t.items {
		if b, err := Compare(item, EQ, value); err == nil {
			yes, _ := b.IsTrue()
			if yes {
				count++
			}
		} else {
			return 0, err
		}
	}
	return count, nil
}

func (t Tuple) Contains(value Object) (bool, error) {
	for _, item := range t.items {
		if b, err := Compare(item, EQ, value); err == nil {
			yes, _ := b.IsTrue()
			if yes {
				return true, nil
			}
		} else {
			return false, err
		}
	}
	return false, nil
}

func (t Tuple) Index(value Object) (int, error) {
	panic("not implemented")
}

func (t Tuple) List() (Object, error) {
	panic("not implemented")
}

func (t Tuple) Tuple() (Object, error) {
	return t, nil
}

