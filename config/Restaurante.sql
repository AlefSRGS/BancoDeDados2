CREATE OR REPLACE FUNCTION calcula_pontos() RETURNS TRIGGER AS $$
BEGIN
    UPDATE cliente
    SET pontos = pontos + FLOOR(NEW.valor / 10)
    WHERE id = NEW.id_cliente;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_calcula_pontos
AFTER INSERT ON venda
FOR EACH ROW
EXECUTE FUNCTION calcula_pontos();

CREATE OR REPLACE FUNCTION atualiza_disponibilidade_prato() RETURNS TRIGGER AS $$
BEGIN
    IF NEW.data_validade < CURRENT_DATE THEN
        UPDATE prato
        SET disponibilidade = FALSE
        WHERE id IN (SELECT id_prato FROM usos WHERE id_ingrediente = NEW.id);
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_atualiza_disponibilidade_prato
AFTER UPDATE ON ingredientes
FOR EACH ROW
EXECUTE FUNCTION atualiza_disponibilidade_prato();

CREATE OR REPLACE FUNCTION atualiza_quantidade_ingrediente() RETURNS TRIGGER AS $$
BEGIN
    UPDATE ingredientes
    SET quantidade = quantidade - NEW.quantidade
    WHERE id IN (SELECT id_ingrediente FROM usos WHERE id_prato = NEW.id_prato);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_atualiza_quantidade_ingrediente
AFTER INSERT ON venda
FOR EACH ROW
EXECUTE FUNCTION atualiza_quantidade_ingrediente();

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

CREATE OR REPLACE FUNCTION calcula_pontos(valor NUMERIC) RETURNS INT AS $$
BEGIN
    RETURN FLOOR(valor / 10);
END;
$$ LANGUAGE plpgsql;

CREATE USER admin WITH PASSWORD 'admin_admin';
GRANT ALL PRIVILEGES ON DATABASE restaurante TO admin;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO admin;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO admin;

CREATE USER gerente WITH PASSWORD 'senha_gerente';
GRANT SELECT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO gerente;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO gerente;

CREATE USER funcionario WITH PASSWORD 'senha_funcionario';
GRANT INSERT, SELECT ON ALL TABLES IN SCHEMA public TO funcionario;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO funcionario;

CREATE VIEW vendas_por_produto AS
SELECT prato.nome, COUNT(venda.id) AS total_vendas
FROM venda
JOIN prato ON venda.id_prato = prato.id
GROUP BY prato.nome;
