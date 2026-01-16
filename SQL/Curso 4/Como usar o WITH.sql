WITH Media_Vendas_Anteriores AS
--Primeira Subconsulta
(SELECT AVG(Qtd_Vendas) AS Media_Qtd_Vendas
FROM (
    SELECT COUNT(*) AS Qtd_Vendas, strftime('%Y', v.data_venda) AS Ano
    FROM vendas v
    WHERE strftime('%m', v.data_venda) = '11' AND Ano != '2022'
    GROUP BY Ano
)),
Vendas_Atual AS
--Segunda Subconsulta
(SELECT Qtd_Vendas AS Qtd_Vendas_Atual
FROM(
    SELECT COUNT(*) AS Qtd_Vendas, strftime('%Y', v.data_venda) AS Ano
    FROM vendas v
    WHERE strftime('%m', v.data_venda) = '11' AND Ano = '2022'
    GROUP BY Ano
    ))
    SELECT
    mva.Media_Qtd_Vendas,
    va.Qtd_Vendas_Atual,
    ROUND((va.Qtd_Vendas_Atual - mva.Media_Qtd_Vendas)/mva.Media_Qtd_Vendas *100.0, 2) || '%' AS Porcentagem
    FROM Vendas_Atual va, Media_Vendas_Anteriores mva
    
 --O WITH serve para criar uma consulta que junta duas subconsultas, de forma que seja possível utilizar os dados das subconsultas na consulta atual
 --Ele faz isso fazendo com que as subconsultas se tornem tabelas temporárias que são criadas apenas para essa consulta