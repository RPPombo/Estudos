<?php

$contador = 1;
 
while ($contador <=15) {
    if ($contador == 10) {
        break;
    }
    echo "#$contador" . PHP_EOL;
    $contador = $contador +1;
}

for ($contador = 1; $contador <=15; $contador ++) {
    if ($contador == 13){
        continue;
    }
    echo "#$contador" . PHP_EOL;
}

//$contador ++ é igual $contador += 1, que também é igual a $contador = $contador + 1