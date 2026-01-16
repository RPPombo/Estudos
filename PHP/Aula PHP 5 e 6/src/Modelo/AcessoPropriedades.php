<?php

namespace Alura\Banco\Modelo;

trait AcessoPropriedades
{
    public function __get(string $atributo)
    {
        $metodo = 'get' . ucfirst($atributo);
        return $this->$metodo();
    }
}

// a trait serve para injetar uma função em alguma classe
// por exemplo essa foi usada na classe endereço