<?php
$lista = array(31, 24, 99, 91, 105, 9);

$lista[] = 812;

for ($i = 0; $i < count($lista); $i ++) {
    echo $lista[$i] . PHP_EOL;
}

$contas = [
    531836581 => [
        'nome' => 'Rodrigo',
        'saldo' => 3000
    ], 
    673216361 => [
        'nome' => 'Sophya',
        'saldo' => 10000
    ], 
    469236478 => [
        'nome' => 'Arthur',
        'saldo' => 1000
    ]
];

foreach ($contas as $cpf => $conta) {
    echo $conta['nome'] . PHP_EOL;
}

// O nome das chaves de um array só pode ser uma string ou inteiro, caso seja de outro tipo o nome será convertido
// se um número estiver em string, ele virará um inteiro: '8' => 8