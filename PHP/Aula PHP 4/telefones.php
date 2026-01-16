<?php

$telefones = ['    (21) 9999-9999', '(24) 9999-9999', '(13) 9999-9999     '];


$string = implode(PHP_EOL, $telefones);
echo trim($string, " ");
// o implode() junta o array em uma string
// o trim() apenas apara as pontas da string
// se quiser aparar apenas uma ponta específica utilize ltrim() ou rtrim()