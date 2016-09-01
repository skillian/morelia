package morelia

import (
	"fmt"
)

// PyObject_HasAttrString for Go.  Creates a string object and passes that to
// the object's HasAttr method.
func HasAttrString(o Object, name string) bool {
	s := NewStringObject(name)
	return o.HasAttr(s)
}

// PyObject_GetAttrString for Go.  Creates a string object and passes that to
// the object's GetAttr method.
func GetAttrString(o Object, name string) (Object, error) {
	s := NewStringObject(name)
	return o.GetAttr(s)
}

// PyObject_SetAttrString for Go.  Creates a string object and passes that to
// the object's SetAttr method.
func SetAttrString(o Object, name string, value Object) error {
	s := NewStringObject(name)
	return o.SetAttr(s, value)
}

// PyObject_DelAttrString for Go.  Creates a string object and passes that to
// the object's DelAttr method.
func DelAttrString(o Object, name string) error {
	s := NewStringObject(name)
	return o.DelAttr(s)
}

// Implementation of PyObject_RichCompare based on CPython's do_richcompare
// function.
func Compare(o1 Object, op Op, o2 Object) (Object, error) {
	var checked_reverse_op bool = false

	// If o2 is an instance of a subclass of o1's type, use its comparison
	// instead:
	t1 := o1.Type()
	t2 := o2.Type()
	if t1 != t2 && IsSubtype(t2, t1) {
		checked_reverse_op = true
		if res, err := o2.RichCompare(o1, SwappedOp[op]); err != NotImplemented {
			// as long as it's implemented, then it should be returned, even
			// if there was an error.
			return res, err
		}
	}
	// otherwise try the first object's comparison
	if res, err := o1.RichCompare(o2, op); err != NotImplemented {
		return res, err
	}
	// if that failed, and we didn't start with the second object's comparison,
	// try that one.
	if !checked_reverse_op {
		if res, err := o2.RichCompare(o1, SwappedOp[op]); err != NotImplemented {
			return res, err
		}
	}
	// else use a "sensible" default :)
	switch op {
	case EQ:
		return NewBoolObject(o1 == o2), nil
	case NE:
		return NewBoolObject(o1 != o2), nil
	default:
		return nil, fmt.Errorf("'%s' is not supported between instances of '%s' and '%s'", op, o1.Type().Name, o2.Type().Name)
	}
}

// walk t1's bases to see if any of them are t2.  If so, return true.  If none
// of t1's bases are t2, then return false.
func typeIsSubtypeBaseChain(t1, t2 *Type) bool {
	for {
		if t1 == t2 {
			return true
		}
		t1 = t1.Base
		if t1 == nil {
			return false
		}
	}
}

// Check if t1 is a subtype of t2 first by looking at t1's MRO tuple.  If that
// isn't set, then walk the base-chain.
func checkSubtype(t1, t2 *Type) bool {
	if t1.Mro != nil {
		mro := t1.Mro.(Tuple)
		l, _ := mro.Length()
		for i := 0; i < l; i++ {
			if c, err := mro.GetItemInt(i); err == nil {
				if t2 == c {
					return true
				}
			} else {
				panic(err)
			}
		}
		return false
	} else {
		return typeIsSubtypeBaseChain(t1, t2)
	}
}

// Check if Type a is a subclass of Type b.
func IsSubtype(a, b Object) bool {
	if t1, ok := a.(*Type); ok {
		if t2, ok := b.(*Type); ok {
			if t1 == t2 {
				return true
			} else {
				return checkSubtype(t1, t2)
			}
		}
	}
	return false;
}

// PyObject_RichCompareBool for Go.  Calls the object's RichCompare method and
// then converts the returned object to a Go bool.
func RichCompareBool(o1, o2 Object, op Op) (bool, error) {
	if r, err := o1.RichCompare(o2, op); err == nil {
		return r.IsTrue()
	} else {
		return false, err
	}
}

// Call the Object's Length method.
func Size(o Object) (int, error) {
	return o.Length()
}

// Number protocol:

func doBinaryOp(o1, o2 Object, fn1, fn2 func(o2 Object) (Object, error)) (Object, error) {
	t1 := o1.Type()
	t2 := o2.Type()
	if fn1 != nil {
		if fn2 != nil && t1 != t2 && IsSubtype(t1, t2) {
			if res, err := fn2(o1); err != NotImplemented {
				return res, err
			}
			fn2 = nil
		}
		if res, err := fn1(o2); err != NotImplemented {
			return res, err
		}
	}
	if fn2 != nil {
		if res, err := fn2(o1); err != NotImplemented {
			return res, err
		}
	}
	return nil, NotImplemented
}

func Add(o1, o2 Object) (Object, error) {
	if n1, ok := o1.(Number); ok {
		if n2, ok := o2.(Number); ok {
			return doBinaryOp(o1, o2, n1.Add, n2.Add)
		}
	}
	return nil, NotImplemented
}

func Subtract(o1, o2 Object) (Object, error) {
	if n1, ok := o1.(Number); ok {
		if n2, ok := o2.(Number); ok {
			return doBinaryOp(o1, o2, n1.Subtract, n2.Subtract)
		}
	}
	return nil, NotImplemented
}

func Multiply(o1, o2 Object) (Object, error) {
	if n1, ok := o1.(Number); ok {
		if n2, ok := o2.(Number); ok {
			return doBinaryOp(o1, o2, n1.Multiply, n2.Multiply)
		}
	}
	return nil, NotImplemented
}

func FloorDivide(o1, o2 Object) (Object, error) {
	if n1, ok := o1.(Number); ok {
		if n2, ok := o2.(Number); ok {
			return doBinaryOp(o1, o2, n1.FloorDivide, n2.FloorDivide)
		}
	}
	return nil, NotImplemented
}

func TrueDivide(o1, o2 Object) (Object, error) {
	if n1, ok := o1.(Number); ok {
		if n2, ok := o2.(Number); ok {
			return doBinaryOp(o1, o2, n1.TrueDivide, n2.TrueDivide)
		}
	}
	return nil, NotImplemented
}

