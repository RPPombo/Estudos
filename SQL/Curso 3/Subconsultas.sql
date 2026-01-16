SELECT nome, telefone 
FROM clientes 
WHERE ID = (
    SELECT idcliente 
    FROM pedidos 
    WHERE datahorapedido = '2023-01-02 08:15:00');
    
-- Essa consulta com uma consulta dentro dela são chamadas de subconsultas e servem para procurar informações utilizando duas tabelas