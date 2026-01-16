SELECT nome, rua, bairro, cidade, estado, cep FROM colaboradores
UNION
SELECT nome, rua, bairro, cidade, estado, cep FROM fornecedores;

-- Esse comando serve para unir duas tabelas
-- Esse comando sempre vai retornar as informações de maneira distinta, ou seja, sem repetições
-- Caso necessite da repetição de valores utilize UNION ALL
-- Para a utilização desse comando as tabelas devem ser compatíveis, ou seja, seus campos devem ser iguais, tanto o nome quanto a sua descrição
-- EXCEPT, serve para retornar os valores que só estão no conjunto A e que não estão no conjunto B
-- INTERSECT, serve para retornar os valores presentes tanto no conjunto A quanto no conjunto B