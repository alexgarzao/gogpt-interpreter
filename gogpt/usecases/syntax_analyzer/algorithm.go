package syntax

import (
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/bytecode"
	"github.com/alexgarzao/gogpt-interpreter/gogpt/entities/constant_pool"
	lexer "github.com/alexgarzao/gogpt-interpreter/gogpt/entities/lexical_analyzer"
)

type Program struct {
	cp *constant_pool.CP
	bc *bytecode.Bytecode
}

func NewProgram() *Program {
	return &Program{
		cp: constant_pool.NewCp(),
		bc: bytecode.NewBytecode(),
	}
}

func (p *Program) TryToParse(l *lexer.Lexer) bool {
	return p.parser(l)
}

// algoritmo
//     : declaracao_algoritmo (var_decl_block)? stm_block EOF
//     ;
func (p *Program) parser(l *lexer.Lexer) bool {
	if p.parserAlgorithmDeclaration(l) == false {
		return false
	}

	// if p.parserOptionalVarDecBlock(l) == false {
	// 	return false
	// }

	if p.parserStmtBlock(l) == false {
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

// declaracao_algoritmo
//     : "algoritmo" T_IDENTIFICADOR ";"
//     ;
func (p *Program) parserAlgorithmDeclaration(l *lexer.Lexer) bool {
	if l.GetNextTokenIf(lexer.ALGORITMO) == nil {
		return false
	}

	if l.GetNextTokenIf(lexer.IDENT) == nil {
		return false
	}

	if l.GetNextTokenIf(lexer.SEMICOLON) == nil {
		return false
	}

	return true
}

func (p *Program) parserStmtBlock(l *lexer.Lexer) bool {

	mb := NewMainBlock().
		SetBytecodeGenRequirements(p.cp, p.bc)

	return mb.TryToParse(l)
}

func (p *Program) GetCP() *constant_pool.CP {
	return p.cp
}

func (p *Program) GetBC() *bytecode.Bytecode {
	return p.bc
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
