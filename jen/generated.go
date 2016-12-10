package jen

// Parens inserts parenthesis
func Parens(c ...Code) *Group { return newStatement().Parens(c...) }

// Parens inserts parenthesis
func (g *Group) Parens(c ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := Parens(c...)
		g.items = append(g.items, s)
		return s
	}
	s := Group{
		syntax: ParensSyntax,
		items:  c,
	}
	g.items = append(g.items, s)
	return g
}

// Braces inserts curly braces
func Braces(c ...Code) *Group { return newStatement().Braces(c...) }

// Braces inserts curly braces
func (g *Group) Braces(c ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := Braces(c...)
		g.items = append(g.items, s)
		return s
	}
	s := Group{
		syntax: BracesSyntax,
		items:  c,
	}
	g.items = append(g.items, s)
	return g
}

// Block inserts curly braces containing a statements list
func Block(c ...Code) *Group { return newStatement().Block(c...) }

// Block inserts curly braces containing a statements list
func (g *Group) Block(c ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := Block(c...)
		g.items = append(g.items, s)
		return s
	}
	s := Group{
		items:  c,
		syntax: BlockSyntax,
	}
	g.items = append(g.items, s)
	return g
}

// Params inserts parenthesis containing a comma separated list
func Params(c ...Code) *Group { return newStatement().Params(c...) }

// Params inserts parenthesis containing a comma separated list
func (g *Group) Params(c ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := Params(c...)
		g.items = append(g.items, s)
		return s
	}
	s := Group{
		syntax: ParamsSyntax,
		items:  c,
	}
	g.items = append(g.items, s)
	return g
}

// List inserts a comma separated list
func List(c ...Code) *Group { return newStatement().List(c...) }

// List inserts a comma separated list
func (g *Group) List(c ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := List(c...)
		g.items = append(g.items, s)
		return s
	}
	s := Group{
		syntax: ListSyntax,
		items:  c,
	}
	g.items = append(g.items, s)
	return g
}

// Values inserts curly braces containing a comma separated list
func Values(c ...Code) *Group { return newStatement().Values(c...) }

// Values inserts curly braces containing a comma separated list
func (g *Group) Values(c ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := Values(c...)
		g.items = append(g.items, s)
		return s
	}
	s := Group{
		items:  c,
		syntax: ValuesSyntax,
	}
	g.items = append(g.items, s)
	return g
}

// Index inserts square brackets containing a colon separated list
func Index(c ...Code) *Group { return newStatement().Index(c...) }

// Index inserts square brackets containing a colon separated list
func (g *Group) Index(c ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := Index(c...)
		g.items = append(g.items, s)
		return s
	}
	s := Group{
		syntax: IndexSyntax,
		items:  c,
	}
	g.items = append(g.items, s)
	return g
}

// Call inserts parenthesis containing a comma separated list
func Call(c ...Code) *Group { return newStatement().Call(c...) }

// Call inserts parenthesis containing a comma separated list
func (g *Group) Call(c ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := Call(c...)
		g.items = append(g.items, s)
		return s
	}
	s := Group{
		syntax: CallSyntax,
		items:  c,
	}
	g.items = append(g.items, s)
	return g
}

// Decls inserts parenthesis containing a statement list
func Decls(c ...Code) *Group { return newStatement().Decls(c...) }

// Decls inserts parenthesis containing a statement list
func (g *Group) Decls(c ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := Decls(c...)
		g.items = append(g.items, s)
		return s
	}
	s := Group{
		items:  c,
		syntax: DeclsSyntax,
	}
	g.items = append(g.items, s)
	return g
}

// Bool inserts the bool identifier
func Bool() *Group { return newStatement().Bool() }

// Bool inserts the bool identifier
func (g *Group) Bool() *Group {
	if startNewStatement(g.syntax) {
		s := Bool()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		Group:   g,
		typ:     identifierToken,
		content: "bool",
	}
	g.items = append(g.items, t)
	return g
}

// Byte inserts the byte identifier
func Byte() *Group { return newStatement().Byte() }

// Byte inserts the byte identifier
func (g *Group) Byte() *Group {
	if startNewStatement(g.syntax) {
		s := Byte()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		content: "byte",
		Group:   g,
		typ:     identifierToken,
	}
	g.items = append(g.items, t)
	return g
}

// Complex64 inserts the complex64 identifier
func Complex64() *Group { return newStatement().Complex64() }

// Complex64 inserts the complex64 identifier
func (g *Group) Complex64() *Group {
	if startNewStatement(g.syntax) {
		s := Complex64()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		content: "complex64",
		Group:   g,
		typ:     identifierToken,
	}
	g.items = append(g.items, t)
	return g
}

// Complex128 inserts the complex128 identifier
func Complex128() *Group { return newStatement().Complex128() }

// Complex128 inserts the complex128 identifier
func (g *Group) Complex128() *Group {
	if startNewStatement(g.syntax) {
		s := Complex128()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		content: "complex128",
		Group:   g,
		typ:     identifierToken,
	}
	g.items = append(g.items, t)
	return g
}

// Error inserts the error identifier
func Error() *Group { return newStatement().Error() }

// Error inserts the error identifier
func (g *Group) Error() *Group {
	if startNewStatement(g.syntax) {
		s := Error()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		typ:     identifierToken,
		content: "error",
		Group:   g,
	}
	g.items = append(g.items, t)
	return g
}

// Float32 inserts the float32 identifier
func Float32() *Group { return newStatement().Float32() }

// Float32 inserts the float32 identifier
func (g *Group) Float32() *Group {
	if startNewStatement(g.syntax) {
		s := Float32()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		Group:   g,
		typ:     identifierToken,
		content: "float32",
	}
	g.items = append(g.items, t)
	return g
}

// Float64 inserts the float64 identifier
func Float64() *Group { return newStatement().Float64() }

// Float64 inserts the float64 identifier
func (g *Group) Float64() *Group {
	if startNewStatement(g.syntax) {
		s := Float64()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		Group:   g,
		typ:     identifierToken,
		content: "float64",
	}
	g.items = append(g.items, t)
	return g
}

// Int inserts the int identifier
func Int() *Group { return newStatement().Int() }

// Int inserts the int identifier
func (g *Group) Int() *Group {
	if startNewStatement(g.syntax) {
		s := Int()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		content: "int",
		Group:   g,
		typ:     identifierToken,
	}
	g.items = append(g.items, t)
	return g
}

// Int8 inserts the int8 identifier
func Int8() *Group { return newStatement().Int8() }

// Int8 inserts the int8 identifier
func (g *Group) Int8() *Group {
	if startNewStatement(g.syntax) {
		s := Int8()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		Group:   g,
		typ:     identifierToken,
		content: "int8",
	}
	g.items = append(g.items, t)
	return g
}

// Int16 inserts the int16 identifier
func Int16() *Group { return newStatement().Int16() }

// Int16 inserts the int16 identifier
func (g *Group) Int16() *Group {
	if startNewStatement(g.syntax) {
		s := Int16()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		Group:   g,
		typ:     identifierToken,
		content: "int16",
	}
	g.items = append(g.items, t)
	return g
}

// Int32 inserts the int32 identifier
func Int32() *Group { return newStatement().Int32() }

// Int32 inserts the int32 identifier
func (g *Group) Int32() *Group {
	if startNewStatement(g.syntax) {
		s := Int32()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		typ:     identifierToken,
		content: "int32",
		Group:   g,
	}
	g.items = append(g.items, t)
	return g
}

// Int64 inserts the int64 identifier
func Int64() *Group { return newStatement().Int64() }

// Int64 inserts the int64 identifier
func (g *Group) Int64() *Group {
	if startNewStatement(g.syntax) {
		s := Int64()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		Group:   g,
		typ:     identifierToken,
		content: "int64",
	}
	g.items = append(g.items, t)
	return g
}

// Rune inserts the rune identifier
func Rune() *Group { return newStatement().Rune() }

// Rune inserts the rune identifier
func (g *Group) Rune() *Group {
	if startNewStatement(g.syntax) {
		s := Rune()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		content: "rune",
		Group:   g,
		typ:     identifierToken,
	}
	g.items = append(g.items, t)
	return g
}

// String inserts the string identifier
func String() *Group { return newStatement().String() }

// String inserts the string identifier
func (g *Group) String() *Group {
	if startNewStatement(g.syntax) {
		s := String()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		Group:   g,
		typ:     identifierToken,
		content: "string",
	}
	g.items = append(g.items, t)
	return g
}

// Uint inserts the uint identifier
func Uint() *Group { return newStatement().Uint() }

// Uint inserts the uint identifier
func (g *Group) Uint() *Group {
	if startNewStatement(g.syntax) {
		s := Uint()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		Group:   g,
		typ:     identifierToken,
		content: "uint",
	}
	g.items = append(g.items, t)
	return g
}

// Uint8 inserts the uint8 identifier
func Uint8() *Group { return newStatement().Uint8() }

// Uint8 inserts the uint8 identifier
func (g *Group) Uint8() *Group {
	if startNewStatement(g.syntax) {
		s := Uint8()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		Group:   g,
		typ:     identifierToken,
		content: "uint8",
	}
	g.items = append(g.items, t)
	return g
}

// Uint16 inserts the uint16 identifier
func Uint16() *Group { return newStatement().Uint16() }

// Uint16 inserts the uint16 identifier
func (g *Group) Uint16() *Group {
	if startNewStatement(g.syntax) {
		s := Uint16()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		Group:   g,
		typ:     identifierToken,
		content: "uint16",
	}
	g.items = append(g.items, t)
	return g
}

// Uint32 inserts the uint32 identifier
func Uint32() *Group { return newStatement().Uint32() }

// Uint32 inserts the uint32 identifier
func (g *Group) Uint32() *Group {
	if startNewStatement(g.syntax) {
		s := Uint32()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		typ:     identifierToken,
		content: "uint32",
		Group:   g,
	}
	g.items = append(g.items, t)
	return g
}

// Uint64 inserts the uint64 identifier
func Uint64() *Group { return newStatement().Uint64() }

// Uint64 inserts the uint64 identifier
func (g *Group) Uint64() *Group {
	if startNewStatement(g.syntax) {
		s := Uint64()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		content: "uint64",
		Group:   g,
		typ:     identifierToken,
	}
	g.items = append(g.items, t)
	return g
}

// Uintptr inserts the uintptr identifier
func Uintptr() *Group { return newStatement().Uintptr() }

// Uintptr inserts the uintptr identifier
func (g *Group) Uintptr() *Group {
	if startNewStatement(g.syntax) {
		s := Uintptr()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		Group:   g,
		typ:     identifierToken,
		content: "uintptr",
	}
	g.items = append(g.items, t)
	return g
}

// True inserts the true identifier
func True() *Group { return newStatement().True() }

// True inserts the true identifier
func (g *Group) True() *Group {
	if startNewStatement(g.syntax) {
		s := True()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		typ:     identifierToken,
		content: "true",
		Group:   g,
	}
	g.items = append(g.items, t)
	return g
}

// False inserts the false identifier
func False() *Group { return newStatement().False() }

// False inserts the false identifier
func (g *Group) False() *Group {
	if startNewStatement(g.syntax) {
		s := False()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		Group:   g,
		typ:     identifierToken,
		content: "false",
	}
	g.items = append(g.items, t)
	return g
}

// Iota inserts the iota identifier
func Iota() *Group { return newStatement().Iota() }

// Iota inserts the iota identifier
func (g *Group) Iota() *Group {
	if startNewStatement(g.syntax) {
		s := Iota()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		Group:   g,
		typ:     identifierToken,
		content: "iota",
	}
	g.items = append(g.items, t)
	return g
}

// Nil inserts the nil identifier
func Nil() *Group { return newStatement().Nil() }

// Nil inserts the nil identifier
func (g *Group) Nil() *Group {
	if startNewStatement(g.syntax) {
		s := Nil()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		Group:   g,
		typ:     identifierToken,
		content: "nil",
	}
	g.items = append(g.items, t)
	return g
}

// Break inserts the break keyword
func Break() *Group { return newStatement().Break() }

// Break inserts the break keyword
func (g *Group) Break() *Group {
	if startNewStatement(g.syntax) {
		s := Break()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		Group:   g,
		typ:     keywordToken,
		content: "break",
	}
	g.items = append(g.items, t)
	return g
}

// Default inserts the default keyword
func Default() *Group { return newStatement().Default() }

// Default inserts the default keyword
func (g *Group) Default() *Group {
	if startNewStatement(g.syntax) {
		s := Default()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		content: "default",
		Group:   g,
		typ:     keywordToken,
	}
	g.items = append(g.items, t)
	return g
}

// Func inserts the func keyword
func Func() *Group { return newStatement().Func() }

// Func inserts the func keyword
func (g *Group) Func() *Group {
	if startNewStatement(g.syntax) {
		s := Func()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		typ:     keywordToken,
		content: "func",
		Group:   g,
	}
	g.items = append(g.items, t)
	return g
}

// Interface inserts the interface keyword
func Interface() *Group { return newStatement().Interface() }

// Interface inserts the interface keyword
func (g *Group) Interface() *Group {
	if startNewStatement(g.syntax) {
		s := Interface()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		Group:   g,
		typ:     keywordToken,
		content: "interface",
	}
	g.items = append(g.items, t)
	return g
}

// Select inserts the select keyword
func Select() *Group { return newStatement().Select() }

// Select inserts the select keyword
func (g *Group) Select() *Group {
	if startNewStatement(g.syntax) {
		s := Select()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		typ:     keywordToken,
		content: "select",
		Group:   g,
	}
	g.items = append(g.items, t)
	return g
}

// Case inserts the case keyword
func Case() *Group { return newStatement().Case() }

// Case inserts the case keyword
func (g *Group) Case() *Group {
	if startNewStatement(g.syntax) {
		s := Case()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		Group:   g,
		typ:     keywordToken,
		content: "case",
	}
	g.items = append(g.items, t)
	return g
}

// Defer inserts the defer keyword
func Defer() *Group { return newStatement().Defer() }

// Defer inserts the defer keyword
func (g *Group) Defer() *Group {
	if startNewStatement(g.syntax) {
		s := Defer()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		typ:     keywordToken,
		content: "defer",
		Group:   g,
	}
	g.items = append(g.items, t)
	return g
}

// Go inserts the go keyword
func Go() *Group { return newStatement().Go() }

// Go inserts the go keyword
func (g *Group) Go() *Group {
	if startNewStatement(g.syntax) {
		s := Go()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		content: "go",
		Group:   g,
		typ:     keywordToken,
	}
	g.items = append(g.items, t)
	return g
}

// Map inserts the map keyword
func Map() *Group { return newStatement().Map() }

// Map inserts the map keyword
func (g *Group) Map() *Group {
	if startNewStatement(g.syntax) {
		s := Map()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		Group:   g,
		typ:     keywordToken,
		content: "map",
	}
	g.items = append(g.items, t)
	return g
}

// Struct inserts the struct keyword
func Struct() *Group { return newStatement().Struct() }

// Struct inserts the struct keyword
func (g *Group) Struct() *Group {
	if startNewStatement(g.syntax) {
		s := Struct()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		typ:     keywordToken,
		content: "struct",
		Group:   g,
	}
	g.items = append(g.items, t)
	return g
}

// Chan inserts the chan keyword
func Chan() *Group { return newStatement().Chan() }

// Chan inserts the chan keyword
func (g *Group) Chan() *Group {
	if startNewStatement(g.syntax) {
		s := Chan()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		content: "chan",
		Group:   g,
		typ:     keywordToken,
	}
	g.items = append(g.items, t)
	return g
}

// Else inserts the else keyword
func Else() *Group { return newStatement().Else() }

// Else inserts the else keyword
func (g *Group) Else() *Group {
	if startNewStatement(g.syntax) {
		s := Else()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		typ:     keywordToken,
		content: "else",
		Group:   g,
	}
	g.items = append(g.items, t)
	return g
}

// Goto inserts the goto keyword
func Goto() *Group { return newStatement().Goto() }

// Goto inserts the goto keyword
func (g *Group) Goto() *Group {
	if startNewStatement(g.syntax) {
		s := Goto()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		typ:     keywordToken,
		content: "goto",
		Group:   g,
	}
	g.items = append(g.items, t)
	return g
}

// Package inserts the package keyword
func Package() *Group { return newStatement().Package() }

// Package inserts the package keyword
func (g *Group) Package() *Group {
	if startNewStatement(g.syntax) {
		s := Package()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		typ:     keywordToken,
		content: "package",
		Group:   g,
	}
	g.items = append(g.items, t)
	return g
}

// Switch inserts the switch keyword
func Switch() *Group { return newStatement().Switch() }

// Switch inserts the switch keyword
func (g *Group) Switch() *Group {
	if startNewStatement(g.syntax) {
		s := Switch()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		Group:   g,
		typ:     keywordToken,
		content: "switch",
	}
	g.items = append(g.items, t)
	return g
}

// Const inserts the const keyword
func Const() *Group { return newStatement().Const() }

// Const inserts the const keyword
func (g *Group) Const() *Group {
	if startNewStatement(g.syntax) {
		s := Const()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		Group:   g,
		typ:     keywordToken,
		content: "const",
	}
	g.items = append(g.items, t)
	return g
}

// Fallthrough inserts the fallthrough keyword
func Fallthrough() *Group { return newStatement().Fallthrough() }

// Fallthrough inserts the fallthrough keyword
func (g *Group) Fallthrough() *Group {
	if startNewStatement(g.syntax) {
		s := Fallthrough()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		Group:   g,
		typ:     keywordToken,
		content: "fallthrough",
	}
	g.items = append(g.items, t)
	return g
}

// If inserts the if keyword
func If() *Group { return newStatement().If() }

// If inserts the if keyword
func (g *Group) If() *Group {
	if startNewStatement(g.syntax) {
		s := If()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		Group:   g,
		typ:     keywordToken,
		content: "if",
	}
	g.items = append(g.items, t)
	return g
}

// Range inserts the range keyword
func Range() *Group { return newStatement().Range() }

// Range inserts the range keyword
func (g *Group) Range() *Group {
	if startNewStatement(g.syntax) {
		s := Range()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		content: "range",
		Group:   g,
		typ:     keywordToken,
	}
	g.items = append(g.items, t)
	return g
}

// Type inserts the type keyword
func Type() *Group { return newStatement().Type() }

// Type inserts the type keyword
func (g *Group) Type() *Group {
	if startNewStatement(g.syntax) {
		s := Type()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		Group:   g,
		typ:     keywordToken,
		content: "type",
	}
	g.items = append(g.items, t)
	return g
}

// Continue inserts the continue keyword
func Continue() *Group { return newStatement().Continue() }

// Continue inserts the continue keyword
func (g *Group) Continue() *Group {
	if startNewStatement(g.syntax) {
		s := Continue()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		typ:     keywordToken,
		content: "continue",
		Group:   g,
	}
	g.items = append(g.items, t)
	return g
}

// For inserts the for keyword
func For() *Group { return newStatement().For() }

// For inserts the for keyword
func (g *Group) For() *Group {
	if startNewStatement(g.syntax) {
		s := For()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		Group:   g,
		typ:     keywordToken,
		content: "for",
	}
	g.items = append(g.items, t)
	return g
}

// Import inserts the import keyword
func Import() *Group { return newStatement().Import() }

// Import inserts the import keyword
func (g *Group) Import() *Group {
	if startNewStatement(g.syntax) {
		s := Import()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		Group:   g,
		typ:     keywordToken,
		content: "import",
	}
	g.items = append(g.items, t)
	return g
}

// Var inserts the var keyword
func Var() *Group { return newStatement().Var() }

// Var inserts the var keyword
func (g *Group) Var() *Group {
	if startNewStatement(g.syntax) {
		s := Var()
		g.items = append(g.items, s)
		return s
	}
	t := Token{
		typ:     keywordToken,
		content: "var",
		Group:   g,
	}
	g.items = append(g.items, t)
	return g
}

// Append inserts the append built in function
func Append(c ...Code) *Group { return newStatement().Append(c...) }

// Append inserts the append built in function
func (g *Group) Append(c ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := Append(c...)
		g.items = append(g.items, s)
		return s
	}
	return g.Id("append").Call(c...)
}

// Cap inserts the cap built in function
func Cap(c ...Code) *Group { return newStatement().Cap(c...) }

// Cap inserts the cap built in function
func (g *Group) Cap(c ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := Cap(c...)
		g.items = append(g.items, s)
		return s
	}
	return g.Id("cap").Call(c...)
}

// Close inserts the close built in function
func Close(c ...Code) *Group { return newStatement().Close(c...) }

// Close inserts the close built in function
func (g *Group) Close(c ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := Close(c...)
		g.items = append(g.items, s)
		return s
	}
	return g.Id("close").Call(c...)
}

// Complex inserts the complex built in function
func Complex(c ...Code) *Group { return newStatement().Complex(c...) }

// Complex inserts the complex built in function
func (g *Group) Complex(c ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := Complex(c...)
		g.items = append(g.items, s)
		return s
	}
	return g.Id("complex").Call(c...)
}

// Copy inserts the copy built in function
func Copy(c ...Code) *Group { return newStatement().Copy(c...) }

// Copy inserts the copy built in function
func (g *Group) Copy(c ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := Copy(c...)
		g.items = append(g.items, s)
		return s
	}
	return g.Id("copy").Call(c...)
}

// Delete inserts the delete built in function
func Delete(c ...Code) *Group { return newStatement().Delete(c...) }

// Delete inserts the delete built in function
func (g *Group) Delete(c ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := Delete(c...)
		g.items = append(g.items, s)
		return s
	}
	return g.Id("delete").Call(c...)
}

// Imag inserts the imag built in function
func Imag(c ...Code) *Group { return newStatement().Imag(c...) }

// Imag inserts the imag built in function
func (g *Group) Imag(c ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := Imag(c...)
		g.items = append(g.items, s)
		return s
	}
	return g.Id("imag").Call(c...)
}

// Len inserts the len built in function
func Len(c ...Code) *Group { return newStatement().Len(c...) }

// Len inserts the len built in function
func (g *Group) Len(c ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := Len(c...)
		g.items = append(g.items, s)
		return s
	}
	return g.Id("len").Call(c...)
}

// Make inserts the make built in function
func Make(c ...Code) *Group { return newStatement().Make(c...) }

// Make inserts the make built in function
func (g *Group) Make(c ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := Make(c...)
		g.items = append(g.items, s)
		return s
	}
	return g.Id("make").Call(c...)
}

// New inserts the new built in function
func New(c ...Code) *Group { return newStatement().New(c...) }

// New inserts the new built in function
func (g *Group) New(c ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := New(c...)
		g.items = append(g.items, s)
		return s
	}
	return g.Id("new").Call(c...)
}

// Panic inserts the panic built in function
func Panic(c ...Code) *Group { return newStatement().Panic(c...) }

// Panic inserts the panic built in function
func (g *Group) Panic(c ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := Panic(c...)
		g.items = append(g.items, s)
		return s
	}
	return g.Id("panic").Call(c...)
}

// Print inserts the print built in function
func Print(c ...Code) *Group { return newStatement().Print(c...) }

// Print inserts the print built in function
func (g *Group) Print(c ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := Print(c...)
		g.items = append(g.items, s)
		return s
	}
	return g.Id("print").Call(c...)
}

// Println inserts the println built in function
func Println(c ...Code) *Group { return newStatement().Println(c...) }

// Println inserts the println built in function
func (g *Group) Println(c ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := Println(c...)
		g.items = append(g.items, s)
		return s
	}
	return g.Id("println").Call(c...)
}

// Real inserts the real built in function
func Real(c ...Code) *Group { return newStatement().Real(c...) }

// Real inserts the real built in function
func (g *Group) Real(c ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := Real(c...)
		g.items = append(g.items, s)
		return s
	}
	return g.Id("real").Call(c...)
}

// Recover inserts the recover built in function
func Recover(c ...Code) *Group { return newStatement().Recover(c...) }

// Recover inserts the recover built in function
func (g *Group) Recover(c ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := Recover(c...)
		g.items = append(g.items, s)
		return s
	}
	return g.Id("recover").Call(c...)
}

// Return inserts the return keyword
func Return(c ...Code) *Group { return newStatement().Return(c...) }

// Return inserts the return keyword
func (g *Group) Return(c ...Code) *Group {
	if startNewStatement(g.syntax) {
		s := Return(c...)
		g.items = append(g.items, s)
		return s
	}
	return g.Id("return").List(c...)
}