<?php

$nome = 'Rodrigo Panisi Pombo';

$familia = str_contains($nome, 'Panisi');

if ($familia) {
    echo "$nome é da minha família" . PHP_EOL;
}
else {
    echo "$nome não é da minha familia" . PHP_EOL;
};

// str_contains() retorna um valor booleano