package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/vinicius457/BancoDeDados2/internal/model"
)

// Interface ClienteRepoInterface define os métodos para o CRUD de Cliente
type ClienteRepoInterface interface {
	Create(ctx context.Context, cliente model.Cliente) error
	Read(ctx context.Context, id int) (model.Cliente, error)
	Update(ctx context.Context, cliente model.Cliente) error
	Delete(ctx context.Context, id int) error
}

// Estrutura ClienteRepo que implementa ClienteRepoInterface
type ClienteRepo struct {
	DB *sql.DB
}

// Construtor para ClienteRepo, retorna a interface ClienteRepoInterface
func NewClienteRepo(db *sql.DB) ClienteRepoInterface {
	return &ClienteRepo{DB: db}
}

// Implementação das funções CRUD para ClienteRepo

// Create - insere um novo cliente no banco de dados
func (repo *ClienteRepo) Create(ctx context.Context, cliente model.Cliente) error {
	query := "INSERT INTO cliente (nome, sexo, idade, nascimento, pontos) VALUES ($1, $2, $3, $4, $5)"
	_, err := repo.DB.ExecContext(ctx, query, cliente.Nome, cliente.Sexo, cliente.Idade, cliente.Nascimento, cliente.Pontos)
	return err
}

// Read - busca um cliente por ID
func (repo *ClienteRepo) Read(ctx context.Context, id int) (model.Cliente, error) {
	query := "SELECT id, nome, sexo, idade, nascimento, pontos FROM cliente WHERE id = $1"
	row := repo.DB.QueryRowContext(ctx, query, id)

	var cliente model.Cliente
	if err := row.Scan(&cliente.ID, &cliente.Nome, &cliente.Sexo, &cliente.Idade, &cliente.Nascimento, &cliente.Pontos); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return cliente, nil // Não encontrou, retorna cliente vazio
		}
		return cliente, err
	}
	return cliente, nil
}

// Update - atualiza um cliente por ID
func (repo *ClienteRepo) Update(ctx context.Context, cliente model.Cliente) error {
	query := "UPDATE cliente SET nome = $1, sexo = $2, idade = $3, nascimento = $4, pontos = $5 WHERE id = $6"
	_, err := repo.DB.ExecContext(ctx, query, cliente.Nome, cliente.Sexo, cliente.Idade, cliente.Nascimento, cliente.Pontos, cliente.ID)
	return err
}

// Delete - exclui um cliente por ID
func (repo *ClienteRepo) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM cliente WHERE id = $1"
	_, err := repo.DB.ExecContext(ctx, query, id)
	return err
}
