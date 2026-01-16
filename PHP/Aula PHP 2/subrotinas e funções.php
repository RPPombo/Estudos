<?php

require 'funcoes.php';

$contas = [
    531836581 => [
        'nome' => 'Rodrigo',
        'saldo' => 3000
    ], 
    673216361 => [
        'nome' => 'Sophya',
        'saldo' => 10000
    ], 
    469236478 => [
        'nome' => 'Arthur',
        'saldo' => 1000
    ]
];

$contas[531836581] = sacar($contas[531836581], 700);
$contas[469236478] = sacar($contas[469236478], 7000);
$contas[673216361] = depositar($contas[673216361], 500);

$contas[469236478] = titularComLetrasMaiusculas($contas[469236478]);

foreach ($contas as $cpf => $conta){
    list('nome' => $nome, 'saldo' => $saldo) = $conta;
    exibeMensagem("$cpf => titular: $nome saldo: $saldo");
};

// uma subrotina é uma função que não retorna nenhum valor e uma função é a que retorna valor
// para importar um arquivo pode-se usar tanto include, quanto require,
// porém caso o arquivo não seja encontradob o include continuará rodando o programa, já o require dará erro na hora
// caso as chaves não tenham nome definido será feito na ordem a atribuição na função list()
// para remover uma variável é só utilizar a função unset()
// quando ocorre o fechamento do código php, tudo abaixo dele será lido como texto
// isso pode ser usado para criar um html
?>
<!DOCTYPE html>
<html lang="pt-br">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>PHP teste</title>
</head>
<body>
    <h1>Contas correntes</h1>

    <dl>
        <?php foreach($contas as $cpf => $conta) { ?>
        <dt>
            <h3><?= $conta['nome']; ?> - <?= $cpf; ?></h3>
        </dt>
        <dd>Saldo: <?= $conta['saldo']; ?></dd>
        <?php } ?>
    </dl>
</body>
</html>