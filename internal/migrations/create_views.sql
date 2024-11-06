
CREATE VIEW vendas_por_produto AS
SELECT prato.nome, COUNT(venda.id) AS total_vendas
FROM venda
JOIN prato ON venda.id_prato = prato.id
GROUP BY prato.nome;


CREATE VIEW clientes_com_pontos AS
SELECT nome, pontos
FROM cliente
WHERE pontos > 0;


CREATE VIEW ingredientes_proximos_vencimento AS
SELECT nome, data_validade
FROM ingrediente
WHERE data_validade < CURRENT_DATE + INTERVAL '7 days';