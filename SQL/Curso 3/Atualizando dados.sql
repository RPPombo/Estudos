PRAGMA foreign_keys = ON;

INSERT INTO produtos (id, nome, descricao, preco, categoria)
VALUES (31, 'Lasanha à Bolonhesa', 'Deliciosa lasanha caseira com molho bolonhesa', 12.5, 'Almoço');

UPDATE produtos SET preco = 13
WHERE id = 31;

UPDATE produtos SET descricao = 'Croissant recheado com amêndoas'
WHERE id = 28;

--O SGBD SQLlite Online precisa que seja ativado a verificação de foreign keys
--Para atualizar dados deve-se sempre colocar o filtro de que linha o dado deve entrar, caso contrário todas as linhas da tabela serão atualizadas com aquele dado