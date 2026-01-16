<?php

$telefones = ['(21) 9999 - 9999', '(24) 9999 - 9999', '(13) 9999 - 9999'];

foreach ($telefones as $telefone){
    $telefoneValido = preg_match('/^\([0-9]{2}\) 9?[0-9]{4} - [0-9]{4}$/', $telefone, $correspondencia);
    if ($telefoneValido){
        echo "$telefone: Esse telefone é válido". PHP_EOL;
    } else{
        echo "$telefone: Esse telefone é inválido". PHP_EOL;
    }
    var_dump($correspondencia);
}

echo preg_replace(
    '/^\(([0-9]{2})\) (9?[0-9]{4} - [0-9]{4})$/',
    '(XX) \2',
    $telefone
) . PHP_EOL;

// o preg_match() serve para comparar uma string a uma expressão regular, ou seja, uma padronização de algo
// o terceiro parâmetro serve para armazenar em uma variável a expressão regular encontrada na string