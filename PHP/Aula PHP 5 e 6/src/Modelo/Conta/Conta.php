<?php
namespace Alura\Banco\Modelo\Conta;

class Conta
{
    private Titular $titular;
    private float $saldo;
    private static $numeroDeContas = 0;

    public function __construct(Titular $titular)
    {
        $this->titular = $titular;
        $this->saldo = 0;

        self::$numeroDeContas++;
    }

    public function sacar(float $valoraSacar) : void
    {
        $tarifaSaque = $valoraSacar * 0.05;
        $valorSaque = $valoraSacar + $tarifaSaque;
        if ($valorSaque > $this->saldo){
            echo 'Saldo insuficiente' . PHP_EOL;
            return;
        } 
        
        $this->saldo -= $valorSaque;
    }


    public function depositar(float $valoraDepositar) : void
    {
        if ($valoraDepositar < 0){
            echo 'O valor deve ser positivo' . PHP_EOL;
            return;
        }

        $this->saldo += $valoraDepositar;
    }

    public function transferir(float $valoraTransferir, Conta $contaDestino) : void
    {
        if ($valoraTransferir > $this->saldo){
            echo 'Saldo insuficiente' . PHP_EOL;
            return;
        }

        $this->sacar($valoraTransferir);
        $contaDestino->depositar($valoraTransferir);
    }

    public function getSaldo() : float
    {
        return $this->saldo;
    }

    public static function recuperaNumeroDeContas() : int
    {
        return self::$numeroDeContas;
    }

    public function __destruct()
    {
        self::$numeroDeContas--;
    }
}

// para criar um objeto do tipo Conta utiliza-se $variavel = new Conta()
// quando é trabalhado com objetos o que é salvo é o endereço para o objeto, não os valores do objeto
// o $this serve para se reefrir ao próprio objeto que chamou a função
// quando é deixado um item privado, significa que ele só pode ser acessado por métodos da própria classe
// uma função com o nome __construct() sempre vai ser ativada quando um novo objeto com essa classe for criada
// um atributo estático é um atributo da própria classe e não de um objeto específico
// um método com o nome __destruct() serve para quando o objeto deixar de existir ele ativar
// caso precise fazer uma muança para uma função específica, como por exemplo em uma conta corrente a tarifa de saque ser 0.05 e em  uma poupança 0.03
// você pode criar uma classe filha que sobrescreve a função, apenas colocando o nome dela como o mesmo
// para não ter que implementar um método no modelo padrão de uma conta e sim somente nas contas filhas, crie o método como abstract e a classe modelo como abstract também