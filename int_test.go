package morelia

import (
	"testing"
)

var (
	Int_1 = NewIntObject(1)
	Int_2 = NewIntObject(2)
	Int_3 = NewIntObject(3)
	Int_4 = NewIntObject(4)
	Int_5 = NewIntObject(5)
	Int_6 = NewIntObject(6)
)


func testBinOp(t *testing.T, fn func(o1, o2 Object) (Object, error), o1, o2 Object, op Op, expected Object, good, bad string) {
	if res, err := fn(o1, o2); err == nil {
		if eq, err := Compare(res, op, expected); err == nil {
			yes, _ := eq.IsTrue()
			if yes {
				t.Log(good)
			} else {
				t.Error(bad + ":", err)
			}
		} else {
			t.Error("Failed to compare result and expected:", err)
		}
	} else {
		t.Error("Failed to perform calculation:", err)
	}
}


// Equality tests:

func TestIntAddEqual(t *testing.T) {
	t.Parallel()
	testBinOp(t, Add, Int_1, Int_2, EQ, Int_3, "1 + 2 = 3", "Apparently 1 + 2 != 3")
}

func TestIntSubtractEqual(t *testing.T) {
	t.Parallel()
	testBinOp(t, Subtract, Int_3, Int_2, EQ, Int_1, "3 - 2 = 1", "3 - 2 != 1")
}

func TestIntMultiplyEqual(t *testing.T) {
	t.Parallel()
	testBinOp(t, Multiply, Int_2, Int_3, EQ, Int_6, "2 * 3 = 6", "2 * 3 != 6")
}

func TestIntFloorDivideEqual(t *testing.T) {
	t.Parallel()
	testBinOp(t, FloorDivide, Int_6, Int_4, EQ, Int_1, "6 // 4 = 1", "6 // 4 != 1")
}

// Inequality tests:

func TestIntAddNotEqual(t *testing.T) {
	t.Parallel()
	testBinOp(t, Add, Int_1, Int_1, NE, Int_1, "1 + 1 != 1", "1 + 1 == 1")
}

func TestIntSubtractNotEqual(t *testing.T) {
	t.Parallel()
	testBinOp(t, Subtract, Int_4, Int_3, NE, Int_2, "4 - 3 != 2", "4 - 3 == 2")
}

func TestIntMultiplyNotEqual(t *testing.T) {
	t.Parallel()
	testBinOp(t, Multiply, Int_3, Int_3, NE, Int_6, "3 * 3 != 6", "3 * 3 == 6")
}

func TestIntFloorDivideNotEqual(t *testing.T) {
	t.Parallel()
	testBinOp(t, FloorDivide, Int_3, Int_2, NE, Int_2, "3 // 2 != 2", "3 // 2 == 2")
}

// Less-than tests:

func TestIntAddLessThan(t *testing.T) {
	t.Parallel()
	testBinOp(t, Add, Int_1, Int_1, LT, Int_3, "1 + 1 < 3", "1 + 1 >= 3")
}

func TestIntSubtractLessThan(t *testing.T) {
	t.Parallel()
	testBinOp(t, Subtract, Int_4, Int_3, LT, Int_2, "4 - 3 < 2", "4 - 3 >= 2")
}

func TestIntMultiplyLessThan(t *testing.T) {
	t.Parallel()
	testBinOp(t, Multiply, Int_2, Int_2, LT, Int_6, "2 * 2 < 6", "2 * 2 >= 6")
}

func TestIntFloorDivideLessThan(t *testing.T) {
	t.Parallel()
	testBinOp(t, FloorDivide, Int_3, Int_2, LT, Int_2, "3 // 2 < 2", "3 // 2 >= 2")
}

// Less-than-or-equal tests:

func TestIntAddLessThanOrEqual(t *testing.T) {
	t.Parallel()
	testBinOp(t, Add, Int_1, Int_1, LE, Int_3, "1 + 1 <= 3", "1 + 1 > 3")
	testBinOp(t, Add, Int_1, Int_2, LE, Int_3, "1 + 2 <= 3", "1 + 2 > 3")
}

func TestIntSubtractLessThanOrEqual(t *testing.T) {
	t.Parallel()
	testBinOp(t, Subtract, Int_4, Int_3, LE, Int_2, "4 - 3 <= 2", "4 - 3 > 2")
	testBinOp(t, Subtract, Int_4, Int_2, LE, Int_2, "4 - 2 <= 2", "4 - 2 > 2")
}

func TestIntMultiplyLessThanOrEqual(t *testing.T) {
	t.Parallel()
	testBinOp(t, Multiply, Int_2, Int_2, LE, Int_6, "2 * 2 <= 6", "2 * 2 > 6")
	testBinOp(t, Multiply, Int_2, Int_3, LE, Int_6, "2 * 3 <= 6", "2 * 3 > 6")
}

func TestIntFloorDivideLessThanOrEqual(t *testing.T) {
	t.Parallel()
	testBinOp(t, FloorDivide, Int_3, Int_2, LE, Int_2, "3 // 2 <= 2", "3 // 2 > 2")
	testBinOp(t, FloorDivide, Int_6, Int_3, LE, Int_2, "6 // 3 <= 2", "6 // 3 > 2")
}

