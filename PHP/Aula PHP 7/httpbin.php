<?php

$contexto = stream_context_create([
    'http' => [
        'method' => 'post',
        'header' => "X-From: PHP\r\nContent-type: text/plain",
        'content' => 'Teste do corpo'
    ]
]);

echo file_get_contents('http://httpbin.org/post', false, $contexto);

// esse código serve para modificar requisições na web, adicionando contexto para essas requisições
// ele pode ser utilizado para outros tipos de wrappers, não apenas para http
// o contexto pode ser adicionado em todos os comandos que utilizam streams