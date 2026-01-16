<?php

require 'Meufiltro.php';

$arquivos = fopen('lista.txt', 'r');

stream_filter_register('alura.partes', 'Meufiltro');
stream_filter_append($arquivos, 'alura.partes');

echo fread($arquivos, filesize('lista.txt'));

fclose($arquivos);