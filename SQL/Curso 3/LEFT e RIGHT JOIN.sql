-- RIGHT JOIN
INSERT INTO Produtos (ID, Nome, Descricao, Preco, Categoria)VALUES 
(31, 'Lasanha à Bolonhesa', 'Deliciosa lasanha caseira com molho bolonhesa', 12.50, 'Almoço');

SELECT pr.nome, ip.idproduto, ip.idpedido
FROM itenspedidos ip
RIGHT JOIN produtos pr
ON ip.idproduto = pr.id;

SELECT pr.nome,  x.idproduto,  x.idpedido 
FROM(
    SELECT ip.idpedido, ip.idproduto
    FROM pedidos p
    JOIN itenspedidos ip 
    ON p.id = ip.idpedido
    WHERE strftime('%m', p.DataHoraPedido) = '10' ) x
RIGHT JOIN produtos pr
ON pr.id =  x.idproduto;


-- LEFT JOIN
INSERT INTO Clientes (id, Nome, Telefone, Email, Endereco)
VALUES (28, 'João Santos', '215555678', 'joao.santos@email.com', 'Avenida Principal, 456, Cidade B'),
       (29, 'Carla Ferreira', '315557890', 'carla.ferreira@email.com', 'Travessa das Ruas, 789, Cidade C');
       
SELECT c.nome, x.id
FROM clientes c 
LEFT JOIN
(
    SELECT p.id, p.idcliente 
    FROM pedidos p 
    WHERE strftime('%m', p.datahorapedido) = '10')x
ON c.id = x.idcliente
WHERE x.idcliente IS NULL;

-- O LEFT e RIGHT JOIN servem de maneira igual, apenas alterando o lado que é desejado da junção
-- A função dos dois é retornar todos os valores de uma tabela mesmo que ela não tenha ainda uma relação feita com a outra tabela