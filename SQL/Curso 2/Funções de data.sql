SELECT id_colaborador, STRFTIME('%Y/%m', datainicio) FROM Licencas;

SELECT id_colaborador, JULIANDAY(datatermino) - JULIANDAY(datacontratacao) FROM HistoricoEmprego
WHERE datatermino IS NOT NULL;

--STRFTIME() altera o formato original da data para ser mostrado na consulta
--JULIANDAY() serve para fazer o calculo entre duas colunas de uma tabela
--Outras funções de data:
--DATE('now', '[modificador]'), serve para ver a data (retorna a data no formato 'YYYY-MM-DD')
--TIME('now', '[modificador]'), serve para ver a hora (retorna a hora no formato 'HH:MM:SS')
--DATETIME('now', '[modificador]'), serve para ver a data e hora (retorna no formato'YYYY-MM-DD HH:MM:SS')
--CURRENT_TIMESTAMP, é uma função de comodidade, funciona do mesmo jeito que colocar DATETIME('now')