SELECT DISTINCT strftime('%Y', data_venda) AS Ano
FROM vendas 
ORDER BY Ano;

SELECT strftime('%Y', data_venda) AS Ano, COUNT(id_venda) AS Total_Vendas
FROM vendas
GROUP BY Ano
ORDER BY Ano;

SELECT strftime('%Y', data_venda) AS Ano, strftime('%m', data_venda) AS Mes, COUNT(id_venda) AS Total_Vendas
FROM vendas
GROUP BY Ano, Mes
ORDER BY Ano;

SELECT strftime('%Y', data_venda) AS Ano, strftime('%m', data_venda) AS Mes, COUNT(id_venda) AS Total_Vendas
FROM vendas
WHERE Mes = '01' OR Mes = '11' OR Mes = '12'
GROUP BY Ano, Mes
ORDER BY Ano;

--Para extrair apenas uma parte de uma data, utiliza-se a função strftime() com a parte da data quer extrair e a coluna na qual está as datas
--A função DISTINCT serve para mostrar apenas uma vez dados que se repetem
--Para melhores formatações utilizando o strftime() pode-se usar os seguintes formatos: %Y(ano), %m(mês), %d(dia), %H(hora), %M(minuto)