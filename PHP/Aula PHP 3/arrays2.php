<?php

$alunos = [
    'Matheus',
    'Leo',
    'Rafael',
    'Gabriel',
    'Thomas'
];

array_push($alunos, 'João', 'Gabriel', 'Marco');
$alunos[] = 'Luiz';
array_unshift($alunos, 'Rafaela', 'Adriana', 'Bernardo');
var_dump($alunos);

// array_push() adiciona elementos a um array e retorna quantos elementos foram adicionados
// o array_unshift() adiciona elementos no começo do arrays
// array_shift() tira o primeiro elemento do array e o retorna para uma variável
// ele também altera as chaves numéricas para começar do zero
// array_pop() tira o ultimo elemento de um array e o retorna