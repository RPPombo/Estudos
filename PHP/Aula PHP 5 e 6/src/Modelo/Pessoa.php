<?php
namespace Alura\Banco\Modelo;
class Pessoa
{
    protected string $nome;
    private CPF $cpf;

    public function __construct(string $nome, CPF $cpf)
    {
        $this->nome = $nome;
        $this->cpf = $cpf;
    }

    public function getNome() : string
    {
        return $this->nome;
    }

    public function getNumeroCPF() : string
    {
        return $this->cpf->getCPF();
    }
}

// um atributo 'protected' é um atributo que só pode ser acessado pela classe mãe ou filhas