package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const title = `
# Stadio

Compendium of functions, data structures, monadic wrappers and more which, hopefully, will be
included as a standard library of the language

## Modules

`

type (
	fun struct {
		name    string
		comment string
		body    *bytes.Buffer
	}

	mod struct {
		title       string
		description string
		funs        []fun
	}
)

func (m *mod) CutreSort() {
	for i := len(m.funs) - 1; i > 0; i-- {
		jmax := i

		for j := 0; j < i; j++ {
			f2 := m.funs[j]

			if strings.Compare(m.funs[jmax].name, f2.name) == -1 {
				jmax = j
			}
		}

		tmp := m.funs[i]
		m.funs[i] = m.funs[jmax]
		m.funs[jmax] = tmp
	}
}

func (m mod) String() string {
	buf := bytes.NewBuffer(nil)

	buf.WriteString("### " + strings.ToUpper(m.title[:1]) + m.title[1:])
	buf.WriteString("\n\n" + m.description)
	buf.WriteString("\n")

	// index
	buf.WriteString("Table of contents\n")

	for _, fn := range m.funs {
		buf.WriteString(fmt.Sprintf("\n- [%s](#%s)", fn.name, fn.name))
	}

	buf.WriteString("\n\n")

	for _, fn := range m.funs {
		buf.WriteString(fn.String())
	}

	return buf.String()
}

func (f fun) String() string {
	buf := bytes.NewBuffer(nil)

	buf.WriteString("#### " + f.name)
	buf.WriteString("\n\n" + f.comment)

	buf.WriteString("\n\n")
	buf.WriteString("<details><summary>Code</summary>")
	buf.WriteString("\n\n```go\n\n" + f.body.String() + "\n```\n\n</details>\n\n")

	return buf.String()
}

func cleanComment(c string) string {
	return strings.ReplaceAll(c, "// ", "")
}

func readme() {
	modules := []string{
		"slices", "maps", "fp",
	}

	buf := bufio.NewWriter(os.Stderr)
	_, _ = buf.WriteString(title)

	mods := make([]*mod, 0)

	for _, modName := range modules {
		_ = filepath.WalkDir(modName, func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() {
				return nil
			}
			if strings.Contains(path, "_test") {
				return nil
			}

			fset := token.NewFileSet()

			fmt.Printf("parsing '%s'...\n", path)
			defer fmt.Println("...done")

			// Parse src but stop after processing the imports.
			f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
			if err != nil {
				fmt.Println(err)
				return nil
			}

			m := new(mod)
			mods = append(mods, m)

			// Remove the first variable declaration from the list of declarations.
			for _, decl := range f.Decls {
				m.title = f.Name.Name
				m.description = cleanComment(f.Doc.Text())

				if fn, ok := decl.(*ast.FuncDecl); ok {
					if fn.Doc == nil {
						// skip non-documented functions
						continue
					}

					if fn.Recv != nil {
						// skip methods
						continue
					}

					if !fn.Name.IsExported() {
						// skip local module functions
						continue
					}

					modFun := fun{
						name:    fn.Name.String(),
						comment: cleanComment(fn.Doc.Text()),
						body:    new(bytes.Buffer),
					}

					// remove comments from node in order to make go/format print the body without comments
					fn.Doc = nil

					if err = format.Node(modFun.body, fset, fn); err != nil {
						fmt.Println(err)
						continue
					}

					m.funs = append(m.funs, modFun)
				}
			}

			return nil
		})
	}

	for _, m := range mods {
		m.CutreSort()

		_, _ = buf.WriteString(m.String())
		_, _ = buf.WriteString("\n\n<br/>\n\n")
	}

	_ = buf.Flush()
}

func main() {
	readme()
}
