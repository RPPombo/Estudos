<?php

$dados = [
    'nome' => 'Rodrigo', 
    'nota' => 7, 
    'idade' => 18];
list('nome' => $nome, 'nota' => $nota, 'idade' => $idade) = $dados;

var_dump($nome, $nota, $idade);

var_dump(compact('nome', 'nota', 'idade'));

// a função list() faz com que o código atribua significados a elementos de um array
// a função extract() tem a mesma funcionalidade em que ela cria variáveis a aprtir das chaves
// porém ela só deve ser usada em dados confiáveis, pois caso contrário ela pode criar variáveis indesejadas
// a função compact() faz o caminho inverso e cria um array a partir das variáveis citadas