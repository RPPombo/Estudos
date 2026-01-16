<?php

error_reporting(E_ALL);

ini_set('display_errors', 1);
ini_set('log_errors', 0);

set_error_handler(function($errno, $errstr, $errfile, $errline){
    switch($errno){
        case E_NOTICE:
            echo "Notice em {$errfile}:{$errline}\n";
            break;
        case E_WARNING:
            echo "Warning em {$errfile}:{$errline}\n";
            break;
    }
});

echo $variavel;
echo $array[12];

echo CONSTANTE;

// o error_reporting() serve para escolher os erros q vão ser mostrados pelo php em um arquivo
// por padrão são mostrados todos a partir do php 8
// em display_errors(), é configurado se é mostrado ou não os erros no display
// em log_errors(), é configurado se os erros vão ser jogados para um log
// em error_log(), é configurado o lugar do log
// a função set_error_handler(), serve para configurar os erros