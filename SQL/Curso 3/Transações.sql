BEGIN TRANSACTION;

SELECT * FROM clientes;

UPDATE Pedidos SET status = 'Concluído' WHERE status = 'Em Andamento';

DELETE FROM clientes;

ROLLBACK;

COMMIT;

--A transação serve para ter uma camada a mais de segurança para atualização de dados, em que para realmente confirmar as alterações deve-se utilizar o commit ou caso não queira fazê-las utilize o rollback