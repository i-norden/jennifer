package main

// Keywords is a list of simple keywords
var Keywords = []string{"break", "default", "func", "select", "chan", "else", "const", "fallthrough", "type", "continue", "var", "goto", "defer", "go", "range"}

// Identifiers is a list of simple identifiers
var Identifiers = []string{"bool", "byte", "complex64", "complex128", "error", "float32", "float64", "int", "int8", "int16", "int32", "int64", "rune", "string", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr", "true", "false", "iota", "nil", "err"}

// Groups is a list of complex groups
var Groups = []struct {
	Name        string   // name of the function / method
	Desc        string   // comment appended to name
	Variadic    bool     // is the parameter variadic?
	Open        string   // opening token
	Close       string   // closing token
	Separator   string   // separator token
	Parameters  []string // parameter names
	PreventFunc bool     // prevent the fooFunc function/method
}{
	{
		Name:       "Parens",
		Desc:       "renders a single item in parenthesis. Use for type conversion or to specify evaluation order.",
		Variadic:   false,
		Open:       "(",
		Close:      ")",
		Separator:  "",
		Parameters: []string{"item"},
	},
	{
		Name:       "List",
		Desc:       "renders a comma separated list with no open or closing tokens. Use for multiple return functions.",
		Variadic:   true,
		Open:       "",
		Close:      "",
		Separator:  ",",
		Parameters: []string{"items"},
	},
	{
		Name:       "Values",
		Desc:       "renders a comma separated list enclosed by curly braces. Use for slice literals.",
		Variadic:   true,
		Open:       "{",
		Close:      "}",
		Separator:  ",",
		Parameters: []string{"values"},
	},
	{
		Name:       "Index",
		Desc:       "renders a colon separated list enclosed by square brackets. Use for array / slice indexes and definitions.",
		Variadic:   true,
		Open:       "[",
		Close:      "]",
		Separator:  ":",
		Parameters: []string{"items"},
	},
	{
		Name:       "Block",
		Desc:       "renders a statement list enclosed by curly braces. Use for all code blocks.",
		Variadic:   true,
		Open:       "{",
		Close:      "}",
		Separator:  "\n",
		Parameters: []string{"statements"},
	},
	{
		Name:       "Defs",
		Desc:       "renders a list of statements enclosed in parenthesis. Use for definition lists.",
		Variadic:   true,
		Open:       "(",
		Close:      ")",
		Separator:  "\n",
		Parameters: []string{"definitions"},
	},
	{
		Name:       "Call",
		Desc:       "renders a comma separated list enclosed by parenthesis. Use for function calls.",
		Variadic:   true,
		Open:       "(",
		Close:      ")",
		Separator:  ",",
		Parameters: []string{"params"},
	},
	{
		Name:       "Params",
		Desc:       "renders a comma separated list enclosed by parenthesis. Use for function parameters and method receivers.",
		Variadic:   true,
		Open:       "(",
		Close:      ")",
		Separator:  ",",
		Parameters: []string{"params"},
	},
	{
		Name:       "CaseBlock",
		Desc:       "renders a statement list preceded by a colon. Use to build switch / select statements.",
		Variadic:   true,
		Open:       ":",
		Close:      "",
		Separator:  "\n",
		Parameters: []string{"statements"},
	},
	{
		Name:       "Assert",
		Desc:       "renders a period followed by a single item enclosed by parenthesis. Use for type assertions.",
		Variadic:   false,
		Open:       ".(",
		Close:      ")",
		Separator:  "",
		Parameters: []string{"typ"},
	},
	{
		Name:       "Map",
		Desc:       "renders the map keyword followed by a single item enclosed by square brackets. Use for map definitions.",
		Variadic:   false,
		Open:       "map[",
		Close:      "]",
		Separator:  "",
		Parameters: []string{"typ"},
	},
	{
		Name:       "If",
		Desc:       "renders the if keyword followed by a semicolon separated list.",
		Variadic:   true,
		Open:       "if ",
		Close:      "",
		Separator:  ";",
		Parameters: []string{"conditions"},
	},
	{
		Name:       "Return",
		Desc:       "renders the return keyword followed by a comma separated list.",
		Variadic:   true,
		Open:       "return ",
		Close:      "",
		Separator:  ",",
		Parameters: []string{"results"},
	},
	{
		Name:       "For",
		Desc:       "renders the for keyword followed by a semicolon separated list.",
		Variadic:   true,
		Open:       "for ",
		Close:      "",
		Separator:  ";",
		Parameters: []string{"conditions"},
	},
	{
		Name:       "Switch",
		Desc:       "renders the switch keyword followed by a semicolon separated list.",
		Variadic:   true,
		Open:       "switch ",
		Close:      "",
		Separator:  ";",
		Parameters: []string{"conditions"},
	},
	{
		Name:       "Interface",
		Desc:       "renders the interface keyword followed by curly braces containing a method list.",
		Variadic:   true,
		Open:       "interface{",
		Close:      "}",
		Separator:  "\n",
		Parameters: []string{"methods"},
	},
	{
		Name:       "Struct",
		Desc:       "renders the struct keyword followed by curly braces containing a field list.",
		Variadic:   true,
		Open:       "struct{",
		Close:      "}",
		Separator:  "\n",
		Parameters: []string{"fields"},
	},
	{
		Name:       "Case",
		Desc:       "renders the case keyword followed by a comma separated list.",
		Variadic:   true,
		Open:       "case ",
		Close:      "",
		Separator:  ",",
		Parameters: []string{"cases"},
	},
	{
		Name:       "Sel",
		Desc:       "renders a chain of selectors separated by periods.",
		Variadic:   true,
		Open:       "",
		Close:      "",
		Separator:  ".",
		Parameters: []string{"selectors"},
	},
	{
		Name:       "Append",
		Desc:       "renders the append built-in function.",
		Variadic:   true,
		Open:       "append(",
		Close:      ")",
		Separator:  ",",
		Parameters: []string{"args"},
	},
	{
		Name:       "Cap",
		Desc:       "renders the cap built-in function.",
		Variadic:   false,
		Open:       "cap(",
		Close:      ")",
		Separator:  ",",
		Parameters: []string{"v"},
	},
	{
		Name:       "Close",
		Desc:       "renders the close built-in function.",
		Variadic:   false,
		Open:       "close(",
		Close:      ")",
		Separator:  ",",
		Parameters: []string{"c"},
	},
	{
		Name:       "Complex",
		Desc:       "renders the complex built-in function.",
		Variadic:   false,
		Open:       "complex(",
		Close:      ")",
		Separator:  ",",
		Parameters: []string{"r", "i"},
	},
	{
		Name:       "Copy",
		Desc:       "renders the copy built-in function.",
		Variadic:   false,
		Open:       "copy(",
		Close:      ")",
		Separator:  ",",
		Parameters: []string{"dst", "src"},
	},
	{
		Name:       "Delete",
		Desc:       "renders the delete built-in function.",
		Variadic:   false,
		Open:       "delete(",
		Close:      ")",
		Separator:  ",",
		Parameters: []string{"m", "key"},
	},
	{
		Name:       "Imag",
		Desc:       "renders the imag built-in function.",
		Variadic:   false,
		Open:       "imag(",
		Close:      ")",
		Separator:  ",",
		Parameters: []string{"c"},
	},
	{
		Name:       "Len",
		Desc:       "renders the len built-in function.",
		Variadic:   false,
		Open:       "len(",
		Close:      ")",
		Separator:  ",",
		Parameters: []string{"v"},
	},
	{
		Name:        "Make",
		Desc:        "renders the make built-in function.",
		Variadic:    true,
		Open:        "make(",
		Close:       ")",
		Separator:   ",",
		Parameters:  []string{"args"},
		PreventFunc: true,
	},
	{
		Name:       "New",
		Desc:       "renders the new built-in function.",
		Variadic:   false,
		Open:       "new(",
		Close:      ")",
		Separator:  ",",
		Parameters: []string{"typ"},
	},
	{
		Name:       "Panic",
		Desc:       "renders the panic built-in function.",
		Variadic:   false,
		Open:       "panic(",
		Close:      ")",
		Separator:  ",",
		Parameters: []string{"v"},
	},
	{
		Name:       "Print",
		Desc:       "renders the print built-in function.",
		Variadic:   true,
		Open:       "print(",
		Close:      ")",
		Separator:  ",",
		Parameters: []string{"args"},
	},
	{
		Name:       "Println",
		Desc:       "renders the println built-in function.",
		Variadic:   true,
		Open:       "println(",
		Close:      ")",
		Separator:  ",",
		Parameters: []string{"args"},
	},
	{
		Name:       "Real",
		Desc:       "renders the real built-in function.",
		Variadic:   false,
		Open:       "real(",
		Close:      ")",
		Separator:  ",",
		Parameters: []string{"c"},
	},
	{
		Name:       "Recover",
		Desc:       "renders the recover built-in function.",
		Variadic:   false,
		Open:       "recover(",
		Close:      ")",
		Separator:  ",",
		Parameters: []string{},
	},
}
