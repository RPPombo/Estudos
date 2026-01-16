<?php
 class Minha_exceção extends Exception {
    public function __construct(float $saldoAtual,float $valorSaque) 
    {
        $mensagem = "Você não pode sacar R$ {$valorSaque}, pois seu saldo atual é de R$ {$saldoAtual}.";
        parent::__construct($mensagem);
    }
  }

  // uma prática comum é de colocar o nome da exceção terminando com  "_Exception"
  // os outros parâmetros que podem ser passados para parent::__construct() são: o código do erro e a pilha  de chamadas.