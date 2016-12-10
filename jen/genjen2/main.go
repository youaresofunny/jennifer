package main

import (
	"context"

	"strings"

	. "github.com/davelondon/jennifer/jen"
	"github.com/davelondon/jennifer/jen/data"
)

func main() {

	file := NewFile()
	for _, b := range data.Blocks {
		b := b // b used in closures
		if b.Name == "" {
			continue
		}
		comment := Commentf("%s inserts %s", b.Name, b.Desc)

		/*
			func {Name}(c ...Code) *Group {
				return newStatement().{Name}(c...)
			}
		*/
		file.Add(comment)
		file.Func().Id(b.Name).Params(
			Id("c").Op("...").Id("Code"),
		).Op("*").Id("Group").Block(
			Return().Id("newStatement").Call().Op(".").Id(b.Name).Call(
				Id("c").Op("..."),
			),
		)

		/*
			func (g *Group) {Name}(c ...Code) *Group {
				if startNewStatement(g.syntax) {
					s := {Name}(c...)
					g.items = append(g.items, s)
					return s
				}
				s := Group{
					syntax: {Syntax},
					items:  c,
				}
				g.items = append(g.items, s)
				return g
			}
		*/
		file.Add(comment)
		file.Func().Params(
			Id("g").Op("*").Id("Group"),
		).Id(b.Name).Params(
			Id("c").Op("...").Id("Code"),
		).Op("*").Id("Group").Block(
			If().Id("startNewStatement").Call(
				Id("g", "syntax"),
			).Block(
				Id("s").Op(":=").Id(b.Name).Call(
					Id("c").Op("..."),
				),
				Id("g", "items").Op("=").Append(
					Id("g", "items"),
					Id("s"),
				),
				Return(Id("s")),
			),
			Id("s").Op(":=").Id("Group").Lit(map[string]Code{
				"syntax": Id(b.Syntax),
				"items":  Id("c"),
			}),
			Id("g", "items").Op("=").Append(
				Id("g", "items"),
				Id("s"),
			),
			Return(Id("g")),
		)
	}

	type token struct {
		name      string
		cap       string
		tokenType string
		tokenDesc string
	}
	tokens := []token{}
	for _, v := range data.Identifiers {
		tokens = append(tokens, token{
			name:      v,
			cap:       strings.ToUpper(v[:1]) + v[1:],
			tokenType: "identifierToken",
			tokenDesc: "identifier",
		})
	}
	for _, v := range data.Keywords {
		tokens = append(tokens, token{
			name:      v,
			cap:       strings.ToUpper(v[:1]) + v[1:],
			tokenType: "keywordToken",
			tokenDesc: "keyword",
		})
	}

	for _, t := range tokens {
		t := t // used in closures
		comment := Commentf(
			"%s inserts the %s %s",
			t.cap,
			t.name,
			t.tokenDesc,
		)
		/*
			func {Name}() *Group {
				return newStatement().{Name}()
			}
		*/
		file.Add(comment)
		file.Func().Id(t.cap).Params().Op("*").Id("Group").Block(
			Return().Id("newStatement").Call().Op(".").Id(t.cap).Call(),
		)

		/*
			func (g *Group) {Name}() *Group {
				if startNewStatement(g.syntax) {
					s := {Name}()
					g.items = append(g.items, s)
					return s
				}
				t := Token{
					Group:    g,
					typ:     {identifierToken|keywordToken},
					content: "{Name}",
				}
				g.items = append(g.items, t)
				return g
			}
		*/
		file.Add(comment)
		file.Func().Params(
			Id("g").Op("*").Id("Group"),
		).Id(t.cap).Params().Op("*").Id("Group").Block(
			If().Id("startNewStatement").Call(
				Id("g", "syntax"),
			).Block(
				Id("s").Op(":=").Id(t.cap).Call(),
				Id("g", "items").Op("=").Append(
					Id("g", "items"),
					Id("s"),
				),
				Return(Id("s")),
			),
			Id("t").Op(":=").Id("Token").Lit(map[string]Code{
				"Group":   Id("g"),
				"typ":     Id(t.tokenType),
				"content": Lit(t.name),
			}),
			Id("g", "items").Op("=").Append(
				Id("g", "items"),
				Id("t"),
			),
			Return(Id("g")),
		)
	}

	for _, f := range data.Functions {
		f := f // used in closure
		capName := strings.ToUpper(f.Name[:1]) + f.Name[1:]
		desc := "built in function"
		if f.NoParens {
			desc = "keyword"
		}
		comment := Commentf(
			"%s inserts the %s %s",
			capName,
			f.Name,
			desc,
		)
		/*
			func {Name}(c ...Code) *Group {
				return newStatement().{Name}(c...)
			}
		*/
		file.Add(comment)
		file.Func().Id(capName).Params(
			Id("c").Op("...").Id("Code"),
		).Op("*").Id("Group").Block(
			Return().Id("newStatement").Call().Op(".").Id(capName).Call(
				Id("c").Op("..."),
			),
		)

		/*
			func (g *Group) {Name}(c ...Code) *Group {
				if startNewStatement(g.syntax) {
					s := {Name}(c...)
					g.items = append(g.items, s)
					return s
				}
				return g.Id("{Name}").{Call|List}(c...)
			}
		*/
		file.Add(comment)
		file.Func().Params(
			Id("g").Op("*").Id("Group"),
		).Id(capName).Params(
			Id("c").Op("...").Id("Code"),
		).Op("*").Id("Group").Block(
			If().Id("startNewStatement").Call(Id("g", "syntax")).Block(
				Id("s").Op(":=").Id(capName).Call(
					Id("c").Op("..."),
				),
				Id("g", "items").Op("=").Append(
					Id("g", "items"),
					Id("s"),
				),
				Return(Id("s")),
			),
			Return().Id("g.Id").Call(
				Lit(f.Name),
			).Op(".").AddFunc(func() Code {
				if f.NoParens {
					return Id("List")
				}
				return Id("Call")
			}).Call(
				Id("c").Op("..."),
			),
		)

	}

	ctx := Context(context.Background(), "jen")
	err := WriteFile(ctx, file, "./generated.go")
	if err != nil {
		panic(err)
	}

}