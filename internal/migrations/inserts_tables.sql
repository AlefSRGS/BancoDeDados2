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