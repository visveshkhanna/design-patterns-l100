package parser

type TokenType int

const (
	IDENTIFIER TokenType = iota
	STRING
	NUMBER

	AND
	OR
	EQUALS
	LESS_THAN

	LPAREN
	RPAREN

	EOF
	ILLEGAL
)

type Token struct {
	Type    TokenType
	Literal string
	Pos     int
}

func (t TokenType) String() string {
	switch t {
	case IDENTIFIER:
		return "IDENTIFIER"
	case STRING:
		return "STRING"
	case NUMBER:
		return "NUMBER"
	case AND:
		return "AND"
	case OR:
		return "OR"
	case EQUALS:
		return "EQUALS"
	case LESS_THAN:
		return "LESS_THAN"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	case EOF:
		return "EOF"
	case ILLEGAL:
		return "ILLEGAL"
	default:
		return "UNKNOWN"
	}
}
