<?php

$url = 'https://alura.com.br';

$teste = str_starts_with($url, 'https');

if ($teste) {
    echo "A url: $url é segura";
} else{
    echo "A url: $url não é segura";
}

$teste2 = str_ends_with($url, '.br');

if ($teste2) {
    echo ' e é brasilera';
} else {
    echo ' e não é brasileira';
}

echo PHP_EOL;