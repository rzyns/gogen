package gen_test

import (
	"testing"

	"github.com/rzyns/gogen/gen"
)

func Test_StructTag_String(t *testing.T) {
	st := gen.NewStructTag().WithValue("fOo", "bar", "").WithValue("fOo", "baz", "derp")

	expected := `fOo:"bar, baz=derp"`
	actual := st.String()

	assertStringsEqual(t, actual, expected)

	st = gen.NewStructTag().WithValue("json", "", "").WithValue("json", "omitempty", "")

	expected = `json:", omitempty"`
	actual = st.String()

	assertStringsEqual(t, actual, expected)
}

func Test_Struct_AddField(t *testing.T) {
	s := gen.NewStruct()
	s.AddField(gen.MakeField("Diarrhea", "*Poop").WithTag("`foo`"))

	expected := "struct {\n\tDiarrhea *Poop `foo`\n}"
	actual := s.String()

	assertStringsEqual(t, actual, expected)
}
