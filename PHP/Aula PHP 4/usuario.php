<?php

$nome = 'Rodrigo Panisi Pombo';
$email = 'rodrigo@gmail.com';
$senha = '123';

if (mb_strlen($senha) < 8){
    echo 'Senha muito fraca' . PHP_EOL;
}

$posicao = strpos($email, '@');

$usuario = substr($email, 0, $posicao);
echo mb_strtoupper($usuario) . PHP_EOL;
echo substr($email, $posicao + 1);

list($nome, $sobrenome) = explode(' ', $nome, 2);

echo "Nome: $nome" . PHP_EOL;
echo "Sobrenome: $sobrenome" .PHP_EOL;

// o strlen() retorna o número de bytes de uma string, ou seja, caracteres especiais, como ç, á, õ, retornam números maiores do que realmente tem na string
// já o mb_strlen() consegue ler os caracteres especiais normalmente
// o explode() separa uma string em um array
// o último parâmetro serve para limitar o número de itens do array, em que o último item vai ficar com tudo o que resta da string