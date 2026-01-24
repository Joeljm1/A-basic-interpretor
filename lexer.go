package main

import (
	"fmt"
)

// ASCII For now
type lexer struct {
	inp     string
	pos     int // current ch index
	readPos int // next char index read
	ch      byte
}

func NewLexer(inp string) *lexer {
	l := &lexer{
		inp: inp,
	}
	l.readChar()
	return l
}

func (l *lexer) readChar() {
	if l.readPos >= len(l.inp) {
		l.ch = 0
	} else {
		l.ch = l.inp[l.readPos]
	}
	l.pos = l.readPos
	l.readPos++
}

// excluding new line
func (l *lexer) skipSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(b byte) bool {
	if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') {
		return true
	}
	return false
}

func isNum(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}

func (l *lexer) readIdentifier() Token {
	if !isLetter(l.ch) && l.ch != '_' {
		panic(fmt.Sprintf("got %v", string(l.ch)))
	}
	start := l.pos
	for isLetter(l.ch) || isNum(l.ch) || l.ch == '_' {
		l.readChar()
	}
	txt := l.inp[start:l.pos]
	tok := Token{
		Value: txt,
	}
	if txt == "print" {
		tok.Type = TokPrint
	} else {
		tok.Type = TokVar
	}
	return tok
}

// returns number and if num is float
func (l *lexer) readNum() Token {
	isFloat := false
	start := l.pos
	if !isNum(l.ch) && l.ch != '.' {
		panic("Error")
	}
	for isNum(l.ch) || l.ch == '.' {
		if l.ch == '.' {
			if !isFloat {
				isFloat = true
			} else {
				return Token{
					Type:  TokErr,
					Value: l.inp[start : l.pos+1],
				}
			}
		}
		l.readChar()
	}
	var tok Token
	if isFloat {
		tok.Type = TokFloat
	} else {
		tok.Type = TokInt
	}
	tok.Value = l.inp[start:l.pos]
	return tok
}

func (l *lexer) next() Token {
	l.skipSpace()
	switch l.ch {
	case '=':
		l.readChar()
		return NewTok(TokEQ, "=")
	case '(':
		l.readChar()
		return NewTok(TokLPara, "(")
	case ')':
		l.readChar()
		return NewTok(TokRPara, ")")
	case '*':
		l.readChar()
		return NewTok(TokMul, "*")
	case '/':
		l.readChar()
		return NewTok(TokDiv, "/")
	case '-':
		l.readChar()
		return NewTok(TokMinus, "-")
	case '+':
		l.readChar()
		return NewTok(TokPlus, "+")
	case '\n':
		l.readChar()
		return NewTok(TokNewLine, "\n")
	case 0:
		return NewTok(TokEOF, "")
	default:
		if isLetter(l.ch) || l.ch == '_' {
			return l.readIdentifier()
		} else if isNum(l.ch) || l.ch == '.' {
			return l.readNum()
		} else {
			l.readChar()
			return NewTok(TokErr, string(l.ch))
		}

	}
}
