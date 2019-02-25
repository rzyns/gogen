package gen

import (
	"go/ast"
	"go/format"
	"go/printer"
	"go/token"
	"strings"
)

// NamedFunction is a wrapper around ast.FuncDecl
type NamedFunction ast.FuncDecl

// NewNamedFunction creates a new *NamedFunction with name `name`
func NewNamedFunction(name string) *NamedFunction {
	obj := ast.NewObj(ast.Fun, name)
	ident := ast.NewIdent(name)

	ident.Obj = obj

	funcDecl := &NamedFunction{
		Name: ident,
		// Recv: &ast.FieldList{},
		Type: &ast.FuncType{},
		Body: &ast.BlockStmt{},
	}

	obj.Decl = funcDecl

	// return (*NamedFunction)(funcDecl)
	return funcDecl
}

// AST returns the go/ast *FuncDecl cast
func (f *NamedFunction) AST() *ast.FuncDecl {
	return (*ast.FuncDecl)(f)
}

// RemoveReceiver removes any receiver set on f
func (f *NamedFunction) RemoveReceiver() {
	f.Recv = nil
}

// SetReceiver sets the function's receiver
func (f *NamedFunction) SetReceiver(name, typ string) {
	field := MakeField(name, typ).AST()

	f.Recv = &ast.FieldList{
		List: []*ast.Field{field},
	}
}

// AddParameter adds a parameter to the function
func (f *NamedFunction) AddParameter(name, typ string) {

	if f.Type.Params == nil {
		f.Type.Params = &ast.FieldList{}
	}

	if f.Type.Params.List == nil {
		f.Type.Params.List = make([]*ast.Field, 0)
	}

	// if the last parameter is the same type as this one, add name to that
	// instead of making a whole new field (e.g., "bar, baz string")
	if len(f.Type.Params.List) > 1 {
		i := len(f.Type.Params.List) - 1
		if t, ok := f.Type.Params.List[i].Type.(*ast.Ident); ok {
			if t.Name == typ && name != "" {
				id := ast.NewIdent(name)
				f.Type.Params.List[i].Names = append(f.Type.Params.List[i].Names, id)
				return
			}
		}
	}

	field := MakeField(name, typ).AST()
	f.Type.Params.List = append(f.Type.Params.List, field)
}

// RemoveParameter removes the named parameter
func (f *NamedFunction) RemoveParameter(name string) {
	// TODO: not done!
	if f.Type.Params == nil {
		return
	}

	if f.Type.Params.List == nil {
		return
	}

	if len(f.Type.Params.List) < 1 {
		return
	}

	for i := 0; i < len(f.Type.Params.List); i++ {
		names := make([]*ast.Ident, 0, len(f.Type.Params.List[i].Names))

		// Make a new list of names excluding the one we're removing
		for _, n := range f.Type.Params.List[i].Names {
			if n.Name != name {
				names = append(names, n)
			}
		}

		// Only replace the old list with our new one if they're different
		l := len(names)
		if l != len(f.Type.Params.List[i].Names) {
			if l > 0 {
				f.Type.Params.List[i].Names = names
			} else {
				// remove Field if no Names left
				copy(f.Type.Params.List[i:], f.Type.Params.List[i+1:])
				f.Type.Params.List[len(f.Type.Params.List)-1] = nil
				f.Type.Params.List = f.Type.Params.List[:len(f.Type.Params.List)-1]
			}
		}
	}
}

// NameString returns the function's actual string name (not an *ast.Ident)
func (f *NamedFunction) NameString() string {
	return f.Name.String()
}

func (f *NamedFunction) String() string {
	b := &strings.Builder{}
	funcDecl := (*ast.FuncDecl)(f)

	fs := token.NewFileSet()
	err := printer.Fprint(b, fs, funcDecl)
	if err != nil {
		panic(err)
	}

	return b.String()
}

// GofmtString uses format.Node() from go/format, which matches gofmt
func (f *NamedFunction) GofmtString() string {
	b := &strings.Builder{}

	funcDecl := (*ast.FuncDecl)(f)
	err := format.Node(b, nil, funcDecl)
	if err != nil {
		return ""
	}

	return b.String()
}
