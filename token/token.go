package token

type TokenType string

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"

	// identifiers + literals
	IDENT     = "IDENT"
	INT       = "INT"
	STRING 	  = "STRING"

	// Operators
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	ASTERISK  = "*"
	SLASH     = "/"

	GT        = ">"
	LT        = "<"
	EQ        = "=="
	NOT_EQ	  = "!="

	// Delimiters
	COMMA     = ","
	COLON 	  = ":"
	SEMICOLON = ";"

	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"

	// keywords
	FUNCTION  = "FUNCTION"
	LET       = "LET"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
	IF        = "IF"
	ELSE      = "ELSE"
	RETURN    = "RETURN"
)

type Token struct {
	Type TokenType
	Literal string
}

var keywords = map[string]TokenType {
	"fn"  	 : FUNCTION,
	"let" 	 : LET,
	"true" 	 : TRUE,
	"false"  : FALSE,
	"if" 	 : IF,
	"else" 	 : ELSE,
	"return" : RETURN,
}

func LookupIdentifier(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}