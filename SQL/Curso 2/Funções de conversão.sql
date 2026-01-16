SELECT ('O faturamento bruto médio foi: ' || CAST(ROUND(AVG(faturamento_bruto), 2) AS TEXT)) FROM faturamento;

--A função CAST() serve para converter uma variável de um tipo para outro tipo
--Outras funções de conversão:
--CONVERT(tipo, expressao [, estilo]), funciona de forma semelhante ao CAST(), porém utilizada em outros SGBDs

--Funções do Oracle:
--TO_NUMBER(string [, formato [, 'nlsparam']]), Converte strings para números]
--TO_CHAR(valor [, formato [, 'nlsparam']]),  números ou datas para strings
--TO_DATE(string [, formato [, 'nlsparam']]), strings para datas

--Funções do SQL SERVER:
--PARSE(string AS tipo USING cultura), tenta converter uma string para um tipo de dados numérico ou de data/hora com um estilo de cultura opcional
--TRY_PARSE(string AS tipo USING cultura), TRY_CONVERT(tipo, expressao [, estilo]), são versões mais seguras que retornam NULL em vez de um erro se a conversão falhar

--Funções do MySQL:
--STR_TO_DATE(string, formato), Converte uma string em um formato de data especificado para uma data

--Funções do PostgreSQL:
--TO_NUMBER(string, formato), converte uma string para um número com um formato especificado
--TO_CHAR(valor, formato), converte um número ou data para uma string com um formato especificado