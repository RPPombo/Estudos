SELECT c.nome, c.id, p.datahorapedido
from clientes c
INNER JOIN pedidos p
ON c.id = p.idcliente;

-- O INNER JOIN serve para retorna os valores que possuem uma associação entre as duas tabelas (FOREIGN KEY)
-- Nesse caso a associação feita foi entre o ID de cliente em que na tabela Cliente corresponde ao id e na tabela Pedidos corresponde a idcliente