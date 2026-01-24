package main

import (
	"bytes"
	"fmt"
)

var variables map[string]float64 = map[string]float64{}

type Node interface {
	String() string // for debugging
}

type Expression interface {
	Node
	Value() float64 // for now doing only float
}

type IntegerLiteral struct {
	tok Token
	Val int
}

func (il *IntegerLiteral) Value() float64 {
	return float64(il.Val)
}

func (il *IntegerLiteral) String() string {
	return il.tok.Value
}

type FloatLiteral struct {
	tok Token
	Val float64
}

func (fl *FloatLiteral) String() string {
	return fl.tok.Value
}
func (fl *FloatLiteral) Value() float64 {
	return fl.Val
}

type PrefixExpression struct {
	Token    Token // prefix token like -
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) Value() float64 {
	if pe.Operator == "-" {
		return -1 * pe.Right.Value()
	}
	panic(fmt.Sprintf("prefix non minus value found : %s", pe.Operator))
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

type Identifier struct {
	Token Token //[token.IDENT] token
	Val   string
}

func (i *Identifier) String() string {
	return i.Val
}

func (pe *Identifier) Value() float64 {
	val, ok := variables[pe.Val]
	if !ok {
		panic(fmt.Sprintf("variable %v not declared ", pe.Val))
	}
	return val
}

// func (pe *Identifier) Value() {}

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

func (ie *InfixExpression) Value() float64 {
	switch ie.Operator {
	case "+":
		return ie.Left.Value() + ie.Right.Value()
	case "-":
		return ie.Left.Value() - ie.Right.Value()
	case "*":
		return ie.Left.Value() * ie.Right.Value()
	case "/":
		return ie.Left.Value() / ie.Right.Value()
	case "=":
		ident, ok := ie.Left.(*Identifier)
		if !ok {
			panic("assigning to non identifier ")
		}
		val := ie.Right.Value()
		variables[ident.Val] = val
		return val
	default:
		panic(fmt.Sprintf("Unknown operator %s", ie.Operator))
	}
}

type PrintExpr struct {
	Token Token
	Val   Expression
}

func (pe *PrintExpr) Value() float64 {
	return pe.Val.Value()
}

func (pe *PrintExpr) String() string {
	var out bytes.Buffer
	out.WriteString("print ")
	out.WriteString(pe.Val.String())
	return out.String()
}

type Program struct {
	Expression []Expression
}

func (p *Program) Value() float64 {
	if len(p.Expression) > 0 {
		return p.Expression[len(p.Expression)-1].Value()
	}
	panic("Empty program")
}

func (p *Program) String() string {
	var out bytes.Buffer
	for i := range p.Expression {
		out.WriteString(p.Expression[i].String())
	}
	return out.String()
}
