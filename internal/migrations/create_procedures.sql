CREATE OR REPLACE PROCEDURE reajuste(percentual NUMERIC)
LANGUAGE plpgsql
AS $$
BEGIN
    UPDATE prato
    SET valor = valor + (valor * percentual / 100);
END;
$$;

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

CREATE OR REPLACE PROCEDURE estatisticas_vendas()
LANGUAGE plpgsql
AS $$
BEGIN

    SELECT prato.nome AS produto_mais_vendido, COUNT(venda.id_prato) AS total_vendas
    FROM venda
    JOIN prato ON venda.id_prato = prato.id
    GROUP BY prato.nome
    ORDER BY total_vendas DESC
    LIMIT 1;

    
    SELECT prato.nome AS produto_menos_vendido, COUNT(venda.id_prato) AS total_vendas
    FROM venda
    JOIN prato ON venda.id_prato = prato.id
    GROUP BY prato.nome
    ORDER BY total_vendas ASC
    LIMIT 1;
END;
$$;
