-- Usuários e permissões

CREATE USER admin WITH PASSWORD 'admin_admin'; -- Cria um "admin" com a senha "admin_admin"
GRANT ALL PRIVILEGES ON DATABASE restaurante TO admin; -- Concede ao usuário "admin" todos os privilégios no banco de dados "restaurante"
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO admin; -- Concede ao usuário "admin" todos os privilégios em todas as tabelas no esquema "public"
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO admin; -- Concede ao usuário "admin" todos os privilégios em todas as sequências no esquema "public"

CREATE USER gerente WITH PASSWORD 'senha_gerente'; -- Cria um "gerente" com a senha "senha_gerente"
GRANT SELECT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO gerente; -- Concede ao usuário "gerente" permissões de SELECT, UPDATE e DELETE em todas as tabelas no esquema "public"
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO gerente; -- Concede ao usuário "gerente" permissões de USAGE e SELECT em todas as sequências no esquema "public"

CREATE USER funcionario WITH PASSWORD 'senha_funcionario'; -- Cria um "funcionario" com a senha "senha_funcionario"
GRANT INSERT, SELECT ON ALL TABLES IN SCHEMA public TO funcionario; -- Concede ao usuário "funcionario" permissões de INSERT e SELECT em todas as tabelas no esquema "public"
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO funcionario; -- Concede ao usuário "funcionario" permissões de USAGE e SELECT em todas as sequências no esquema "public"