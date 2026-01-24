package main

import "testing"

func Test_isLetter(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		b    byte
		want bool
	}{
		{
			name: "1",
			b:    'h',
			want: true,
		},
		{
			name: "2",
			b:    '1',
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isLetter(tt.b)
			if tt.want != got {
				t.Errorf("isLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lexer_readIdentifier(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		inp  string
		want Token
	}{
		{
			name: "1",
			inp:  "hello dfsaf sa",
			want: Token{
				Type:  TokVar,
				Value: "hello",
			},
		},
		{
			name: "2",
			inp:  "hello_world sa",
			want: Token{
				Type:  TokVar,
				Value: "hello_world",
			},
		},
		{
			name: "3",
			inp:  "hello_world",
			want: Token{
				Type:  TokVar,
				Value: "hello_world",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer(tt.inp)
			got := l.readIdentifier()
			if tt.want.Type != got.Type || tt.want.Value != got.Value {
				t.Errorf("readIdentifier() = %v, want %v in %v", got, tt.want, tt.name)
			}
		})
	}
}

func Test_lexer_readNum(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		inp  string
		want Token
	}{
		{
			name: "1",
			inp:  ".012",
			want: Token{
				Type:  TokFloat,
				Value: ".012",
			},
		},
		{
			name: "2",
			inp:  "1.43",
			want: Token{
				Type:  TokFloat,
				Value: "1.43",
			},
		},
		{
			name: "3",
			inp:  "12.",
			want: Token{
				Type:  TokFloat,
				Value: "12.",
			},
		},
		{
			name: "4",
			inp:  "12",
			want: Token{
				Type:  TokInt,
				Value: "12",
			},
		},
		{
			name: "5",
			inp:  "12.1.1",
			want: Token{
				Type:  TokErr,
				Value: "12.1.",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer(tt.inp)
			got := l.readNum()
			// TODO: update the condition below to compare got with tt.want.
			if got.Type != tt.want.Type || got.Value != tt.want.Value {
				t.Errorf("readNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lexer_next(t *testing.T) {
	inp := `a=1*1+2/(1-1)/5
			b=43+12
			print (a+b)`

	l := NewLexer(inp)
	tests := []struct {
		name string
		want Token
	}{
		{
			name: "variable a",
			want: Token{Type: TokVar, Value: "a"},
		},
		{
			name: "equals",
			want: Token{Type: TokEQ, Value: "="},
		},
		{
			name: "integer 1",
			want: Token{Type: TokInt, Value: "1"},
		},
		{
			name: "multiply",
			want: Token{Type: TokMul, Value: "*"},
		},
		{
			name: "integer 1 again",
			want: Token{Type: TokInt, Value: "1"},
		},
		{
			name: "plus",
			want: Token{Type: TokPlus, Value: "+"},
		},
		{
			name: "integer 2",
			want: Token{Type: TokInt, Value: "2"},
		},
		{
			name: "divide",
			want: Token{Type: TokDiv, Value: "/"},
		},
		{
			name: "left paren",
			want: Token{Type: TokLPara, Value: "("},
		},
		{
			name: "integer 1 paren",
			want: Token{Type: TokInt, Value: "1"},
		},
		{
			name: "minus",
			want: Token{Type: TokMinus, Value: "-"},
		},
		{
			name: "integer 1 minus",
			want: Token{Type: TokInt, Value: "1"},
		},
		{
			name: "right paren",
			want: Token{Type: TokRPara, Value: ")"},
		},
		{
			name: "divide after paren",
			want: Token{Type: TokDiv, Value: "/"},
		},
		{
			name: "integer 5",
			want: Token{Type: TokInt, Value: "5"},
		},
		{
			name: "newline",
			want: Token{Type: TokNewLine, Value: "\n"},
		},
		{
			name: "variable b",
			want: Token{Type: TokVar, Value: "b"},
		},
		{
			name: "equals b",
			want: Token{Type: TokEQ, Value: "="},
		},
		{
			name: "integer 43",
			want: Token{Type: TokInt, Value: "43"},
		},
		{
			name: "plus b",
			want: Token{Type: TokPlus, Value: "+"},
		},
		{
			name: "integer 12",
			want: Token{Type: TokInt, Value: "12"},
		},
		{
			name: "newline b",
			want: Token{Type: TokNewLine, Value: "\n"},
		},
		{
			name: "print",
			want: Token{Type: TokPrint, Value: "print"},
		},
		{
			name: "left paren print",
			want: Token{Type: TokLPara, Value: "("},
		},
		{
			name: "variable a in print",
			want: Token{Type: TokVar, Value: "a"},
		},
		{
			name: "plus in print",
			want: Token{Type: TokPlus, Value: "+"},
		},
		{
			name: "variable b in print",
			want: Token{Type: TokVar, Value: "b"},
		},
		{
			name: "right paren print",
			want: Token{Type: TokRPara, Value: ")"},
		},
		{
			name: "EOF",
			want: Token{Type: TokEOF, Value: ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := l.next()
			if tt.want.Type != got.Type || tt.want.Value != got.Value {
				t.Errorf("next() = %v, want %v", got, tt.want)
			}
		})
	}
}
