package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
)

// parse parses the source code of a single Go source file and returns
// the corresponding ast.File node. The source code may be provided via
// the filePath of the source file via the parameter.
func parse(filePath string) (f *ast.File, err error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}

	fset := token.NewFileSet()
	return parser.ParseFile(fset, "", string(data), 0)
}

// GetStructNames returns all struct names defined in a .go file.
// It uses Go's ast package to retrieve struct properties such as name.
func GetStructNames(filePath string) ([]string, error) {
	var structNames []string

	f, err := parse(filePath)
	if err != nil {
		return nil, err
	}

	// Decls holds top-level declarations
	for _, d := range f.Decls {
		// A GenDecl node (generic declaration node) represents an import, constant, type or variable declaration.
		gen, ok := d.(*ast.GenDecl)
		if !ok || gen.Tok != token.TYPE {
			continue
		}

		for _, spec := range gen.Specs {
			// A TypeSpec node represents a type declaration
			ts, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			structNames = append(structNames, ts.Name.Name)
		}
	}
	return structNames, err
}
