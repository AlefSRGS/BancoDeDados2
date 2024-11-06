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