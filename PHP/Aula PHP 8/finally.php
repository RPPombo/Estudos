<?php

function a() : int
{
    try  {
        throw new Exception();
        echo'Código' . PHP_EOL;
        return 0;
    } catch(Throwable){
        echo  'Exceção' . PHP_EOL;
        return 1;
    } finally{
        echo 'Finally' . PHP_EOL;
    }
}

echo a();

// o finally sempre vai executar independente se deu erro ou não