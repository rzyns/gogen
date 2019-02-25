package gen

import (
	"go/ast"
	"go/token"
	"strings"
)

// Field is an alias for ast.Field
type Field ast.Field

// NewField makes a new field of type typ
func NewField(typ ast.Node) *Field {
	field := &ast.Field{Names: make([]*ast.Ident, 0, 1)}

	return (*Field)(field)
}

// Ptr returns a new copy of f as an *ast.StarExpr
func (f *Field) Ptr() *Field {
	var field *Field
	if f != nil {
		*field = *f
		field.Type = &ast.StarExpr{X: field.Type}
		return field
	}

	return nil
}

// SetName sets f.Names to the single value, name
func (f *Field) SetName(name string) {
	f.Names = []*ast.Ident{ast.NewIdent(name)}
}

// WithTag returns the field's tag
func (f *Field) WithTag(tag string) *Field {
	if len(tag) > 0 {
		f.Tag = &ast.BasicLit{
			Kind:  token.STRING,
			Value: tag,
		}
	}

	return f

	// cpy := (*Field)(astcopy.Field(f.AST()))

	// cpy.Tag = &ast.BasicLit{
	// 	Kind:  token.STRING,
	// 	Value: "`" + strings.Replace(tag, "`", "", -1) + "`",
	// }

	// return cpy
}

// MakeField makes a field from name and type _string_ (e.g., "Foo" or "*Foo")
func MakeField(name, typ string) *Field {
	var typeName = strings.TrimPrefix(typ, "*")

	typeIdent := ast.NewIdent(typeName)
	typeIdent.Obj = ast.NewObj(ast.Typ, typeName)

	// field := &ast.Field{Names: []*ast.Ident{nameIdent}}
	field := &ast.Field{Names: make([]*ast.Ident, 0, 1)}
	if len(name) > 0 {
		nameIdent := ast.NewIdent(name)
		nameIdent.Obj = ast.NewObj(ast.Var, name)
		field.Names = append(field.Names, nameIdent)
	}

	if strings.HasPrefix(typ, "*") {
		field.Type = &ast.StarExpr{X: typeIdent}
	} else {
		field.Type = typeIdent
	}

	return (*Field)(field)
}

// AST returns the go/ast *Field
func (f *Field) AST() *ast.Field {
	return (*ast.Field)(f)
}
