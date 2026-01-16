<?php

$arquivo = fopen('lista.txt', 'a');

$curso = 'Design Patterns PHP I: boa prática de programação';

fwrite($arquivo, $curso, 21);

fclose($arquivo);

// o modo de abertura 'w' sobrescreve o que já estava escrito em um arquivo
// o modo de abertura 'a' serve para acrescentar coisas ao arquivo
// para adicionar algo em um arquivo sem abrí-lo é só usar a função file_put_contents($nome do arquivo, $o que vai coloacar, $flags)
// as flags são marcações, como por exemplo FILE_APPEND, em que faz com que a função não sobrescreva e sim adicione ao arquivo