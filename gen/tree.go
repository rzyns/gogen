package gen

import (
	"go/ast"
	"go/printer"
	"go/token"
	"strings"
)

// Tree is a stand-alone AST
type Tree struct {
	PackageName string
	FileSet     *token.FileSet
	AST         *ast.File
}

// NewTree news a thing
func NewTree(pkgName string) *Tree {
	var fs = token.NewFileSet()

	f := &ast.File{
		Name:       ast.NewIdent(pkgName),
		Decls:      make([]ast.Decl, 0),
		Scope:      ast.NewScope(nil),
		Imports:    make([]*ast.ImportSpec, 0),
		Unresolved: make([]*ast.Ident, 0),
		Comments:   make([]*ast.CommentGroup, 0),
	}

	// parser.ParseFile(fs, "", fmt.Sprintf())

	// printer.Fprint(os.Stdout, fs, 0)
	// ast.Print(fs, f)
	// printer.Fprint(os.Stdout, fs, f)

	return &Tree{
		PackageName: pkgName,
		FileSet:     fs,
		AST:         f,
	}
}

func (t *Tree) String() string {
	b := &strings.Builder{}
	printer.Fprint(b, t.FileSet, t.AST)

	return b.String()
}

// AddFunction adds a named function declaration to the tree
func (t *Tree) AddFunction(f *NamedFunction) {
	t.AST.Decls = append(t.AST.Decls, f.AST())
}

// AddStructDef adds a struct typedef
func (t *Tree) AddStructDef(name string, s *Struct) {
	ident := ast.NewIdent(name)
	ident.Obj = ast.NewObj(ast.Typ, name)

	decl := &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: ident,
				Type: s.AST(),
			},
		},
	}

	t.AST.Decls = append(t.AST.Decls, decl)
}
