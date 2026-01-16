<?php

$teclado = fopen('php://stdin', 'r');

$novoCurso = fgets($teclado);

file_put_contents('lista.txt',  $novoCurso . "\n", FILE_APPEND);

// o php://stdin é uma forma de entrada padrão para o teclado
// outra maneira de receber input é fgets(SDTIN), pois o php já entende automaticamente como entrada de dados do usuário
// o console pode ser tratado como um arquivo e para utilizá-lo como, basta utilizar STDOUT