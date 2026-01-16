<?php
$idade = 18;

echo "Você só pode entrar se tiver 18 anos ou mais";

if ($idade >= 18) {
    echo "Você tem $idade anos" . PHP_EOL;
    echo "Pode entrar";
}

else {
    echo "Você não pode entrar" . PHP_EOL;
    echo "Você só tem $idade anos";
}
// Se vc for colocar uma decisão com 'ou' pode-se ser usado tanto o 'or' quanto ||
// Para o 'e' pode-se ser usado o 'and' ou '&&'