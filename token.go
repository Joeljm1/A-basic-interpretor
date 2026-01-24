package main

type TokenType int

const (
	TokErr TokenType = iota
	TokEOF
	TokInt
	TokFloat
	TokVar
	TokPlus
	TokMinus
	TokDiv
	TokMul
	TokPrint
	TokEQ
	TokLPara
	TokRPara
	TokNewLine
)

func (t TokenType) String() string {
	switch t {
	case TokErr:
		return "Error "
	case TokEOF:
		return "EOF"
	case TokInt:
		return "Number"
	case TokFloat:
		return "Float"
	case TokVar:
		return "Variable"
	case TokPlus:
		return "Plus"
	case TokMinus:
		return "Minus"
	case TokDiv:
		return "Div"
	case TokMul:
		return "Mul"
	case TokEQ:
		return "EQ"
	case TokLPara:
		return "LPara"
	case TokRPara:
		return "RPara"
	case TokPrint:
		return "Print"
	case TokNewLine:
		return "NewLine"
	default:
		return "Invalid Token type"
	}
}

type Token struct {
	Type  TokenType
	Value string
}

func NewTok(ty TokenType, val string) Token {
	return Token{
		Type:  ty,
		Value: val,
	}
}
