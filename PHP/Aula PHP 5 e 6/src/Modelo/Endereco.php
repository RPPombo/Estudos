<?php
namespace Alura\Banco\Modelo;
final class Endereco
{
    use AcessoPropriedades;
    private string $cidade;
    private string $bairro;
    private string $rua;
    private string $numero;

    public function __construct(string $cidade, string $bairro, string $rua, string $numero)
    {
        $this->cidade = $cidade;
        $this->bairro = $bairro;
        $this->rua = $rua;
        $this->numero = $numero;
    }

    public function getCidade() : string
    {
        return $this->cidade;
    }

    public function getCBairro() : string
    {
        return $this->bairro;
    }

    public function getRua() : string
    {
        return $this->rua;
    }

    public function getNumero() : string
    {
        return $this->numero;
    }

    public function __toString() : string
    {
        return "$this->cidade, $this->bairro, $this->rua, $this->numero";
    }
}

// a função __toString() faz com que caso o objeto seja chamado como string ele irá mostrar o que foi colocado na função
// o __call() serve para caso seja chamado uma função que é privada ou protegida
// o __get() é chamado quando o nome de um atributo da classe é chamado mesmo quando privado
// o final impede que um método ou uma classe seja extendida para outra classe ou um método sobrescrito