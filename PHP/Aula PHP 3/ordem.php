<?php

$notas = [
    [
        'aluno' => 'Rodrigo',
        'nota' => '10'
    ],
    [
        'aluno' => 'Sophya',
        'nota' => '7'
    ],
    [
        'aluno' => 'Arthur',
        'nota' => '6'
    ]
];

function ordenaNotas(array $nota1, array $nota2): int
{
    return $nota1['nota'] <=> $nota2['nota'];
    /*
    if ($nota1['nota'] > $nota2['nota']){
        return 1;
    }
    elseif($nota2['nota'] > $nota1['nota']) {
        return -1;
    }
    else{
        return 0;
    }
    */
    // dessa maneira ele irá fazer na ordem crescente, caso queira na ordem decrescente é só mudar a posição da nota1 com a nota2
} 

usort($notas, 'ordenaNotas');

foreach($notas as $registro){
    echo $registro['nota'] . PHP_EOL;
}

// para ordenações mais simples basta utilizar as funções sort() e rsort(), a primeira ordena na ordem crescente e a segunda na ordem decrescente
// quando você usa uma dessas funções, as chaves acabam se perdendo
// para mantê-las basta usar asort() e arsort()
// para ordenar um array usando as chaves, por exemplo em ordem alfabética, utiliza-se ksort() e krsort()