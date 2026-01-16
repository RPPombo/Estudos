<?php

$arquivo = fopen('lista.txt','r');

while(feof($arquivo)){
    $linha = fgets($arquivo);
    echo $linha;
}

fclose($arquivo);





// fopen() serve para abrir arquivos
// fclose() fecha o arquivo
// fgets() swerve para ler uma linha de um arquivo até o final
// feof() verifica se o arquivo chegou no final
// para ler o arquivo no geral utilize fread($arquivo, $tamanho do arquivo)
// o tamanho do arquivo pode ser visto por meio da função filesize()
// se utilizar a função file_get_contents() é possível ler o arquivo de uma vez, ao invés de ter que usar duas funções para isso
// a função file() transforma cada linha do arquivo em um item de array
// é possível abrir conteúdos de sites utilizando as mesmas funções utilizadas para abrir arquivos locais
// é possível ler e escrever arquivos .zip sem descompactá-los