SELECT UPPER('A pessoa colaboradora ' || nome || ' de cpf ' || cpf || ' possui o seguinte endereço: ' || endereco) as texto FROM Colaboradores;


--As || servem para concatenar strings e variaveis
--Existem também as funções UPPER e LOWER
--Além dessas funções apresentadas agora também existem: 
--TRIM(string, [caractere_para_trimar]), responsável por tirar os caracteres do começo e do fim de uma string (Espaço por padrão)
--INSTR(string, substring), retorna a posição de uma substring dentro de uma string (0, caso não encontre)
--REPLACE(string, substring_a_substituir, substring_para_substituir), troca uma substring por outra
--SUBSTR(string, inicio[, comprimento]), extrai a parte de uma string com base no ponto de início e comprimento dados