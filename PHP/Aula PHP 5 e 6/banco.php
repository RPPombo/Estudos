<?php
spl_autoload_register(function(string $nomeCompletoDasClasse){
    $caminhoDaClasse = str_replace('Alura\\Banco', 'src', $nomeCompletoDasClasse);
    $caminhoDaClasse = str_replace('\\',DIRECTORY_SEPARATOR,$caminhoDaClasse);
    $caminhoDaClasse .= '.php';

    if(file_exists($caminhoDaClasse)){
        require_once $caminhoDaClasse;
    }
});

use Alura\Banco\Modelo\Conta\{Conta, Titular};
use Alura\Banco\Modelo\{CPF, Endereco, Funcionario, Pessoa};

// esse arquivo é um modelo para fazer um arquivo com várias classes carregar
// quando colocar .= vai adicionar ao final de uma string outra string
// spl_autoload_register() serve para carregar automaticamente os arquivos
// ao invés de ter que colocar require em cada um dos arquivos