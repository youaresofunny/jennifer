package jen_test

import (
	"bytes"
	"context"
	"fmt"
	"go/format"
	"reflect"
	"strings"
	"testing"

	. "github.com/davelondon/jennifer/jen"
)

var cases = []tc{
	{
		desc:   `empty block`,
		code:   Block(),
		expect: `{}`,
	},
	{
		desc:   `string literal`,
		code:   Lit("a"),
		expect: `"a"`,
	},
	{
		desc:   `int literal`,
		code:   Lit(1),
		expect: `1`,
	},
	{
		desc:   `simple id`,
		code:   Id("a"),
		expect: `a`,
	},
	{
		desc:   `foreign id`,
		code:   Id("x.y/z.a"),
		expect: `z.a`,
		expectImports: map[string]string{
			"x.y/z": "z",
		},
	},
	{
		desc:   `var decl`,
		code:   Var().Id("a").Op("=").Lit("b"),
		expect: `var a = "b"`,
	},
	{
		desc:   `short var decl`,
		code:   Id("a").Op(":=").Lit("b"),
		expect: `a := "b"`,
	},
	{
		desc:   `simple if`,
		code:   If().Id("a").Op("==").Lit("b").Block(),
		expect: `if a == "b" {}`,
	},
	{
		desc: `simple if`,
		code: If().Id("a").Op("==").Lit("b").Block(
			Id("a").Op("++"),
		),
		expect: `if a == "b" { a++ }`,
	},
	{
		desc:   `pointer`,
		code:   Op("*").Id("a"),
		expect: `*a`,
	},
	{
		desc:   `address`,
		code:   Op("&").Id("a"),
		expect: `&a`,
	},
	{
		desc: `simple call`,
		code: Id("a").Call(
			Lit("b"),
			Lit("c"),
		),
		expect: `a("b", "c")`,
	},
	{
		desc: `call fmt.Sprintf`,
		code: Id("fmt.Sprintf").Call(
			Lit("b"),
			Id("c"),
		),
		expect: `fmt.Sprintf("b", c)`,
	},
	{
		desc: `slices`,
		code: Id("a").Index(
			Lit(1),
			Empty(),
		),
		expect: `a[1:]`,
	},
	{
		desc:   `return`,
		code:   Return(Id("a")),
		expect: `return a`,
	},
	{
		desc:   `double return`,
		code:   Return(Id("a"), Id("b")),
		expect: `return a, b`,
	},
	{
		desc: `func`,
		code: Func().Id("a").Params(
			Id("a").String(),
		).Block(
			Return(Id("a")),
		),
		expect: `func a(a string){ return a }`,
	},
	{
		desc:   `built in func`,
		code:   New(Id("a"), Id("b")),
		expect: `new(a, b)`,
	},
	{
		desc:   `multip`,
		code:   Id("a").Op("*").Id("b"),
		expect: `a * b`,
	},
	{
		desc:   `multip ptr`,
		code:   Id("a").Op("*").Op("*").Id("b"),
		expect: `a * *b`,
	},
	{
		desc:   `field`,
		code:   Id("a", "b"),
		expect: `a.b`,
	},
	{
		desc:   `method`,
		code:   Id("a", "b").Call(Id("c"), Id("d")),
		expect: `a.b(c, d)`,
	},
	{
		desc: `if else`,
		code: If().Id("a").Op("==").Lit(1).Block(
			Id("b").Op("=").Lit(1),
		).Else().If().Id("a").Op("==").Lit(2).Block(
			Id("b").Op("=").Lit(2),
		).Else().Block(
			Id("b").Op("=").Lit(3),
		),
		expect: `if a == 1 { b = 1 } else if a == 2 { b = 2 } else { b = 3 }`,
	},
	{
		desc:   `literal array`,
		code:   Index().String().Values(Lit("a"), Lit("b")),
		expect: `[]string{"a", "b"}`,
	},
	{
		desc:   `comment`,
		code:   Comment("a"),
		expect: `// a`,
	},
	{
		desc:   `null`,
		code:   Id("a").Params(Id("b"), Null(), Id("c")),
		expect: `a(b, c)`,
	},
	{
		desc: `map literal`,
		code: Id("a").MapLit(map[Code]Code{
			Id("b"): Id("c"),
		}),
		expect: `a{
			b: c,
		}`,
	},
	{
		desc: `map literal`,
		code: Id("a").Lit(map[Code]Code{
			Id("b"): Id("c"),
		}),
		expect: `a{
			b: c,
		}`,
	},
	{
		desc: `map literal string keys`,
		code: Id("a").Lit(map[string]Code{
			"b": Id("c"),
		}),
		expect: `a{
			b: c,
		}`,
	},
	{
		desc:   `multi id`,
		code:   Id("a", "b", "c"),
		expect: `a.b.c`,
	},
	{
		desc:   `add func`,
		code:   Id("a").Op(".").AddFunc(func() Code { return Id("b") }),
		expect: `a.b`,
	},
}

func TestJen(t *testing.T) {
	for i, c := range cases {
		onlyTest := ""
		if onlyTest != "" && c.desc != onlyTest {
			continue
		}
		b := &bytes.Buffer{}

		if c.path == "" {
			c.path = "a.b/c"
		}
		ctx := Context(context.Background(), c.path)

		err := c.code.Render(ctx, b)
		if err != nil {
			t.Errorf("Error in test case %d failed. Description: %s\nError:\n%s", i, c.desc, err)
		}

		rendered, err := format.Source(b.Bytes())
		if err != nil {
			t.Errorf("Error formatting rendered source in test case %d. Description: %s\nError:\n%s\nSource:\n%s", i, c.desc, err, b.String())
		}

		expected, err := format.Source([]byte(c.expect))
		if err != nil {
			panic(fmt.Sprintf("Error formatting expected source in test case %d. Description: %s\nError:\n%s", i, c.desc, err))
		}

		if strings.TrimSpace(string(rendered)) != strings.TrimSpace(string(expected)) {
			t.Errorf("Test case %d failed. Description: %s\nExpected:\n%s\nOutput:\n%s", i, c.desc, expected, rendered)
		}

		if c.expectImports != nil {
			f := FromContext(ctx)
			if !reflect.DeepEqual(f.Imports, c.expectImports) {
				t.Errorf("Test case %d failed. Description: %s\nImports expected:\n%s\nOutput:\n%s", i, c.desc, c.expectImports, f.Imports)
			}
		}
	}
}

// a test case
type tc struct {
	// path
	path string
	// description for locating the test case
	desc string
	// code to generate
	code Code
	// expected generated source
	expect string
	// expected imports
	expectImports map[string]string
}