<?php

$diretorioAtual = dir('.');

while($arquivo = $diretorioAtual -> read()) {
    echo $arquivo . PHP_EOL;
}

// o '.' significa diretório atual em qualquer sistema operacional