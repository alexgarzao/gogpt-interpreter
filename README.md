# gogpt-interpreter

Um interpretador para a linguagem G-Portugol, escrito em Golang :-)

A linguagem G-Portugol foi criada pelo Thiago Silva, tendo por finalidade o ensino de algoritmos. Maiores informações sobre a linguagem podem ser obtidas [aqui](http://inf.ufes.br/~mberger/Disciplinas/2015_2/Compiladores/manualGPortugol.pdf).

[Aqui](https://pt.wikipedia.org/wiki/G-Portugol) também existe uma breve descrição do propósito do projeto.

A versão atual do gogpt consegue interpretar e executar apenas algoritmos simples em G-Portugol, algoritmos como "Olá mundo" e leitura de dados alfanuméricos.

## Requisitos

Como este projeto foi escrito em Golang, é necessário ter o ambiente de Golang instalado. Maiores informações sobre como instalar podem ser obtidas [aqui](https://golang.org/).

## Como executar

    go get github.com/alexgarzao/gogpt-interpreter/gogpt

    cat <<EOT >> hello_world.gpt
    algoritmo olá_mundo;

    início
        imprima("Olá mundo!");
    fim
    EOT


    gogpt hello_world.gpt

## Para gerar o binário a partir dos fontes

    mkdir -p ~/go/src/github.com/alexgarzao/
    cd ~/go/src/github.com/alexgarzao/
    git clone https://github.com/alexgarzao/gogpt-interpreter.git
    cd gogpt-interpreter/gogpt
    make unittests
    make build
    ./gogpt/gogpt ../samples/hello_world.gpt

## Arquitetura do projeto

Como a ideia do projeto é praticar Golang, somente foram utilizadas as bibliotecas padrões de Golang.

Este projeto é composto por um compilador que lê os fontes em G-Portugol e gera um bytecode para ser executado por uma VM interna.

A VM é "stack based", e este modelo foi escolhido por ser de fácil implementação. A versão atual somente tem suporte aos tipos primitivos "int" e "string", e os tipos são inferidos durante a execução da VM.

## O que falta

* Expressões relacionais
* Expressões aritméticas
* Estruturas de repetição
* Tipos numéricos, lógicos, matrizes
* Definição de funções
* ...


## Qual a finalidade deste projeto

A ideia inicial é o "just for fun". Eu tenho muito interesse nas áreas de compiladores e máquinas virtuais, e como queria aprender mais de Golang, nada melhor do que um projeto para focar :-)
