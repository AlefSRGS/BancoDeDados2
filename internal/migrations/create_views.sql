-- Views

-- View 1: Vendas por produto

CREATE VIEW vendas_por_produto AS
SELECT prato.nome, COUNT(venda.id) AS total_vendas
FROM venda
JOIN prato ON venda.id_prato = prato.id
GROUP BY prato.nome;

-- View 2: Clientes e seus pontos

CREATE VIEW clientes_com_pontos AS
SELECT nome, pontos
FROM cliente
WHERE pontos > 0;

-- View 3: Ingredientes próximos ao vencimento

CREATE VIEW ingredientes_proximos_vencimento AS
SELECT nome, data_validade
FROM ingredientes
WHERE data_validade < CURRENT_DATE + INTERVAL '7 days';