# BandoDeDados2

1. Conectar com o Banco de Dados (feito para testes, mas terá que ser feito de outro jeito, pois não dá pra levar o pc)

2. Criação e Destruição do Banco de Dados(podemos utilizar o docker para levantar e derrubar)

3. Criação das Tabelas (feito, falta apenas testar)

4. Inserção de Dados Iniciais*

5. Settar os Triggers(falta testar, mas já implementados)

6. Criar Users & Roles(users feitos)

7. Criar Views & Procedimentos(um dos views feito)

8. Criar Funções(funções feitas)

9. Back-end em Go ---CRUD e linkar com as funções 


# Estrutura do projeto:

### `cmd/`
Contém o ponto de entrada da aplicação. O arquivo `main.go` é responsável por inicializar o servidor e configurar as rotas.

### `internal/`
Contém a lógica interna do aplicativo, incluindo a definição de modelos, interação com o banco de dados, regras de negócios e manipulação de requisições HTTP.

- **`model/`**: Define os modelos de dados que representam as entidades do sistema, como usuários.
    - Ex: User: Campos como ID, Name, Email, Password.
    - NewUser(): Construtor para criar uma nova instância de User.
- **`schemas/`**: Define esquemas e validações para garantir que os dados estejam corretos antes de serem processados.
- **`repository/`**: Responsável pela interação com a base de dados, implementando as operações CRUD.
- **`service/`**: Contém a lógica de negócios, incluindo validações e regras específicas para a manipulação de dados.
- **`handler/`**: Manipula as requisições HTTP e encaminha os dados para os serviços apropriados.
- **`routes/`**: Define as rotas da API, ligando os endpoints aos manipuladores.

### `pkg/`
Contém pacotes que podem ser reutilizados em diferentes partes do projeto ou em outros projetos.

- **`utils/`**: Funções utilitárias que podem ser utilizadas em todo o projeto para evitar duplicação de código.
    - Ex: HashPassword(): para hashar senhas antes de armazenar no banco.

### `migrations/`
Contém scripts SQL para migrações de banco de dados, permitindo a criação e modificação de tabelas conforme necessário.

### `test/`
Contém testes unitários para verificar a funcionalidade dos componentes do sistema.