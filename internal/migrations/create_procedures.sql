-- Procedimentos

-- Procedimento 1: Reajustar valores de pratos

CREATE OR REPLACE PROCEDURE reajuste(percentual NUMERIC)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE prato
    SET valor = valor + (valor * percentual / 100);
END;
$$;

-- Procedimento 2: Sortear cliente para premiação

CREATE OR REPLACE PROCEDURE sorteio_cliente()
LANGUAGE plpgsql
AS $$
DECLARE
    cliente_id INT;
BEGIN
    SELECT id INTO cliente_id
    FROM cliente
    ORDER BY RANDOM()
    LIMIT 1;

    UPDATE cliente
    SET pontos = pontos + 100
    WHERE id = cliente_id;
END;
$$;

-- Procedimento 3: Exibir estatísticas de vendas

CREATE OR REPLACE PROCEDURE estatisticas_vendas()
LANGUAGE plpgsql
AS $$
BEGIN
    -- Produto mais vendido
    SELECT prato.nome AS produto_mais_vendido, COUNT(venda.id_prato) AS total_vendas
    FROM venda
    JOIN prato ON venda.id_prato = prato.id
    GROUP BY prato.nome
    ORDER BY total_vendas DESC
    LIMIT 1;

    -- Produto menos vendido
    SELECT prato.nome AS produto_menos_vendido, COUNT(venda.id_prato) AS total_vendas
    FROM venda
    JOIN prato ON venda.id_prato = prato.id
    GROUP BY prato.nome
    ORDER BY total_vendas ASC
    LIMIT 1;
END;
$$;