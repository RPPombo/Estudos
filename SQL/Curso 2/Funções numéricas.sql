SELECT AVG(faturamento_bruto), ROUND(AVG(faturamento_bruto), 2) FROM faturamento;

SELECT CEIL(faturamento_bruto), CEIL(despesas) FROM faturamento;

SELECT FLOOR(faturamento_bruto), FLOOR(despesas) FROM faturamento;

--ROUND(), arredonda o número para a quantidade de casas decimais determinadas
--CEIL(), arredonda para o número inteiro maior mais próximo
--FLOOR(), arredonda para o número inteiro menor mais próximo
--Outras funções numéricas:
--POWER(base, expoente), serve para elevar um número a uma potência desejada
--SQRT(numero), serve para descobrir a raiz quadrada
--RANDOM(), gera um número inteiro aleatório
--ABS(numero), retorna o número absoluto, ou seja, sem sinal
--HEX(numero_ou_string), retorna um número ou uma string na sua forma hexadecimal