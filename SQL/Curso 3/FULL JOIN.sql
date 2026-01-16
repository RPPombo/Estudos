SELECT c.nome, p.id 
FROM clientes c
FULL JOIN pedidos p
ON c.id = p.idcliente
WHERE strftime('%m', p.datahorapedido) = '10';

SELECT c.nome, p.id 
FROM clientes c
FULL JOIN pedidos p
ON c.id = p.idcliente
WHERE c.id IS NULL;

-- O FULL JOIN serve para mostrar todos os dados das duas tabelas mesmo que ela não tenha uma correspondência na outra
-- Outra forma de junção é o CROSS JOIN, que a função é muito pouco frequente, mas serve para juntar de todas as maneiras possíveis duas tabelas, mesmo que essa junção n tenha uma relação lógica