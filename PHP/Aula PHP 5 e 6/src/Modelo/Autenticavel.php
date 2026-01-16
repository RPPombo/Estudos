<?php

namespace Alura\Banco\Modelo;

interface Autenticavel
{
    public function podeAutenticar(string $senha) : bool;
}

// a interface serve para fazer uma 'classe' com todos os métodos abstratos