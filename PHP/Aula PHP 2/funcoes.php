<?php
function exibeMensagem ($mensagem) 
{
    echo $mensagem . PHP_EOL;
};

function sacar(array $conta, float $valor) : array
{
    if ($valor > $conta['saldo']){
        exibeMensagem('Você não pode sacar');
    }
    else{
        $conta['saldo'] -= $valor;
    };
    return $conta;
}

function depositar(array $conta, float $valor) : array
{
    if ($valor < 0){
        exibeMensagem('Você precisa depositar um número positivo');
    }
    else{
       $conta['saldo'] += $valor;
    };
    return $conta;
}

function titularComLetrasMaiusculas(array $conta) : array
{
    $conta['nome'] = mb_strtoupper($conta['nome']);
    return $conta;
}