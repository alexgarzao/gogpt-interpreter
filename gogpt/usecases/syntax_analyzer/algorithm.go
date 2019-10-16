package syntax

import (
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/bytecode"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/constant_pool"
	lexer "github.com/alexgarzao/gogpt-interpreter/gogpt/entities/lexical_analyzer"
)

type Algorithm struct {
	cp *constant_pool.CP
	bc *bytecode.Bytecode
}

func NewAlgorithm() *Algorithm {
	return &Algorithm{
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
func (a *Algorithm) Parser(l *lexer.Lexer) bool {
	if a.parserAlgorithmDeclaration(l) == false {
		return false
	}

	// if p.parserOptionalVarDeclBlock(l) == false {
	// 	return false
	// }

	if a.ParserStmBlock(l) == false {
		return false
	}

	return true
}

// declaracao_algoritmo
//     : "algoritmo"
//       T_IDENTIFICADOR
//       ";"
//     ;
func (a *Algorithm) parserAlgorithmDeclaration(l *lexer.Lexer) bool {
	if l.GetNextTokenIf(lexer.ALGORITMO) == nil || l.GetNextTokenIf(lexer.IDENT) == nil || l.GetNextTokenIf(lexer.SEMICOLON) == nil {
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

func (a *Algorithm) ParserStmBlock(l *lexer.Lexer) bool {
	l.SaveBacktrackingPoint()
	if a.isValidStmBlock(l) {
		return true
	}

	l.BackTracking()
	return false
}

// stm_block
//     : "início" (stm_list)* "fim"
//     ;
func (a *Algorithm) isValidStmBlock(l *lexer.Lexer) bool {
	if l.GetNextTokenIf(lexer.INICIO) == nil {
		return false
	}

	fc := NewFunctionCall().
		SetBytecodeGenRequirements(a.cp, a.bc)

	for fc.TryToParse(l) {
		if l.GetNextTokenIf(lexer.SEMICOLON) == nil {
			return false
		}
	}

	if l.GetNextTokenIf(lexer.FIM) == nil {
		return false
	}

	return true
}

/*
stm_list
    : stm_attr
    | fcall ";"
    | stm_ret
    | stm_se
    | stm_enquanto
    | stm_para
    ;

stm_ret
    : "retorne" expr? ";"
    ;

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

fcall
    : T_IDENTIFICADOR "(" fargs? ")"
    ;

fargs
    : expr ("," expr)*
    ;

literal
    : T_STRING_LIT
    | T_INT_LIT
    | T_KW_VERDADEIRO
    | T_KW_FALSO
    ;
*/
