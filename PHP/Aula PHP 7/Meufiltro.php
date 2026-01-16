<?php

class Meufiltro extends php_user_filter
{
    public $stream;
    public function onCreate() : bool
    {
        $this->stream = fopen('php://temp', 'w+');
        return $this->stream !== false;
    }

    public function filter($in, $out, &$consumed, $closing) : int
    {
        $saida = '';
        while ($bucket = stream_bucket_make_writeable($in)){
            $linhas = explode("\n", $bucket->data);
            
            foreach ($linhas as $linha){
                if (str_contains($linha, 'parte') !== false){
                    $saida .= $linha . PHP_EOL;
                }
                if (str_contains($linha, 'Parte') !== false){
                    $saida .= $linha . PHP_EOL;
                }
            }
        }
        $bucketSaida = stream_bucket_new($this->stream, $saida);
        stream_bucket_append($out, $bucketSaida);

        return PSFS_PASS_ON;
    }
}

// o while na função filter() faz com que enquanto tiver conteúdo de entrada que pode ser usado, o código será rodado