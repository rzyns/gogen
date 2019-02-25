package gen_test

// https://segment.com/blog/5-advanced-testing-techniques-in-go/

import (
	"testing"

	"github.com/rzyns/gogen/gen"
)

func assertStringsEqual(t *testing.T, actual, expected string) {
	if actual != expected {
		t.Fatalf("\nExpected: %q\nActual:   %q", expected, actual)
	}
}

func Test_NewNamedFunction(t *testing.T) {
	fn := gen.NewNamedFunction("superFunc")

	expected := "func superFunc() {\n}"

	actual := fn.String()

	assertStringsEqual(t, actual, expected)
}

func Test_SetReceiver(t *testing.T) {
	fn := gen.NewNamedFunction("Be")
	fn.SetReceiver("foo", "bee")

	expected := "func (foo bee) Be() {\n}"
	actual := fn.String()

	assertStringsEqual(t, actual, expected)
}

func Test_SetReceiver_Unnamed(t *testing.T) {
	fn := gen.NewNamedFunction("Be")
	fn.SetReceiver("", "bee")

	expected := "func (bee) Be() {\n}"
	actual := fn.String()

	assertStringsEqual(t, actual, expected)
}

func Test_SetReceiverStar(t *testing.T) {
	fn := gen.NewNamedFunction("Be")
	fn.SetReceiver("foo", "*bee")

	expected := "func (foo *bee) Be() {\n}"
	actual := fn.String()

	assertStringsEqual(t, actual, expected)
}

func Test_SetReceiverStar_Unnamed(t *testing.T) {
	fn := gen.NewNamedFunction("Be")
	fn.SetReceiver("", "*bee")

	expected := "func (*bee) Be() {\n}"
	actual := fn.String()

	assertStringsEqual(t, actual, expected)
}

func Test_RemoveReceiver(t *testing.T) {
	fn := gen.NewNamedFunction("Be")
	fn.SetReceiver("foo", "*bee")
	fn.RemoveReceiver()

	expected := "func Be() {\n}"
	actual := fn.String()

	assertStringsEqual(t, actual, expected)
}

func Test_AddParameter(t *testing.T) {
	fn := gen.NewNamedFunction("takesParam")
	fn.AddParameter("foo", "Bar")

	expected := "func takesParam(foo Bar) {\n}"
	actual := fn.String()

	assertStringsEqual(t, actual, expected)
}

func Test_AddParameters(t *testing.T) {
	fn := gen.NewNamedFunction("takesParam")
	fn.AddParameter("foo", "Bar")
	fn.AddParameter("baz", "blam")

	expected := "func takesParam(foo Bar, baz blam) {\n}"
	actual := fn.String()

	assertStringsEqual(t, actual, expected)

	fn.AddParameter("bill", "blam")
	fn.AddParameter("zip", "Bar")

	expected = "func takesParam(foo Bar, baz, bill blam, zip Bar) {\n}"
	actual = fn.String()

	assertStringsEqual(t, actual, expected)
}

func Test_AddParameters_Unnamed(t *testing.T) {
	fn := gen.NewNamedFunction("takesParam")
	fn.AddParameter("", "Bar")
	fn.AddParameter("", "blam")

	expected := "func takesParam(Bar, blam) {\n}"
	actual := fn.String()

	assertStringsEqual(t, actual, expected)

	fn.AddParameter("", "blam")
	fn.AddParameter("", "Bar")

	expected = "func takesParam(Bar, blam, blam, Bar) {\n}"
	actual = fn.String()

	assertStringsEqual(t, actual, expected)
}

func Test_RemoveParameter(t *testing.T) {
	fn := gen.NewNamedFunction("takesParam")
	fn.AddParameter("foo", "Bar")
	fn.AddParameter("baz", "blam")
	fn.AddParameter("bill", "blam")
	fn.AddParameter("zip", "Bar")

	fn.RemoveParameter("baz")

	expected := "func takesParam(foo Bar, bill blam, zip Bar) {\n}"
	actual := fn.String()
	assertStringsEqual(t, actual, expected)

	fn.RemoveParameter("bill")

	expected = "func takesParam(foo Bar, zip Bar) {\n}"
	actual = fn.String()
	assertStringsEqual(t, actual, expected)
}
