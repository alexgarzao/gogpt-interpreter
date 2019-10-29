package parser

import (
	"errors"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/bytecode"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/cp"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/model/symboltable"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/instructions"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/services/lexer"
)

// Algorithm keeps data about the whole parser process.
type Algorithm struct {
	l         *lexer.Lexer
	cp        *cp.CP
	bc        *bytecode.Bytecode
	symbol    *symboltable.SymbolTable
	argsCount int
}

// ParserResult keeps data about the parser result.
type ParserResult struct {
	Parsed bool
	Err    error
}

// New creates a new Algorithm.
func New(l *lexer.Lexer) *Algorithm {
	return &Algorithm{
		l:      l,
		cp:     cp.New(),
		bc:     bytecode.New(),
		symbol: symboltable.New(),
	}
}

// Parser is responsible for parser an algorithm.
//
// algoritmo
//     : declaracao_algoritmo
//       (var_decl_block)?
//       stm_block
//       EOF
//     ;
func (a *Algorithm) Parser() ParserResult {
	pr := a.parserAlgorithmDeclaration()
	if pr.Parsed == false {
		return pr
	}

	pr = a.parserVarDeclBlock()
	if pr.Err != nil {
		return pr
	}

	pr = a.parserStmBlock()
	if pr.Parsed == false {
		return ParserResult{false, errors.New("Expected StmtBlock")}
	}

	return ParserResult{true, nil}
}

// declaracao_algoritmo
//     : "algoritmo"
//       T_IDENTIFICADOR
//       ";"
//     ;
func (a *Algorithm) parserAlgorithmDeclaration() ParserResult {
	if a.l.GetNextTokenIf(lexer.ALGORITHM) == nil {
		return ParserResult{false, nil}
	}

	if a.l.GetNextTokenIf(lexer.IDENT) == nil {
		return ParserResult{false, errors.New("Expected IDENT")}
	}

	if a.l.GetNextTokenIf(lexer.SEMICOLON) == nil {
		return ParserResult{false, errors.New("Expected SEMICOLON")}
	}

	return ParserResult{true, nil}
}

// var_decl_block
//     : "variáveis"
//       (var_decl ";")+
//       "fim-variáveis"
//     ;
func (a *Algorithm) parserVarDeclBlock() ParserResult {
	if a.l.GetNextTokenIf(lexer.VARSBEGIN) == nil {
		return ParserResult{false, nil}
	}

	pr := a.parserVarDecl()
	for ; pr.Parsed; pr = a.parserVarDecl() {
	}

	if pr.Parsed == false && pr.Err != nil {
		return pr
	}

	if a.l.GetNextTokenIf(lexer.VARSEND) == nil {
		return ParserResult{false, errors.New("Expected FIM-VARIÁVEIS")}
	}

	return ParserResult{true, nil}
}

// var_decl
//     : T_IDENTIFICADOR ("," T_IDENTIFICADOR)* ":" tp_primitivo
//     ;
func (a *Algorithm) parserVarDecl() ParserResult {
	varID := a.l.GetNextTokenIf(lexer.IDENT)
	if varID == nil {
		return ParserResult{false, nil}
	}

	if a.l.GetNextTokenIf(lexer.COLON) == nil {
		return ParserResult{false, errors.New("Expected ':'")}
	}

	pr := a.parserPrimitiveType()
	if pr.Parsed == false {
		return ParserResult{false, errors.New("Expected type definition")}
	}

	if a.l.GetNextTokenIf(lexer.SEMICOLON) == nil {
		return ParserResult{false, errors.New("Expected ;")}
	}

	if a.symbol.Add(varID.Value) == -1 {
		return ParserResult{false, errors.New("Duplicated variable")}
	}

	return ParserResult{true, nil}
}

// tp_primitivo
//     : "inteiro"
//     | "literal"
//     | "lógico"
//     ;
func (a *Algorithm) parserPrimitiveType() ParserResult {
	if a.l.GetNextTokenIf(lexer.INT) != nil || a.l.GetNextTokenIf(lexer.STRING) != nil {
		return ParserResult{true, nil}
	}

	return ParserResult{false, nil}
}

// stm_block
//     : "início"
//       (stm_list)*
//       "fim"
//     ;
func (a *Algorithm) parserStmBlock() ParserResult {
	if a.l.GetNextTokenIf(lexer.BLOCKBEGIN) == nil {
		return ParserResult{false, nil}
	}

	pr := a.parserStmList()
	for ; pr.Parsed; pr = a.parserStmList() {
	}

	if pr.Parsed == false && pr.Err != nil {
		return pr
	}

	if a.l.GetNextTokenIf(lexer.BLOCKEND) == nil {
		return ParserResult{false, errors.New("Expected FIM")}
	}

	return ParserResult{true, nil}
}

// stm_list
//     : stm_attr
//     | fcall ";"
//     | stm_ret
//     | stm_se
//     | stm_enquanto
//     | stm_para
//     ;
func (a *Algorithm) parserStmList() ParserResult {
	pr := a.parserFunctionCall()
	if pr.Parsed == true {
		// Ensure that a ";" is presented at the EOL.
		if a.l.GetNextTokenIf(lexer.SEMICOLON) == nil {
			return ParserResult{false, errors.New("Expected SEMICOLON")}
		}
		return ParserResult{true, nil}
	}

	pr = a.parserStmAttr()
	if pr.Parsed == true {
		// Ensure that a ";" is presented at the EOL.
		if a.l.GetNextTokenIf(lexer.SEMICOLON) == nil {
			return ParserResult{false, errors.New("Expected SEMICOLON")}
		}
		return ParserResult{true, nil}
	}

	return ParserResult{false, nil}
}

// stm_attr
//     : T_IDENTIFICADOR
//       ":="
//       expr
//       ";"
//     ;
func (a *Algorithm) parserStmAttr() ParserResult {
	id, _ := a.l.GetNextsTokensIf(lexer.IDENT, lexer.ATTR)
	if id == nil {
		return ParserResult{false, nil}
	}

	pr := a.parserExpr()
	if pr.Parsed == false {
		return ParserResult{false, errors.New("Expected Expr")}
	}

	a.bc.Add(instructions.STV, a.symbol.Index(id.Value))

	return ParserResult{true, nil}
}

// expr
//     : expr ("ou"|"||") expr
//     | expr ("e"|"&&") expr
//     | expr ("="|"<>") expr
//     | expr (">"|">="|"<"|"<=") expr
//     | expr ("+" | "-") expr
//     | expr ("/"|"*") expr
//     | ("+"|"-"|"não")? termo
//     ;

// termo
//     : fcall
//     | T_IDENTIFICADOR
//     | literal
//     | "(" expr ")"
//     ;

func (a *Algorithm) parserExpr() ParserResult {
	pr := a.parserFunctionCall()
	if pr.Parsed == true {
		return ParserResult{true, nil}
	}

	id := a.l.GetNextTokenIf(lexer.IDENT)
	if id != nil {
		a.bc.Add(instructions.LDV, a.symbol.Index(id.Value))
		return ParserResult{true, nil}
	}

	token := a.l.GetNextTokenIf(lexer.STRING)
	if token != nil {
		cpIndex := a.cp.Add(token.Value)
		a.bc.Add(instructions.LDC, cpIndex)
		return ParserResult{true, nil}
	}

	return ParserResult{false, nil}
}

// fcall
//     : T_IDENTIFICADOR "(" fargs? ")"
//     ;
func (a *Algorithm) parserFunctionCall() ParserResult {
	token, _ := a.l.GetNextsTokensIf(lexer.IDENT, lexer.LPAREN)
	if token == nil {
		return ParserResult{false, nil}
	}

	funcIndex := -1
	if token.Value == "imprima" {
		funcIndex = a.cp.Add("io.println")
	} else if token.Value == "leia" {
		funcIndex = a.cp.Add("io.readln")
		// } else {
		// 	return ParserResult{false, errors.New("Undefined function name")}
	}

	pr := a.parserFunctionArgs()
	if pr.Parsed == false && pr.Err != nil {
		return pr
	}

	if token.Value == "imprima" {
		argsCountIndex := a.cp.Add(a.argsCount)
		a.bc.Add(instructions.LDC, argsCountIndex)
	}

	a.bc.Add(instructions.CALL, funcIndex)

	if a.l.GetNextTokenIf(lexer.RPAREN) == nil {
		return ParserResult{false, errors.New("Expected RPAREN")}
	}

	return ParserResult{true, nil}
}

// fargs
//     : expr ("," expr)*
//     ;
func (a *Algorithm) parserFunctionArgs() ParserResult {
	a.argsCount = 0

	pr := a.parserExpr()
	if pr.Parsed == false {
		return pr
	}

	a.argsCount++

	for {
		if a.l.GetNextTokenIf(lexer.COMMA) == nil {
			return ParserResult{true, nil}
		}
		pr := a.parserExpr()
		if pr.Parsed == false {
			return ParserResult{false, errors.New("Expected EXPR")}
		}

		a.argsCount++
	}
}

// GetCP gets the constant pool resulted of a parser.
func (a *Algorithm) GetCP() *cp.CP {
	return a.cp
}

// GetBC gets the bytecode resulted of a parser.
func (a *Algorithm) GetBC() *bytecode.Bytecode {
	return a.bc
}

// stm_ret
//     : "retorne" expr? ";"
//     ;
// TODO

/*
stm_se
    : "se" expr "então" stm_list ("senão" stm_list)? "fim-se"
    ;

stm_enquanto
    : "enquanto" expr "faça" stm_list "fim-enquanto"
    ;

stm_para
    : "para" T_IDENTIFICADOR "de" expr "até" expr passo? "faça" stm_list "fim-para"
    ;

passo
    : "passo" ("+"|"-")? T_INT_LIT
    ;

literal
    : T_STRING_LIT
    | T_INT_LIT
    | T_KW_VERDADEIRO
    | T_KW_FALSO
    ;
*/
