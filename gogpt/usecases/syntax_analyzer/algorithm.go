package syntax

import (
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/bytecode"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/constant_pool"
	lexer "github.com/alexgarzao/gogpt-interpreter/gogpt/entities/lexical_analyzer"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/usecases/opcodes"
)

type Algorithm struct {
	l  *lexer.Lexer
	cp *constant_pool.CP
	bc *bytecode.Bytecode
}

func NewAlgorithm(l *lexer.Lexer) *Algorithm {
	return &Algorithm{
		l:  l,
		cp: constant_pool.NewCp(),
		bc: bytecode.NewBytecode(),
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
func (a *Algorithm) Parser() bool {
	if a.ParserAlgorithmDeclaration() == false {
		return false
	}

	// if p.parserOptionalVarDeclBlock() == false {
	// 	return false
	// }

	if a.ParserStmBlock() == false {
		return false
	}

	return true
}

// declaracao_algoritmo
//     : "algoritmo"
//       T_IDENTIFICADOR
//       ";"
//     ;
func (a *Algorithm) ParserAlgorithmDeclaration() bool {
	if a.l.GetNextTokenIf(lexer.ALGORITMO) == nil || a.l.GetNextTokenIf(lexer.IDENT) == nil || a.l.GetNextTokenIf(lexer.SEMICOLON) == nil {
		return false
	}

	return true
}

// var_decl_block
//     : "variáveis" (var_decl ";")+ "fim-variáveis"
//     ;
//
// var_decl
//     : T_IDENTIFICADOR ("," T_IDENTIFICADOR)* ":" tp_primitivo
//     ;
//
// tp_primitivo
//     : "inteiro"
//     | "literal"
//     | "lógico"
//     ;

// stm_block
//     : "início"
//       (stm_list)*
//       "fim"
//     ;
func (a *Algorithm) ParserStmBlock() bool {
	// TODO: SaveBack could be a method of Algorithm???
	a.l.SaveBacktrackingPoint()

	if a.isValidStmBlock() {
		return true
	}

	a.l.BackTracking()
	return false
}

func (a *Algorithm) isValidStmBlock() bool {
	if a.l.GetNextTokenIf(lexer.INICIO) == nil {
		return false
	}

	for a.ParserStmList() {
	}

	if a.l.GetNextTokenIf(lexer.FIM) == nil {
		return false
	}

	return true
}

// stm_list
//     : stm_attr
//     | fcall ";"
//     | stm_ret
//     | stm_se
//     | stm_enquanto
//     | stm_para
//     ;
//
// stm_ret
//     : "retorne" expr? ";"
//     ;
func (a *Algorithm) ParserStmList() bool {
	a.l.SaveBacktrackingPoint()

	if a.ParserFunctionCall() == true {
		// Ensure that a ";" is presented at the EOL.
		if a.l.GetNextTokenIf(lexer.SEMICOLON) != nil {
			return true
		}
	}

	a.l.BackTracking()
	return false
}

// fcall
//     : T_IDENTIFICADOR "(" fargs? ")"
//     ;
//
// fargs
//     : expr ("," expr)*
//     ;
func (a *Algorithm) ParserFunctionCall() bool {
	a.l.SaveBacktrackingPoint()
	if a.isValidFunctionCall() {
		return true
	}

	a.l.BackTracking()
	return false
}

func (a *Algorithm) isValidFunctionCall() bool {
	var token *lexer.Token
	if token = a.l.GetNextTokenIf(lexer.IDENT); token == nil {
		return false
	}

	if a.l.GetNextTokenIf(lexer.LPAREN) == nil {
		return false
	}

	funcIndex := -1
	if token.Value == "imprima" {
		funcIndex = a.cp.Add("io.println")
	}

	if token = a.l.GetNextTokenIf(lexer.STRING); token != nil {
		for {
			cpIndex := a.cp.Add(token.Value)
			a.bc.Add(opcodes.Ldc, cpIndex)
			if a.l.GetNextTokenIf(lexer.COMMA) == nil {
				break
			}
			if a.l.GetNextTokenIf(lexer.STRING) == nil {
				return false
			}
		}
	}

	a.bc.Add(opcodes.Call, funcIndex)

	if a.l.GetNextTokenIf(lexer.RPAREN) == nil {
		return false
	}

	return true
}

/*
stm_attr
    : T_IDENTIFICADOR ":=" expr ";"
    ;

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

expr
    : expr ("ou"|"||") expr
    | expr ("e"|"&&") expr
    | expr ("="|"<>") expr
    | expr (">"|">="|"<"|"<=") expr
    | expr ("+" | "-") expr
    | expr ("/"|"*") expr
    | ("+"|"-"|"não")? termo
    ;

termo
    : fcall
    | T_IDENTIFICADOR
    | literal
    | "(" expr ")"
    ;

literal
    : T_STRING_LIT
    | T_INT_LIT
    | T_KW_VERDADEIRO
    | T_KW_FALSO
    ;
*/
