<?php

$alunos2023 = [
    'Matheus',
    'Leo',
    'Rafael',
    'Gabriel',
    'Thomas'
];

$novosAlunos = [
    'Antonio',
    'Yuji',
    'Bianca',
    'Jefferson',
    'Ana',
    'Fernanda'
];

$alunos2024 = array_merge($alunos2023, $novosAlunos);
var_dump($alunos2024);

$alunos2024 = $alunos2023 + $novosAlunos;
var_dump($alunos2024);

// o array_merge() não preserva as chaves
// caso as chaves de string forem iguais, o valor será sobrescrito do anterior
// caso as chaves sejam numéricas, o número mudará para o próximo logo em seguida
// no caso de soma, as keys iguais seram ignoradas no segundo array e o só seram colocadas as que não existem
// para desempacotar um array basta colocar '...' antes do nome do array
// se os '...' forem usados para a definição de uma função. EX: function funcao(int ...$numeros){}
// o spread oprator faz com que possa ser possível colocar vários números
// esses números se transformam em um array com todos eles