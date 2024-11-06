-- Funções

-- Função para cálculo de pontos

CREATE OR REPLACE FUNCTION calcula_pontos(valor NUMERIC) RETURNS INT AS $$
BEGIN
    RETURN FLOOR(valor / 10);
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION atualiza_quantidade_ingrediente() RETURNS TRIGGER AS $$
BEGIN
    UPDATE ingrediente
    SET quantidade = quantidade - NEW.quantidade
    WHERE id IN (SELECT id_ingrediente FROM usos WHERE id_prato = NEW.id_prato);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION impedir_compra_prato_indisponivel() RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT disponibilidade FROM prato WHERE id = NEW.id_prato) = FALSE THEN
        RAISE EXCEPTION 'Prato indisponível para compra';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

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

CREATE OR REPLACE FUNCTION calcula_pontos() RETURNS TRIGGER AS $$
BEGIN
    UPDATE cliente
    SET pontos = pontos + FLOOR(NEW.valor / 10)
    WHERE id = NEW.id_cliente;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;