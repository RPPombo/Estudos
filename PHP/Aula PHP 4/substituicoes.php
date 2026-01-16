<?php

$texto = 'porra, porra, caralho, estou descontrolado';
echo  str_replace(['porra', 'caralho'], '***', $texto). PHP_EOL;

echo strtr($texto, 'porra', '***') . PHP_EOL;

// o str_replace(), subbstitui exatamente o que está sendo pedido
// o strtr() substitui os caracteres para outros
// porém se colocado um array em strtr() no segundo parâmetro as chaves do array serão as palavras substituidas e os valores pelo o que elas serão substituidas