<?php

$notas = [
    'Rafael' => 8,
    'Matheus' => 10,
    'Rodrigo' => 7,
    'Arthur' => 6,
    'Sophya' => null,
    'Leo' => 7,
];

echo 'É um array?' . PHP_EOL;
var_dump(is_array($notas));

echo 'É uma lista?' . PHP_EOL;
var_dump( array_is_list($notas));

echo 'O Rafael está no array?' . PHP_EOL;
var_dump(array_key_exists('Rafael', $notas));

echo 'A Sophya fez a prova?' . PHP_EOL;
var_dump(isset($notas['Sophya']));

echo 'Teve nota 7?' . PHP_EOL;
var_dump(in_array(7, $notas, true));

echo 'Quem tirou 7?' . PHP_EOL;
echo array_search(7, $notas);

// o array_is_list() verifica se o array é um array ordenado numericamente, retornando uma variavel booleana
// o array_key_exists() verifica se uma key existe em um array
// o isset() verifica se a variável existe e se ela não é nula
// o in_array() busca por elementos dentro das chaves do array
// o true no final significa que a função está fazendo uma comparação mais minuciosa, ou seja, o php não vai igualar 10 com '10'
// o array_search() procura a key que possui o elemento necessitado, porém ele só irá mostrar a primeira key encontrada com o elemento
// caso ele não encontre ele irá retornar o valor booleano FALSE