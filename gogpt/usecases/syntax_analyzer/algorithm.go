package syntax

import (
	"errors"

	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/bytecode"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/constant_pool"
	lexer "github.com/alexgarzao/gogpt-interpreter/gogpt/entities/lexical_analyzer"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/usecases/opcodes"
)

type Algorithm struct {
	l      *lexer.Lexer
	cp     *constant_pool.CP
	bc     *bytecode.Bytecode
	symbol *SymbolTable
}

type ParserResult struct {
	Parsed bool
	Err    error
}

func NewAlgorithm(l *lexer.Lexer) *Algorithm {
	return &Algorithm{
		l:      l,
		cp:     constant_pool.NewCp(),
		bc:     bytecode.NewBytecode(),
		symbol: NewSymbolTable(),
	}
}

func (a *Algorithm) GetCP() *constant_pool.CP {
	return a.cp
}

func (a *Algorithm) GetBC() *bytecode.Bytecode {
	return a.bc
}

// algoritmo
//     : declaracao_algoritmo
//       (var_decl_block)?
//       stm_block
//       EOF
//     ;
func (a *Algorithm) Parser() ParserResult {
	pr := a.ParserAlgorithmDeclaration()
	if pr.Parsed == false {
		return pr
	}

	pr = a.ParserVarDeclBlock()
	if pr.Err != nil {
		return pr
	}

	pr = a.ParserStmBlock()
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
func (a *Algorithm) ParserAlgorithmDeclaration() ParserResult {
	if a.l.GetNextTokenIf(lexer.ALGORITMO) == nil {
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
func (a *Algorithm) ParserVarDeclBlock() ParserResult {
	if a.l.GetNextTokenIf(lexer.VARIAVEIS) == nil {
		return ParserResult{false, nil}
	}

	pr := a.ParserVarDecl()
	for ; pr.Parsed; pr = a.ParserVarDecl() {
	}

	if pr.Parsed == false && pr.Err != nil {
		return pr
	}

	if a.l.GetNextTokenIf(lexer.FIMVARIAVEIS) == nil {
		return ParserResult{false, errors.New("Expected FIM-VARIÁVEIS")}
	}

	return ParserResult{true, nil}
}

// var_decl
//     : T_IDENTIFICADOR ("," T_IDENTIFICADOR)* ":" tp_primitivo
//     ;
func (a *Algorithm) ParserVarDecl() ParserResult {
	varId := a.l.GetNextTokenIf(lexer.IDENT)
	if varId == nil {
		return ParserResult{false, nil}
	}

	if a.l.GetNextTokenIf(lexer.COLON) == nil {
		return ParserResult{false, errors.New("Expected :")}
	}

	pr := a.ParserPrimitiveType()
	if pr.Parsed == false {
		return ParserResult{false, errors.New("Expected type definition")}
	}

	if a.l.GetNextTokenIf(lexer.SEMICOLON) == nil {
		return ParserResult{false, errors.New("Expected ;")}
	}

	if a.symbol.Add(varId.Value) == -1 {
		return ParserResult{false, errors.New("Duplicated variable")}
	}

	return ParserResult{true, nil}
}

// tp_primitivo
//     : "inteiro"
//     | "literal"
//     | "lógico"
//     ;
func (a *Algorithm) ParserPrimitiveType() ParserResult {
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
func (a *Algorithm) ParserStmBlock() ParserResult {
	if a.l.GetNextTokenIf(lexer.INICIO) == nil {
		return ParserResult{false, nil}
	}

	pr := a.ParserStmList()
	for ; pr.Parsed; pr = a.ParserStmList() {
	}

	if pr.Parsed == false && pr.Err != nil {
		return pr
	}

	if a.l.GetNextTokenIf(lexer.FIM) == nil {
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
func (a *Algorithm) ParserStmList() ParserResult {
	pr := a.ParserFunctionCall()
	if pr.Parsed == true {
		// Ensure that a ";" is presented at the EOL.
		if a.l.GetNextTokenIf(lexer.SEMICOLON) == nil {
			return ParserResult{false, errors.New("Expected SEMICOLON")}
		}
		return ParserResult{true, nil}
	}

	pr = a.ParserStmAttr()
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
func (a *Algorithm) ParserStmAttr() ParserResult {
	id, _ := a.l.GetNextsTokensIf(lexer.IDENT, lexer.ATTR)
	if id == nil {
		return ParserResult{false, nil}
	}

	pr := a.ParserExpr()
	if pr.Parsed == false {
		return ParserResult{false, errors.New("Expected Expr")}
	}

	a.bc.Add(opcodes.Stv, a.symbol.Index(id.Value))

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

func (a *Algorithm) ParserExpr() ParserResult {
	pr := a.ParserFunctionCall()
	if pr.Parsed == true {
		return ParserResult{true, nil}
	}

	id := a.l.GetNextTokenIf(lexer.IDENT)
	if id != nil {
		a.bc.Add(opcodes.Ldv, a.symbol.Index(id.Value))
		return ParserResult{true, nil}
	}

	token := a.l.GetNextTokenIf(lexer.STRING)
	if token != nil {
		cpIndex := a.cp.Add(token.Value)
		a.bc.Add(opcodes.Ldc, cpIndex)
		return ParserResult{true, nil}
	}

	return ParserResult{false, nil}
}

// fcall
//     : T_IDENTIFICADOR "(" fargs? ")"
//     ;
func (a *Algorithm) ParserFunctionCall() ParserResult {
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

	pr := a.ParserFunctionArgs()
	if pr.Parsed == false && pr.Err != nil {
		return pr
	}

	a.bc.Add(opcodes.Call, funcIndex)

	if a.l.GetNextTokenIf(lexer.RPAREN) == nil {
		return ParserResult{false, errors.New("Expected RPAREN")}
	}

	return ParserResult{true, nil}
}

// fargs
//     : expr ("," expr)*
//     ;
func (a *Algorithm) ParserFunctionArgs() ParserResult {
	pr := a.ParserExpr()
	if pr.Parsed == false {
		return pr
	}

	for {
		if a.l.GetNextTokenIf(lexer.COMMA) == nil {
			return ParserResult{true, nil}
		}
		pr := a.ParserExpr()
		if pr.Parsed == false {
			return ParserResult{false, errors.New("Expected EXPR")}
		}
	}
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
