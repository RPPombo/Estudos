<?php
namespace Alura\Banco\Service;

use Alura\Banco\Modelo\Funcionario\Funcionario;

class ControladorDeBonificacoes
{
    private $totalDeBonificacoes = 0;

    public function adicionaBonificacao(Funcionario $funcionario)
    {
        $this->totalDeBonificacoes += $funcionario->calculaBonificacao();
    }

    public function recuperaTotal() : float
    {
        return $this->totalDeBonificacoes;
    }
}