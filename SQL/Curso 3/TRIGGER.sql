CREATE TABLE FaturamentoDiario (
    Dia DATE,
    FaturamentoTotal DECIMAL(10, 2)
);

CREATE TRIGGER CalculaFaturamentoDiario
AFTER INSERT ON itenspedidos
FOR EACH ROW
BEGIN
DELETE FROM FaturamentoDiario;
INSERT INTO FaturamentoDiario (Dia, FaturamentoTotal)
SELECT DATE(datahorapedido) AS Dia, SUM(ip.precounitario) AS FaturamentoDiario
FROM pedidos p
JOIN itenspedidos ip
ON p.id = ip.idpedido
GROUP BY Dia
ORDER BY Dia;
END

INSERT INTO Pedidos(ID, IDCliente, DataHoraPedido, Status)
VALUES (451, 27, '2023-10-07 14:30:00', 'Em Andamento');

INSERT INTO ItensPedidos(IDPedido, IDProduto, Quantidade, PrecoUnitario)
VALUES (451, 14, 1, 6.0),
       (451, 13, 1, 7.0);
       
INSERT INTO Pedidos (ID, IDCliente, DataHoraPedido, Status) 
VALUES (452, 28, '2023-10-07 14:35:00', 'Em Andamento');

INSERT INTO ItensPedidos (IDPedido, IDProduto, Quantidade, PrecoUnitario) 
VALUES (452, 10, 1, 5.0),
       (452, 31, 1, 12.50);

SELECT * FROM FaturamentoDiario;
-- Um TRIGGER serve para criar uma automatização de um código, em que sempre que uma ação ocorrer um bloco de código vai rodar