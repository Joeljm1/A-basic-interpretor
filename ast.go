package main

import "bytes"

type Node interface {
	String() string // for debugging
}

type Expression interface {
	Node
}

type IntegerLiteral struct {
	tok Token
	Val int
}

func (il *IntegerLiteral) String() string {
	return il.tok.Value
}

func (il *IntegerLiteral) ExpressionNode() {}

type FloatLiteral struct {
	tok Token
	Val float64
}

func (fl *FloatLiteral) String() string {
	return fl.tok.Value
}

func (fl *FloatLiteral) ExpressionNode() {}

type PrefixExpression struct {
	Token    Token // prefix token like ! or -
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

func (pe *PrefixExpression) ExpressionNode() {}

type Identifier struct {
	Token Token //[token.IDENT] token
	Value string
}

func (i *Identifier) String() string {
	return i.Value
}

type InfixExpression struct {
	Token    Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(ie.Operator)
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	return out.String()
}

type ExpressionStatement struct {
	Token      Token
	Expression Expression
}

func (es *ExpressionStatement) String() string {
	if es != nil {
		return es.Expression.String()
	}
	return ""
}

type Program struct {
	Expression []Expression
}

func (p *Program) String() string {
	var out bytes.Buffer
	for i := range p.Expression {
		out.WriteString(p.Expression[i].String())
	}
	return out.String()
}
