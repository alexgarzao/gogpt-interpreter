package parser

import (
	"errors"

	"github.com/alexgarzao/gogpt-interpreter/pkg/domain"
	"github.com/alexgarzao/gogpt-interpreter/pkg/domain/entities/bytecode"
	"github.com/alexgarzao/gogpt-interpreter/pkg/domain/entities/cp"
	"github.com/alexgarzao/gogpt-interpreter/pkg/domain/entities/symboltable"
	"github.com/alexgarzao/gogpt-interpreter/pkg/domain/usecases/instructions"
	"github.com/alexgarzao/gogpt-interpreter/pkg/domain/usecases/lexer"
)

// Parser keeps data about the whole parser process.
type Parser struct {
	l         *lexer.Lexer
	cp        *cp.CP
	bc        *bytecode.Bytecode
	symbol    *symboltable.SymbolTable
	argsCount int
}

// New creates a new Parser.
func New(l *lexer.Lexer) *Parser {
	return &Parser{
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
func (p *Parser) Parser() error {
	if err := p.parserAlgorithmDeclaration(); err != nil {
		return err
	}

	if err := p.parserVarDeclBlock(); err != nil {
		return err
	}

	if err := p.parserStmBlock(); err != nil {
		return errors.New("Expected StmtBlock")
	}

	return nil
}

// declaracao_algoritmo
//     : "algoritmo"
//       T_IDENTIFICADOR
//       ";"
//     ;
func (p *Parser) parserAlgorithmDeclaration() error {
	if p.l.GetNextTokenIf(lexer.ALGORITHM) == nil {
		return errors.New("Expected ALGORITHM")
	}

	if p.l.GetNextTokenIf(lexer.IDENT) == nil {
		return errors.New("Expected IDENT")
	}

	if p.l.GetNextTokenIf(lexer.SEMICOLON) == nil {
		return errors.New("Expected SEMICOLON")
	}

	return nil
}

// var_decl_block
//     : "variáveis"
//       (var_decl ";")+
//       "fim-variáveis"
//     ;
func (p *Parser) parserVarDeclBlock() error {
	if p.l.GetNextTokenIf(lexer.VARSBEGIN) == nil {
		return nil
	}

	var err error

	for {
		err = p.parserVarDecl()
		if err != nil {
			break
		}
	}

	if err != nil && err != domain.NotParsed {
		return err
	}

	if p.l.GetNextTokenIf(lexer.VARSEND) == nil {
		return errors.New("Expected FIM-VARIÁVEIS")
	}

	return nil
}

// var_decl
//     : T_IDENTIFICADOR ("," T_IDENTIFICADOR)* ":" tp_primitivo
//     ;
func (p *Parser) parserVarDecl() error {
	varID := p.l.GetNextTokenIf(lexer.IDENT)
	if varID == nil {
		return domain.NotParsed
	}

	if p.l.GetNextTokenIf(lexer.COLON) == nil {
		return errors.New("Expected ':'")
	}

	err := p.parserPrimitiveType()
	if err != nil {
		return errors.New("Expected type definition")
	}

	if p.l.GetNextTokenIf(lexer.SEMICOLON) == nil {
		return errors.New("Expected ;")
	}

	if p.symbol.Add(varID.Value) == -1 {
		return errors.New("Duplicated variable")
	}

	return nil
}

// tp_primitivo
//     : "inteiro"
//     | "literal"
//     | "lógico"
//     ;
func (p *Parser) parserPrimitiveType() error {
	if p.l.GetNextTokenIf(lexer.INT) == nil && p.l.GetNextTokenIf(lexer.STRING) == nil {
		return errors.New("Expected primitive type")
	}

	return nil
}

// stm_block
//     : "início"
//       (stm_list)*
//       "fim"
//     ;
func (p *Parser) parserStmBlock() error {
	if p.l.GetNextTokenIf(lexer.BLOCKBEGIN) == nil {
		return errors.New("Expected block begin")
	}

	var err error

	for {
		err = p.parserStmList()
		if err != nil {
			break
		}
	}

	if err != domain.NotParsed {
		return err
	}

	if p.l.GetNextTokenIf(lexer.BLOCKEND) == nil {
		return errors.New("Expected FIM")
	}

	return nil
}

// stm_list
//     : stm_attr
//     | fcall ";"
//     | stm_ret
//     | stm_se
//     | stm_enquanto
//     | stm_para
//     ;
func (p *Parser) parserStmList() error {
	err := p.parserFunctionCall()
	if err == nil {
		// Ensure that a ";" is presented at the EOL.
		if p.l.GetNextTokenIf(lexer.SEMICOLON) == nil {
			return errors.New("Expected SEMICOLON")
		}
		return nil
	}

	err = p.parserStmAttr()
	if err == nil {
		// Ensure that a ";" is presented at the EOL.
		if p.l.GetNextTokenIf(lexer.SEMICOLON) == nil {
			return errors.New("Expected SEMICOLON")
		}
		return nil
	}

	return domain.NotParsed
}

// stm_attr
//     : T_IDENTIFICADOR
//       ":="
//       expr
//       ";"
//     ;
func (p *Parser) parserStmAttr() error {
	id, _ := p.l.GetNextsTokensIf(lexer.IDENT, lexer.ATTR)
	if id == nil {
		return domain.NotParsed
	}

	err := p.parserExpr()
	if err != nil {
		return errors.New("Expected Expr")
	}

	p.bc.Add(instructions.STV, p.symbol.Index(id.Value))

	return nil
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

func (p *Parser) parserExpr() error {
	err := p.parserFunctionCall()
	if err == nil {
		return nil
	}

	id := p.l.GetNextTokenIf(lexer.IDENT)
	if id != nil {
		p.bc.Add(instructions.LDV, p.symbol.Index(id.Value))
		return nil
	}

	token := p.l.GetNextTokenIf(lexer.STRING)
	if token != nil {
		cpIndex := p.cp.Add(token.Value)
		p.bc.Add(instructions.LDC, cpIndex)
		return nil
	}

	return domain.NotParsed
}

// fcall
//     : T_IDENTIFICADOR "(" fargs? ")"
//     ;
func (p *Parser) parserFunctionCall() error {
	token, _ := p.l.GetNextsTokensIf(lexer.IDENT, lexer.LPAREN)
	if token == nil {
		return domain.NotParsed
	}

	funcIndex := -1
	if token.Value == "imprima" {
		funcIndex = p.cp.Add("io.println")
	} else if token.Value == "leia" {
		funcIndex = p.cp.Add("io.readln")
		// } else {
		// 	return Result{false, errors.New("Undefined function name")}
	}

	err := p.parserFunctionArgs()
	if err != nil && err != domain.NotParsed {
		return err
	}

	if token.Value == "imprima" {
		argsCountIndex := p.cp.Add(p.argsCount)
		p.bc.Add(instructions.LDC, argsCountIndex)
	}

	p.bc.Add(instructions.CALL, funcIndex)

	if p.l.GetNextTokenIf(lexer.RPAREN) == nil {
		return errors.New("Expected RPAREN")
	}

	return nil
}

// fargs
//     : expr ("," expr)*
//     ;
func (p *Parser) parserFunctionArgs() error {
	p.argsCount = 0

	err := p.parserExpr()
	if err != nil {
		return err
	}

	p.argsCount++

	for {
		if p.l.GetNextTokenIf(lexer.COMMA) == nil {
			return nil
		}

		err := p.parserExpr()
		if err != nil {
			return errors.New("Expected EXPR")
		}

		p.argsCount++
	}
}

// GetCP gets the constant pool resulted of a parser.
func (p *Parser) GetCP() *cp.CP {
	return p.cp
}

// GetBC gets the bytecode resulted of a parser.
func (p *Parser) GetBC() *bytecode.Bytecode {
	return p.bc
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
