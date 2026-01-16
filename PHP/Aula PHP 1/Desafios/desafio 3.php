<?php
$altura = 1.74;
$peso = 71;

$IMC = $peso / ($altura ** 2);

echo "A pessoa está ";

if ($IMC <=18.5) {
    echo "abaixo do peso ideal";
}

else if ($IMC <=24.9) {
    echo "no pesso ideal";
}

else {
    echo "acima do peso ideal";
}