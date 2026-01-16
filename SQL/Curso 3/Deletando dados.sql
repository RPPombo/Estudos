DELETE FROM colaboradores
WHERE id =3;

DELETE FROM clientes
Where id = 27;

--A nossa tabela de clientes ao ser criada teve a opção de 'ON DELETE CASCADE', iss fez que todas as informações relecionadas a esse cliente fossem deletadas, independente de qual tabela ela estava
--Deve=se tomar cuidado quando deve-se ou não colocar essa opção em uma tabela
--Por outro lado, caso ao criar uma tabela com a opção 'ON DELETE RESTRICT', fará com que caso alguma informação de outra tabela relacionada a algum dos dados daquela tabela tente ser deletado, ocorrerá um erro