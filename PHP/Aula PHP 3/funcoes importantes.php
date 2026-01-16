<?php

$numbers = [1, 2, 3, 4, 5];
$evenNumbers = array_filter($numbers, function ($num) {
    return $num % 2 == 0;
});

$squaredNumbers = array_map(function ($num) {
    return $num * $num;
}, $numbers);

$sum = array_reduce($numbers, function ($carry, $num) {
    return $carry + $num;
}, 0);

var_dump($evenNumbers);
var_dump($squaredNumbers);
var_dump($sum);

// array_filter(): Essa função é utilizada para filtrar os elementos de um array com base em uma função de callback.
// array_map(): Essa função é utilizada para aplicar uma função a cada elemento de um array e retorna um novo array com os resultados.
// array_reduce(): Essa função é utilizada para reduzir um array a um único valor aplicando uma função de callback de forma cumulativa.