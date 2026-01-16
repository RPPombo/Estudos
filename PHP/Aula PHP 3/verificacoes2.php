<?php

$notas1 = [
    'Rafael' => 8,
    'Matheus' => 10,
    'Rodrigo' => 7,
    'Arthur' => 6,
    'Sophya' => 9,
    'Leo' => 7,
];

$notas2 = [
    'Rafael' => 8,
    'Rodrigo' => 7,
    'Arthur' => 8,
    'Sophya' => 10,
    'Leo' => 7,
];

var_dump(array_diff($notas1, $notas2));
var_dump(array_diff_key($notas1, $notas2));
var_dump(array_diff_assoc($notas1, $notas2));

$alunosFaltantes = array_diff_key($notas1, $notas2);
var_dump(array_keys($alunosFaltantes));

// array_diff() verifica os valores que não estão presentes nos dois arrays
// o próximo verifica apenas as chaves
// e o assoc verifica os dois
// para juntar dois arrays sendo um as chaves e outro os valores, usa-se array_combine(chave, valor), porém os dois devem ter o mesmo tamanho
// para inverter as chaves com os valores de um array, usa-se array_flip()