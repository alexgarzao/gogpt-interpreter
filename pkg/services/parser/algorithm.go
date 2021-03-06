package parser

import (
	"errors"

	"github.com/alexgarzao/gogpt-interpreter/pkg/model/bytecode"
	"github.com/alexgarzao/gogpt-interpreter/pkg/model/cp"
	"github.com/alexgarzao/gogpt-interpreter/pkg/model/symboltable"
	"github.com/alexgarzao/gogpt-interpreter/pkg/services/instructions"
	"github.com/alexgarzao/gogpt-interpreter/pkg/services/lexer"
)

// Parser keeps data about the whole parser process.
type Parser struct {
	l         *lexer.Lexer
	cp        *cp.CP
	bc        *bytecode.Bytecode
	symbol    *symboltable.SymbolTable
	argsCount int
}

// Result keeps data about the parser result.
type Result struct {
	Parsed bool
	Err    error
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
func (p *Parser) Parser() Result {
	pr := p.parserAlgorithmDeclaration()
	if pr.Parsed == false {
		return pr
	}

	pr = p.parserVarDeclBlock()
	if pr.Err != nil {
		return pr
	}

	pr = p.parserStmBlock()
	if pr.Parsed == false {
		return Result{false, errors.New("Expected StmtBlock")}
	}

	return Result{true, nil}
}

// declaracao_algoritmo
//     : "algoritmo"
//       T_IDENTIFICADOR
//       ";"
//     ;
func (p *Parser) parserAlgorithmDeclaration() Result {
	if p.l.GetNextTokenIf(lexer.ALGORITHM) == nil {
		return Result{false, nil}
	}

	if p.l.GetNextTokenIf(lexer.IDENT) == nil {
		return Result{false, errors.New("Expected IDENT")}
	}

	if p.l.GetNextTokenIf(lexer.SEMICOLON) == nil {
		return Result{false, errors.New("Expected SEMICOLON")}
	}

	return Result{true, nil}
}

// var_decl_block
//     : "variáveis"
//       (var_decl ";")+
//       "fim-variáveis"
//     ;
func (p *Parser) parserVarDeclBlock() Result {
	if p.l.GetNextTokenIf(lexer.VARSBEGIN) == nil {
		return Result{false, nil}
	}

	pr := p.parserVarDecl()
	for ; pr.Parsed; pr = p.parserVarDecl() {
	}

	if pr.Parsed == false && pr.Err != nil {
		return pr
	}

	if p.l.GetNextTokenIf(lexer.VARSEND) == nil {
		return Result{false, errors.New("Expected FIM-VARIÁVEIS")}
	}

	return Result{true, nil}
}

// var_decl
//     : T_IDENTIFICADOR ("," T_IDENTIFICADOR)* ":" tp_primitivo
//     ;
func (p *Parser) parserVarDecl() Result {
	varID := p.l.GetNextTokenIf(lexer.IDENT)
	if varID == nil {
		return Result{false, nil}
	}

	if p.l.GetNextTokenIf(lexer.COLON) == nil {
		return Result{false, errors.New("Expected ':'")}
	}

	pr := p.parserPrimitiveType()
	if pr.Parsed == false {
		return Result{false, errors.New("Expected type definition")}
	}

	if p.l.GetNextTokenIf(lexer.SEMICOLON) == nil {
		return Result{false, errors.New("Expected ;")}
	}

	if p.symbol.Add(varID.Value) == -1 {
		return Result{false, errors.New("Duplicated variable")}
	}

	return Result{true, nil}
}

// tp_primitivo
//     : "inteiro"
//     | "literal"
//     | "lógico"
//     ;
func (p *Parser) parserPrimitiveType() Result {
	if p.l.GetNextTokenIf(lexer.INT) != nil || p.l.GetNextTokenIf(lexer.STRING) != nil {
		return Result{true, nil}
	}

	return Result{false, nil}
}

// stm_block
//     : "início"
//       (stm_list)*
//       "fim"
//     ;
func (p *Parser) parserStmBlock() Result {
	if p.l.GetNextTokenIf(lexer.BLOCKBEGIN) == nil {
		return Result{false, nil}
	}

	pr := p.parserStmList()
	for ; pr.Parsed; pr = p.parserStmList() {
	}

	if pr.Parsed == false && pr.Err != nil {
		return pr
	}

	if p.l.GetNextTokenIf(lexer.BLOCKEND) == nil {
		return Result{false, errors.New("Expected FIM")}
	}

	return Result{true, nil}
}

// stm_list
//     : stm_attr
//     | fcall ";"
//     | stm_ret
//     | stm_se
//     | stm_enquanto
//     | stm_para
//     ;
func (p *Parser) parserStmList() Result {
	pr := p.parserFunctionCall()
	if pr.Parsed == true {
		// Ensure that a ";" is presented at the EOL.
		if p.l.GetNextTokenIf(lexer.SEMICOLON) == nil {
			return Result{false, errors.New("Expected SEMICOLON")}
		}
		return Result{true, nil}
	}

	pr = p.parserStmAttr()
	if pr.Parsed == true {
		// Ensure that a ";" is presented at the EOL.
		if p.l.GetNextTokenIf(lexer.SEMICOLON) == nil {
			return Result{false, errors.New("Expected SEMICOLON")}
		}
		return Result{true, nil}
	}

	return Result{false, nil}
}

// stm_attr
//     : T_IDENTIFICADOR
//       ":="
//       expr
//       ";"
//     ;
func (p *Parser) parserStmAttr() Result {
	id, _ := p.l.GetNextsTokensIf(lexer.IDENT, lexer.ATTR)
	if id == nil {
		return Result{false, nil}
	}

	pr := p.parserExpr()
	if pr.Parsed == false {
		return Result{false, errors.New("Expected Expr")}
	}

	p.bc.Add(instructions.STV, p.symbol.Index(id.Value))

	return Result{true, nil}
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

func (p *Parser) parserExpr() Result {
	pr := p.parserFunctionCall()
	if pr.Parsed == true {
		return Result{true, nil}
	}

	id := p.l.GetNextTokenIf(lexer.IDENT)
	if id != nil {
		p.bc.Add(instructions.LDV, p.symbol.Index(id.Value))
		return Result{true, nil}
	}

	token := p.l.GetNextTokenIf(lexer.STRING)
	if token != nil {
		cpIndex := p.cp.Add(token.Value)
		p.bc.Add(instructions.LDC, cpIndex)
		return Result{true, nil}
	}

	return Result{false, nil}
}

// fcall
//     : T_IDENTIFICADOR "(" fargs? ")"
//     ;
func (p *Parser) parserFunctionCall() Result {
	token, _ := p.l.GetNextsTokensIf(lexer.IDENT, lexer.LPAREN)
	if token == nil {
		return Result{false, nil}
	}

	funcIndex := -1
	if token.Value == "imprima" {
		funcIndex = p.cp.Add("io.println")
	} else if token.Value == "leia" {
		funcIndex = p.cp.Add("io.readln")
		// } else {
		// 	return Result{false, errors.New("Undefined function name")}
	}

	pr := p.parserFunctionArgs()
	if pr.Parsed == false && pr.Err != nil {
		return pr
	}

	if token.Value == "imprima" {
		argsCountIndex := p.cp.Add(p.argsCount)
		p.bc.Add(instructions.LDC, argsCountIndex)
	}

	p.bc.Add(instructions.CALL, funcIndex)

	if p.l.GetNextTokenIf(lexer.RPAREN) == nil {
		return Result{false, errors.New("Expected RPAREN")}
	}

	return Result{true, nil}
}

// fargs
//     : expr ("," expr)*
//     ;
func (p *Parser) parserFunctionArgs() Result {
	p.argsCount = 0

	pr := p.parserExpr()
	if pr.Parsed == false {
		return pr
	}

	p.argsCount++

	for {
		if p.l.GetNextTokenIf(lexer.COMMA) == nil {
			return Result{true, nil}
		}
		pr := p.parserExpr()
		if pr.Parsed == false {
			return Result{false, errors.New("Expected EXPR")}
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
