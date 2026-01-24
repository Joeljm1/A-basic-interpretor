package main

import (
	"fmt"
	"strconv"
)

type (
	prefixParseFn func() Expression
	infixParseFn  func(Expression) Expression
)
type parser struct {
	l         *lexer
	currToken Token
	peekToken Token
	errors    []string

	prefixParseFns map[TokenType]prefixParseFn
	infixParseFns  map[TokenType]infixParseFn
}

func (p *parser) registerPrefix(t TokenType, fn prefixParseFn) {
	p.prefixParseFns[t] = fn
}
func (p *parser) registerInfix(t TokenType, fn infixParseFn) {
	p.infixParseFns[t] = fn
}

const (
	_ int = iota
	LOWEST
	EQUALS  // ==
	SUM     // +
	PRODUCT // *
	PREFIX  // -X or !X
	CALL    // myFunction(X)
)

var precedanceMap = map[TokenType]int{
	TokEQ:    EQUALS,
	TokPlus:  SUM,
	TokMinus: SUM,
	TokDiv:   PRODUCT,
	TokMul:   PRODUCT,
}

func NewParser(inp string) *parser {
	p := parser{
		l:              NewLexer(inp),
		errors:         []string{},
		prefixParseFns: make(map[TokenType]prefixParseFn),
		infixParseFns:  make(map[TokenType]infixParseFn),
	}
	p.registerPrefix(TokMinus, p.ParsePrefix)
	p.registerPrefix(TokVar, p.ParseIdentifier)
	p.registerPrefix(TokLPara, p.ParseGroupedExpr)
	p.registerPrefix(TokInt, p.ParseIntegerLiteral)
	p.registerPrefix(TokFloat, p.ParseFloatLiteral)

	p.registerInfix(TokPlus, p.parseInfixExpression)
	p.registerInfix(TokMinus, p.parseInfixExpression)
	p.registerInfix(TokMul, p.parseInfixExpression)
	p.registerInfix(TokDiv, p.parseInfixExpression)
	p.registerInfix(TokEQ, p.parseInfixExpression)

	p.NextTok()
	p.NextTok() // twice to fill currTok and peekTok
	return &p
}

func (p *parser) currTokIs(t TokenType) bool {
	if p.currToken.Type == t {
		return true
	}
	return false
}

func (p *parser) peekTokIs(t TokenType) bool {
	if p.peekToken.Type == t {
		return true
	}
	return false
}

func (p *parser) expectPeekTok(t TokenType) bool {
	if p.peekTokIs(t) {
		p.NextTok()
		return true
	}
	p.errors = append(p.errors, fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type))
	return false
}

func (p *parser) ParsePrefix() Expression {
	expr := PrefixExpression{
		Token:    p.currToken,
		Operator: p.currToken.Value,
	}
	p.NextTok()
	expr.Right = p.ParseExpr(PREFIX)
	return &expr
}

func (p *parser) NextTok() {
	p.currToken = p.peekToken
	p.peekToken = p.l.next()
}

func (p *parser) SkipSpace() {
	p.l.skipSpace()
}

func (p *parser) ParseIntegerLiteral() Expression {
	il := &IntegerLiteral{
		tok: p.currToken,
	}
	val, err := strconv.Atoi(p.currToken.Value)
	if err != nil {
		p.errors = append(p.errors, fmt.Sprintf("expected integral literal but go something %s", p.currToken.Value))
		return nil
	}
	il.Val = val
	return il
}

func (p *parser) ParseFloatLiteral() Expression {
	fl := &FloatLiteral{
		tok: p.currToken,
	}
	val, err := strconv.ParseFloat(p.currToken.Value, 64)
	if err != nil {
		p.errors = append(p.errors, fmt.Sprintf("expected float literal but go something %s", p.currToken.Value))
		return nil
	}
	fl.Val = val
	return fl
}

func (p *parser) ParseIdentifier() Expression {
	return &Identifier{
		Token: p.currToken,
		Val:   p.currToken.Value,
	}
}

func (p *parser) ParseGroupedExpr() Expression {
	p.NextTok()
	expr := p.ParseExpr(LOWEST)
	p.expectPeekTok(TokRPara)
	return expr
}

func (p *parser) parseInfixExpression(left Expression) Expression {
	infix := &InfixExpression{
		Token:    p.currToken,
		Left:     left,
		Operator: p.currToken.Value,
	}
	precedance := precedanceMap[p.currToken.Type]
	p.NextTok()
	infix.Right = p.ParseExpr(precedance)
	return infix
}

func (p *parser) ParseExpr(precedance int) Expression {
	prefix := p.prefixParseFns[p.currToken.Type]
	if prefix == nil {
		p.errors = append(p.errors, fmt.Sprintf("no prefix parse fn found for %s", p.currToken.Type))
		return nil
	}
	left := prefix()
	for precedance < precedanceMap[p.peekToken.Type] {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return left
		}
		p.NextTok()
		left = p.parseInfixExpression(left)
	}
	return left
}

func (p *parser) ParseProgram() *Program {
	program := &Program{}
	for p.currToken.Type != TokEOF {
		expr := p.ParseExpr(LOWEST)
		if expr != nil {
			program.Expression = append(program.Expression, expr)
		}
		p.NextTok()
		if p.currTokIs(TokNewLine) {
			p.NextTok()
		}
	}
	return program
}
