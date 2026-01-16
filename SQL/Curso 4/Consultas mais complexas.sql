SELECT strftime('%Y/%m', v.data_venda) AS "Ano/Mês",f.nome AS "Nome do Forncedor", COUNT(iv.produto_id) AS "Número de Vendas" 
FROM itens_venda iv
JOIN vendas v ON iv.venda_id = v.id_venda
JOIN produtos p ON iv.produto_id = p.id_produto
JOIN fornecedores f ON f.id_fornecedor = p.fornecedor_id
WHERE strftime('%m', v.data_venda) = '11'
GROUP BY f.nome, strftime('%Y/%m', v.data_venda)
ORDER BY f.nome;

SELECT strftime('%Y/%m', v.data_venda) AS "Ano/Mês",c.nome_categoria AS "Nome da Categoria", COUNT(iv.produto_id) AS "Número de Vendas" 
FROM itens_venda iv
JOIN vendas v ON iv.venda_id = v.id_venda
JOIN produtos p ON iv.produto_id = p.id_produto
JOIN categorias c ON p.categoria_id = c.id_categoria
WHERE strftime('%m', v.data_venda) = '11'
GROUP BY c.nome_categoria, strftime('%Y/%m', v.data_venda)
ORDER BY c.nome_categoria;