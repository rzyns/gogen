package gen_test

import (
	"testing"

	"github.com/rzyns/gogen/gen"
)

func Test_NewTree(t *testing.T) {
	tree := gen.NewTree("myname")
	expected := "package myname\n"
	actual := tree.String()

	assertStringsEqual(t, actual, expected)
}

func Test_AddFunction(t *testing.T) {
	tree := gen.NewTree("myname")
	fn := gen.NewNamedFunction("superFunc")

	expected := "package myname\n\nfunc superFunc() {\n}\n"

	tree.AddFunction(fn)

	actual := tree.String()

	assertStringsEqual(t, actual, expected)
}

func Test_AddStructDef(t *testing.T) {
	tree := gen.NewTree("myname")
	s := gen.NewStruct()

	tree.AddStructDef("foo", s)

	expected := "package myname\n\ntype foo struct {\n}\n"
	actual := tree.String()

	assertStringsEqual(t, actual, expected)
}
