package main

import (
	"fmt"
	"go/token"

	"github.com/go-toolsmith/strparse"
	"github.com/go-toolsmith/pkgload"
	"golang.org/x/tools/go/packages"
)

func zmain() {
  d := strparse.Decl("func foo() {}")
  fmt.Println(d)
  fmt.Println(strparse.BadDecl)
  fmt.Println(d == strparse.BadDecl)
}

func fmain() {
	fset := token.NewFileSet()
	cfg := packages.Config{
		Mode:  packages.LoadSyntax,
		Tests: true,
		Fset:  fset,
	}
	patterns := []string{"mypackage"}
	pkgs, err := packages.Load(&cfg, patterns...)
	if err != nil {
    panic(err)
	}

	result := pkgs[:0]
	pkgload.VisitUnits(pkgs, func(u *pkgload.Unit) {
		if u.ExternalTest != nil {
			result = append(result, u.ExternalTest)
		}
		result = append(result, u.Base)
	})

  fmt.Println(result)
}
