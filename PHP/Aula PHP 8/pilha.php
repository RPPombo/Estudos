<?php

function funcao1()
{
    echo 'Entrando na função 1'. PHP_EOL;

    try{
        funcao2();
    } catch(Throwable $problema){
        echo  "Erro encontrado em função 1: {$problema->getMessage()}".PHP_EOL;
        echo $problema->getLine(). PHP_EOL;
        echo $problema->getTraceAsString(). PHP_EOL;
    }

    echo 'Saindo da função 1'. PHP_EOL;
}

function funcao2()
{
    echo 'Entrando na função 2'. PHP_EOL;

    for($i = 1;$i<5; $i++) {
        echo $i. PHP_EOL;
    }
    
    throw  new RuntimeException('Deu merda patrão');
    
    echo  'Saindo da função 2'. PHP_EOL;
}   

echo 'Entrando na main'. PHP_EOL;
funcao1();
echo 'Saindo da main'. PHP_EOL;

// o bloco try-catch serve para tratar um throwable gerada no código selecionado pelo try
// a função getTraceAsString() serve para ver a pilha de execução até dar o erro
// a função getLine() mostra a linha em que deu o problema
// a função getMessage() diz qual foi o problema
// na criação de um objeto da classe de exceção os parâmetros colocados são a mensagem da exceção, o código dela e a exceção anterior