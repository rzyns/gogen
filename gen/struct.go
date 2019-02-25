package gen

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"strings"
)

// Struct is an alias for ast.StructType
type Struct ast.StructType

// NewStruct creates a new struct definition
func NewStruct() *Struct {
	return (*Struct)(&ast.StructType{
		Fields: &ast.FieldList{List: make([]*ast.Field, 0, 1)},
	})
}

// AST returns the cast go/ast StructType of s
func (s *Struct) AST() *ast.StructType {
	return (*ast.StructType)(s)
}

// AddField adds a Field to the Struct
func (s *Struct) AddField(f *Field) {
	if s.Fields == nil {
		s.Fields = &ast.FieldList{List: make([]*ast.Field, 0, 1)}
	}

	s.Fields.List = append(s.Fields.List, f.AST())
}

func (s *Struct) String() string {
	b := &strings.Builder{}
	structType := (*ast.StructType)(s)

	fs := token.NewFileSet()
	err := printer.Fprint(b, fs, structType)
	if err != nil {
		panic(err)
	}

	return b.String()
}

// StructTag represents a struct tag
type StructTag struct {
	props       map[string]map[string]string
	propKeys    []string
	propSubKeys []string
}

// NewStructTag initializes a struct tag
func NewStructTag() *StructTag {
	return &StructTag{
		props:       make(map[string]map[string]string),
		propKeys:    make([]string, 0, 1),
		propSubKeys: make([]string, 0, 1),
	}
}

// WithValue sets "key" to "value" on tag "tag", e.g.:
// `s.WithTagValue("json", "", "")`
func (s *StructTag) WithValue(tag, key, value string) *StructTag {
	if _, exists := s.props[tag]; !exists {
		s.props[tag] = make(map[string]string)
		s.propKeys = append(s.propKeys, tag)
	}

	if _, exists := s.props[tag][key]; !exists {
		s.propSubKeys = append(s.propSubKeys, key)
	}

	s.props[tag][key] = value

	return s
}

func (s *StructTag) String() string {
	tags := make([]string, 0, len(s.props))

	for _, propKey := range s.propKeys {
		values := make([]string, 0, len(s.props[propKey]))

		for _, propSubKey := range s.propSubKeys {
			value := s.props[propKey][propSubKey]
			if len(value) > 0 {
				values = append(values, fmt.Sprintf("%s=%s", propSubKey, value))
			} else {
				values = append(values, propSubKey)
			}
		}

		tags = append(tags, fmt.Sprintf("%s:%q", propKey, strings.Join(values, ", ")))
	}

	return strings.Join(tags, " ")
}

// AST returns the *ast.BasicLit representation of the struct tag
func (s *StructTag) AST() *ast.BasicLit {
	return &ast.BasicLit{
		Kind:  token.STRING,
		Value: s.String(),
	}
}
