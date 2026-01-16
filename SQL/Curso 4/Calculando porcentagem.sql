SELECT "Nome da Categoria", "Número de Vendas", ROUND(100.0 * "Número de Vendas"/(SELECT COUNT(*) FROM itens_venda), 2) || '%' AS "Porcentagem"
FROM(
  SELECT c.nome_categoria AS "Nome da Categoria", COUNT(iv.produto_id) AS "Número de Vendas" 
  FROM itens_venda iv
  JOIN vendas v ON iv.venda_id = v.id_venda
  JOIN produtos p ON iv.produto_id = p.id_produto
  JOIN categorias c ON p.categoria_id = c.id_categoria
  GROUP BY c.nome_categoria
  ORDER BY "Número de Vendas" ASC);