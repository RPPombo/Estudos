<?php

$listaCursos = file('lista.txt');

$arquivoCSV = fopen('cursos.csv', 'w');

foreach($listaCursos as $curso) {
    $linha = [trim(mb_convert_encoding($curso, 'UTF-8')), 'Sim'];

    fputcsv($arquivoCSV, $linha, ';');
}

fclose($arquivoCSV);

// para fazer o processo contrário apenas troque fputcsv() por fgetcsv() e escreva as strings no arquivo desejado
// a função mb_convert_encoding() serve para converter o jeito de encodificar alguma string, podendo escolher o tipo de encoding