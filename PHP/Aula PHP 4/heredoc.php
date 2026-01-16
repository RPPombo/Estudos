<?php

function geraEmail(string $nome): void
{
    $conteudoEmail = <<<Delimitador
    Olá, $nome

    Entramos em contato para
    {conteúdo do e-mail}

    {assinatura}
    Delimitador;

    echo $conteudoEmail;
}

geraEmail('Rodrigo');

// o delimitador pode ter qualquer nome e ele serve para retirar todos os TAB atrás dele, quando é fechado
// ele tem uma função similar às aspas duplas
// para funcionar como aspas simples basta colocar <<<'Delimitador' assim