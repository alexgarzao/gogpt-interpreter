package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfUselessCharsHasBeenRemoved(t *testing.T) {
	l := NewLexer("   algoritmo\t\t  \t\t\n\t\r meuid;")
	assert.Equal(t, &Token{ALGORITMO, ALGORITMO}, l.NextToken())
	assert.Equal(t, &Token{IDENT, "meuid"}, l.NextToken())
	assert.Equal(t, &Token{SEMICOLON, SEMICOLON}, l.NextToken())
	assert.Equal(t, &Token{EOF, EOF}, l.NextToken())
}

func TestValidKeywords(t *testing.T) {
	l := NewLexer("algoritmo início fim variáveis fim-variáveis")
	assert.Equal(t, &Token{ALGORITMO, ALGORITMO}, l.NextToken())
	assert.Equal(t, &Token{INICIO, INICIO}, l.NextToken())
	assert.Equal(t, &Token{FIM, FIM}, l.NextToken())
	assert.Equal(t, &Token{VARIAVEIS, VARIAVEIS}, l.NextToken())
	assert.Equal(t, &Token{FIMVARIAVEIS, FIMVARIAVEIS}, l.NextToken())
	assert.Equal(t, &Token{EOF, EOF}, l.NextToken())
}

func TestInvalidKeywords(t *testing.T) {
	l := NewLexer("algoritmoa in ício inicio fimm afim")
	assert.Equal(t, &Token{IDENT, "algoritmoa"}, l.NextToken())
	assert.Equal(t, &Token{IDENT, "in"}, l.NextToken())
	assert.Equal(t, &Token{INVALID, ""}, l.NextToken())
	assert.Equal(t, &Token{IDENT, "cio"}, l.NextToken())
	assert.Equal(t, &Token{IDENT, "inicio"}, l.NextToken())
	assert.Equal(t, &Token{IDENT, "fimm"}, l.NextToken())
	assert.Equal(t, &Token{IDENT, "afim"}, l.NextToken())
	assert.Equal(t, &Token{EOF, EOF}, l.NextToken())
}

func TestDelimiters(t *testing.T) {
	l := NewLexer("(,:)( ());):= ) := ")
	assert.Equal(t, &Token{LPAREN, LPAREN}, l.NextToken())
	assert.Equal(t, &Token{COMMA, COMMA}, l.NextToken())
	assert.Equal(t, &Token{COLON, COLON}, l.NextToken())
	assert.Equal(t, &Token{RPAREN, RPAREN}, l.NextToken())
	assert.Equal(t, &Token{LPAREN, LPAREN}, l.NextToken())
	assert.Equal(t, &Token{LPAREN, LPAREN}, l.NextToken())
	assert.Equal(t, &Token{RPAREN, RPAREN}, l.NextToken())
	assert.Equal(t, &Token{RPAREN, RPAREN}, l.NextToken())
	assert.Equal(t, &Token{SEMICOLON, SEMICOLON}, l.NextToken())
	assert.Equal(t, &Token{RPAREN, RPAREN}, l.NextToken())
	assert.Equal(t, &Token{ATTR, ATTR}, l.NextToken())
	assert.Equal(t, &Token{RPAREN, RPAREN}, l.NextToken())
	assert.Equal(t, &Token{ATTR, ATTR}, l.NextToken())
	assert.Equal(t, &Token{EOF, EOF}, l.NextToken())
}

func TestValidIDs(t *testing.T) {
	l := NewLexer("i ix ir2 if_ _Id Ix iX Camel_Case NotCamelCase ComAcentuação")
	assert.Equal(t, &Token{IDENT, "i"}, l.NextToken())
	assert.Equal(t, &Token{IDENT, "ix"}, l.NextToken())
	assert.Equal(t, &Token{IDENT, "ir2"}, l.NextToken())
	assert.Equal(t, &Token{IDENT, "if_"}, l.NextToken())
	assert.Equal(t, &Token{IDENT, "_Id"}, l.NextToken())
	assert.Equal(t, &Token{IDENT, "Ix"}, l.NextToken())
	assert.Equal(t, &Token{IDENT, "iX"}, l.NextToken())
	assert.Equal(t, &Token{IDENT, "Camel_Case"}, l.NextToken())
	assert.Equal(t, &Token{IDENT, "NotCamelCase"}, l.NextToken())
	assert.Equal(t, &Token{IDENT, "ComAcentuação"}, l.NextToken())
	assert.Equal(t, &Token{EOF, EOF}, l.NextToken())
}

func TestValidStrings(t *testing.T) {
	l := NewLexer(`i "" "a" "aaaaaa" "bb bbbb bbbb" "ccc cccc c  " "  ccc   "`)
	assert.Equal(t, &Token{IDENT, "i"}, l.NextToken())
	assert.Equal(t, &Token{STRING, ""}, l.NextToken())
	assert.Equal(t, &Token{STRING, "a"}, l.NextToken())
	assert.Equal(t, &Token{STRING, "aaaaaa"}, l.NextToken())
	assert.Equal(t, &Token{STRING, "bb bbbb bbbb"}, l.NextToken())
	assert.Equal(t, &Token{STRING, "ccc cccc c  "}, l.NextToken())
	assert.Equal(t, &Token{STRING, "  ccc   "}, l.NextToken())
	assert.Equal(t, &Token{EOF, EOF}, l.NextToken())
}

func TestInvalidStrings(t *testing.T) {
	l := NewLexer("\"a\t\" \"aaa\nbbb\" \"cc \"dddd \r eeee\" \"\"\"\"")
	assert.Equal(t, &Token{STRING, "a\t"}, l.NextToken())
	assert.Equal(t, &Token{INVALID, "\"aaa"}, l.NextToken())
	assert.Equal(t, &Token{IDENT, "bbb"}, l.NextToken())
	assert.Equal(t, &Token{STRING, " "}, l.NextToken())
	assert.Equal(t, &Token{IDENT, "cc"}, l.NextToken())
	assert.Equal(t, &Token{INVALID, "\"dddd "}, l.NextToken())
	assert.Equal(t, &Token{IDENT, "eeee"}, l.NextToken())
	assert.Equal(t, &Token{STRING, " "}, l.NextToken())
	assert.Equal(t, &Token{STRING, ""}, l.NextToken())
	assert.Equal(t, &Token{INVALID, "\""}, l.NextToken())
	assert.Equal(t, &Token{EOF, EOF}, l.NextToken())
}

func TestValidInts(t *testing.T) {
	l := NewLexer(`1 123 0123 123456789 1230`)
	assert.Equal(t, &Token{INT, "1"}, l.NextToken())
	assert.Equal(t, &Token{INT, "123"}, l.NextToken())
	assert.Equal(t, &Token{INT, "123"}, l.NextToken())
	assert.Equal(t, &Token{INT, "123456789"}, l.NextToken())
	assert.Equal(t, &Token{INT, "1230"}, l.NextToken())
	assert.Equal(t, &Token{EOF, EOF}, l.NextToken())
}

func TestInvalidInts(t *testing.T) {
	l := NewLexer("1a 1b23 c123 123\"\" 12\r .89 30 4\n567")
	assert.Equal(t, &Token{INVALID, "1a"}, l.NextToken())
	assert.Equal(t, &Token{INVALID, "1b"}, l.NextToken())
	assert.Equal(t, &Token{INT, "23"}, l.NextToken())
	assert.Equal(t, &Token{IDENT, "c123"}, l.NextToken())
	assert.Equal(t, &Token{INVALID, "123\""}, l.NextToken())
	assert.Equal(t, &Token{INVALID, "\" 12"}, l.NextToken())
	assert.Equal(t, &Token{INVALID, ""}, l.NextToken())
	assert.Equal(t, &Token{INT, "89"}, l.NextToken())
	assert.Equal(t, &Token{INT, "30"}, l.NextToken())
	assert.Equal(t, &Token{INT, "4"}, l.NextToken())
	assert.Equal(t, &Token{INT, "567"}, l.NextToken())
	assert.Equal(t, &Token{EOF, EOF}, l.NextToken())
}

func TestValidTokensWithoutPontuations(t *testing.T) {
	l := NewLexer("algoritmo meuid; imprima fim(),")
	assert.Equal(t, &Token{ALGORITMO, ALGORITMO}, l.NextToken())
	assert.Equal(t, &Token{IDENT, "meuid"}, l.NextToken())
	assert.Equal(t, &Token{SEMICOLON, SEMICOLON}, l.NextToken())
	assert.Equal(t, &Token{IDENT, "imprima"}, l.NextToken())
	assert.Equal(t, &Token{FIM, FIM}, l.NextToken())
	assert.Equal(t, &Token{LPAREN, LPAREN}, l.NextToken())
	assert.Equal(t, &Token{RPAREN, RPAREN}, l.NextToken())
	assert.Equal(t, &Token{COMMA, COMMA}, l.NextToken())
	assert.Equal(t, &Token{EOF, EOF}, l.NextToken())
}

func TestTokensWithPontuations(t *testing.T) {
	l := NewLexer("algoritmo olá; início")
	assert.Equal(t, &Token{ALGORITMO, ALGORITMO}, l.NextToken())
	assert.Equal(t, &Token{IDENT, "olá"}, l.NextToken())
	assert.Equal(t, &Token{SEMICOLON, ";"}, l.NextToken())
	assert.Equal(t, &Token{INICIO, "INÍCIO"}, l.NextToken())
	assert.Equal(t, &Token{EOF, EOF}, l.NextToken())
}

func TestLiteralsOrIdentifiersTokens(t *testing.T) {
	l := NewLexer(`olá "oi !" 123`)
	assert.Equal(t, &Token{IDENT, "olá"}, l.NextToken())
	assert.Equal(t, &Token{STRING, "oi !"}, l.NextToken())
	assert.Equal(t, &Token{INT, "123"}, l.NextToken())
	assert.Equal(t, &Token{EOF, EOF}, l.NextToken())
}

func TestValidHelloWorld(t *testing.T) {
	c :=
		`algoritmo olá_mundo;
início
	imprima("Olá mundo!");
fim`

	l := NewLexer(c)
	assert.Equal(t, &Token{ALGORITMO, "ALGORITMO"}, l.NextToken())
	assert.Equal(t, &Token{IDENT, "olá_mundo"}, l.NextToken())
	assert.Equal(t, &Token{SEMICOLON, ";"}, l.NextToken())
	assert.Equal(t, &Token{INICIO, "INÍCIO"}, l.NextToken())
	assert.Equal(t, &Token{IDENT, "imprima"}, l.NextToken())
	assert.Equal(t, &Token{LPAREN, "("}, l.NextToken())
	assert.Equal(t, &Token{STRING, "Olá mundo!"}, l.NextToken())
	assert.Equal(t, &Token{RPAREN, ")"}, l.NextToken())
	assert.Equal(t, &Token{SEMICOLON, ";"}, l.NextToken())
	assert.Equal(t, &Token{FIM, "FIM"}, l.NextToken())
	assert.Equal(t, &Token{EOF, EOF}, l.NextToken())
}

func TestGetIf(t *testing.T) {
	l := NewLexer("algoritmo início ( xxx ) fim")
	assert.Nil(t, l.GetNextTokenIf(IDENT))
	assert.Equal(t, &Token{ALGORITMO, ALGORITMO}, l.GetNextTokenIf(ALGORITMO))
	assert.Equal(t, &Token{INICIO, INICIO}, l.GetNextTokenIf(INICIO))
	assert.Equal(t, &Token{LPAREN, LPAREN}, l.GetNextTokenIf(LPAREN))
	assert.Equal(t, &Token{IDENT, "xxx"}, l.GetNextTokenIf(IDENT))
	assert.Equal(t, &Token{RPAREN, RPAREN}, l.GetNextTokenIf(RPAREN))
	assert.Equal(t, &Token{FIM, FIM}, l.GetNextTokenIf(FIM))
	assert.Equal(t, &Token{EOF, EOF}, l.GetNextTokenIf(EOF))
}
