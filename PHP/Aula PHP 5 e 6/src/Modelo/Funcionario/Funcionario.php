<?php
namespace Alura\Banco\Modelo\Funcionario;

use Alura\Banco\Modelo\{Pessoa,CPF};
abstract class Funcionario extends Pessoa
{
    protected float $salario;

    public function __construct(string $nome, CPF $cpf, float $salario)
    {
        parent::__construct($nome, $cpf);
        $this->salario = $salario;
    }

    abstract public function calculaBonificacao() : float;
    public function getSalario() : float
    {
        return $this->salario;
    }
}

// quando uma classe extende outra significa que essa classe vai ter os atributos da calsse a qual ela está estendendo