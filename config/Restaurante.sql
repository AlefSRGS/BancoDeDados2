CREATE TABLE IF NOT EXISTS cliente ( -- Tabela cliente
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100),
    sexo CHAR(1) CHECK (sexo IN ('m', 'f', 'o')),
    idade INT,
    nascimento DATE,
    pontos INT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS prato ( -- Tabela prato
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100),
    descricao TEXT,
    valor NUMERIC(10, 2),
    disponibilidade BOOLEAN DEFAULT TRUE
);

CREATE TABLE IF NOT EXISTS fornecedor ( -- Tabela fornecedor
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100),
    estado_origem CHAR(2) CHECK (estado_origem IN ('AC', 'AL', 'AP', 'AM', 'BA', 'CE', 'DF', 'ES', 'GO', 'MA', 'MT', 'MS', 'MG', 'PA', 'PB', 'PR', 'PE', 'PI', 'RJ', 'RN', 'RS', 'RO', 'RR', 'SC', 'SP', 'SE', 'TO'))
);

CREATE TABLE IF NOT EXISTS ingrediente ( -- Tabela ingredientes
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100),
    data_fabricacao DATE,
    data_validade DATE,
    quantidade INT,
    observacao TEXT
);

CREATE TABLE IF NOT EXISTS usos ( -- Tabela usos (relaciona prato com seus ingrediente)
    id_prato INT,
    id_ingrediente INT,
    PRIMARY KEY (id_prato, id_ingrediente),
    FOREIGN KEY (id_prato) REFERENCES prato(id),
    FOREIGN KEY (id_ingrediente) REFERENCES ingrediente(id)
);

CREATE TABLE IF NOT EXISTS venda ( -- Tabela venda
    id SERIAL PRIMARY KEY,
    id_cliente INT REFERENCES cliente(id),
    id_prato INT REFERENCES prato(id),
    quantidade INT,
    dia DATE,
    hora TIME,
    valor NUMERIC(10, 2)
);

INSERT INTO cliente (nome, sexo, idade, nascimento) VALUES -- Inserções de Clientes (10 registros)
('João', 'm', 30, '1994-02-15'),
('Maria', 'f', 25, '1999-10-22'),
('Carlos', 'm', 35, '1989-06-10'),
('Ana', 'f', 40, '1984-03-05'),
('Pedro', 'm', 28, '1996-09-12'),
('Beatriz', 'f', 22, '2002-01-18'),
('Ricardo', 'm', 45, '1978-11-30'),
('Fernanda', 'f', 32, '1992-04-21'),
('Lucas', 'm', 29, '1995-12-08'),
('Sofia', 'f', 27, '1997-07-15');

INSERT INTO prato (nome, descricao, valor) VALUES -- Inserções de Pratos (10 registros)
('Pizza Margherita', 'Queijo, tomate e manjericão', 35.00),
('Lasanha', 'Massa, queijo e molho de carne', 45.00),
('Salada Caesar', 'Alface, croutons, frango e molho Caesar', 25.00),
('Hambúrguer', 'Pão, carne, queijo e salada', 20.00),
('Sopa de Legumes', 'Sopa com diversos legumes frescos', 18.00),
('Frango Assado', 'Frango temperado e assado', 30.00),
('Espaguete', 'Espaguete com molho de tomate', 28.00),
('Bife à Parmegiana', 'Bife, queijo e molho de tomate', 40.00),
('Risoto de Camarão', 'Arroz com camarão e temperos', 55.00),
('Torta de Limão', 'Sobremesa doce de limão', 15.00);

INSERT INTO fornecedor (nome, estado_origem) VALUES -- Inserções de Fornecedores (10 registros)
('Fornecedor A', 'SP'),
('Fornecedor B', 'RJ'),
('Fornecedor C', 'MG'),
('Fornecedor D', 'BA'),
('Fornecedor E', 'PR'),
('Fornecedor F', 'SC'),
('Fornecedor G', 'PE'),
('Fornecedor H', 'RS'),
('Fornecedor I', 'GO'),
('Fornecedor J', 'ES');

INSERT INTO ingrediente (nome, data_fabricacao, data_validade, quantidade, observacao) VALUES -- Inserções de Ingredientes (10 registros)
('Tomate', '2024-01-01', '2024-03-01', 100, 'Fresco'),
('Alface', '2024-02-01', '2024-04-01', 50, 'Orgânico'),
('Queijo', '2024-01-15', '2024-03-15', 200, 'Mussarela'),
('Frango', '2024-01-20', '2024-02-20', 150, 'Congelado'),
('Carne', '2024-02-05', '2024-04-05', 180, 'Fresca'),
('Manjericão', '2024-02-10', '2024-03-10', 70, 'Fresco'),
('Camarão', '2024-02-15', '2024-03-15', 90, 'Descongelado'),
('Arroz', '2024-01-05', '2024-05-05', 300, 'Integral'),
('Limão', '2024-02-01', '2024-04-01', 120, 'Taiti'),
('Farinha', '2024-01-10', '2024-05-10', 250, 'De trigo');

INSERT INTO venda (id_cliente, id_prato, quantidade, dia, hora, valor) VALUES -- Inserção inicial na tabela venda (A SE DISCUTIR)
(1, 1, 2, '2024-03-01', '12:30:00', 70.00),
(2, 2, 1, '2024-03-02', '13:15:00', 45.00),
(3, 3, 3, '2024-03-03', '14:00:00', 75.00),
(4, 4, 1, '2024-03-04', '15:45:00', 20.00),
(5, 5, 2, '2024-03-05', '16:20:00', 36.00),
(6, 6, 1, '2024-03-06', '17:35:00', 30.00),
(7, 7, 2, '2024-03-07', '18:50:00', 80.00),
(8, 8, 3, '2024-03-08', '19:10:00', 120.00),
(9, 9, 2, '2024-03-09', '20:25:00', 56.00),
(10, 10, 1, '2024-03-10', '21:40:00', 15.00);

-- Triggers

-- Trigger 1: Calcula Pontos

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

-- Trigger 2: Atualizar disponibilidade de pratos com ingredientes vencidos

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
AFTER UPDATE ON ingrediente
FOR EACH ROW
EXECUTE FUNCTION atualiza_disponibilidade_prato();

-- Trigger 3: Impedir compra de prato indisponível

CREATE OR REPLACE FUNCTION impedir_compra_prato_indisponivel() RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT disponibilidade FROM prato WHERE id = NEW.id_prato) = FALSE THEN
        RAISE EXCEPTION 'Prato indisponível para compra';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_impedir_compra_prato_indisponivel
BEFORE INSERT ON venda
FOR EACH ROW
EXECUTE FUNCTION impedir_compra_prato_indisponivel();

-- Trigger 4: Atualizar a quantidade de ingredientes de venda

CREATE OR REPLACE FUNCTION atualiza_quantidade_ingrediente() RETURNS TRIGGER AS $$
BEGIN
    UPDATE ingrediente
    SET quantidade = quantidade - NEW.quantidade
    WHERE id IN (SELECT id_ingrediente FROM usos WHERE id_prato = NEW.id_prato);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_atualiza_quantidade_ingrediente
AFTER INSERT ON venda
FOR EACH ROW
EXECUTE FUNCTION atualiza_quantidade_ingrediente();

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

-- Funções

-- Função para cálculo de pontos

CREATE OR REPLACE FUNCTION calcula_pontos(valor NUMERIC) RETURNS INT AS $$
BEGIN
    RETURN FLOOR(valor / 10);
END;
$$ LANGUAGE plpgsql;

-- Usuários e permissões

CREATE USER admin WITH PASSWORD 'admin_admin'; -- Cria um "admin" com a senha "admin_admin"
GRANT ALL PRIVILEGES ON DATABASE restaurante_db TO admin; -- Concede ao usuário "admin" todos os privilégios no banco de dados "restaurante_db"
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO admin; -- Concede ao usuário "admin" todos os privilégios em todas as tabelas no esquema "public"
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO admin; -- Concede ao usuário "admin" todos os privilégios em todas as sequências no esquema "public"

CREATE USER gerente WITH PASSWORD 'senha_gerente'; -- Cria um "gerente" com a senha "senha_gerente"
GRANT SELECT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO gerente; -- Concede ao usuário "gerente" permissões de SELECT, UPDATE e DELETE em todas as tabelas no esquema "public"
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO gerente; -- Concede ao usuário "gerente" permissões de USAGE e SELECT em todas as sequências no esquema "public"

CREATE USER funcionario WITH PASSWORD 'senha_funcionario'; -- Cria um "funcionario" com a senha "senha_funcionario"
GRANT INSERT, SELECT ON ALL TABLES IN SCHEMA public TO funcionario; -- Concede ao usuário "funcionario" permissões de INSERT e SELECT em todas as tabelas no esquema "public"
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO funcionario; -- Concede ao usuário "funcionario" permissões de USAGE e SELECT em todas as sequências no esquema "public"

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
FROM ingrediente
WHERE data_validade < CURRENT_DATE + INTERVAL '7 days';