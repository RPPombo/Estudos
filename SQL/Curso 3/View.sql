CREATE VIEW ViewCliente AS
SELECT nome, endereco FROM clientes;

SELECT * FROM ViewCliente;

CREATE VIEW ValorTotalPedido AS
SELECT p.id, c.nome, p.datahorapedido, SUM(ip.precounitario) "Soma do pedido"
FROM pedidos p
JOIN itenspedidos ip ON p.id = ip.idpedido
JOIN clientes c ON p.idcliente =c.id
GROUP BY p.id, c.nome;

SELECT * FROM ValorTotalPedido;

-- As Views servem para criar um atalho para pesquisas que muita vezes são recorrentes
-- Elas fuuncionam como um tipo de "tabela virtual"