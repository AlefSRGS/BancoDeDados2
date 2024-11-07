package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/vinicius457/BancoDeDados2/internal/model"
)

// Interface FornecedorRepoInterface define os métodos para o CRUD de Fornecedor
type FornecedorRepoInterface interface {
	Create(ctx context.Context, fornecedor model.Fornecedor) error
	Read(ctx context.Context, id int) (model.Fornecedor, error)
	Update(ctx context.Context, fornecedor model.Fornecedor) error
	Delete(ctx context.Context, id int) error
}

// Estrutura FornecedorRepo que implementa FornecedorRepoInterface
type FornecedorRepo struct {
	DB *sql.DB
}

// Construtor para FornecedorRepo, retorna a interface FornecedorRepoInterface
func NewFornecedorRepo(db *sql.DB) FornecedorRepoInterface {
	return &FornecedorRepo{DB: db}
}

// Create - insere um novo fornecedor no banco de dados
func (repo *FornecedorRepo) Create(ctx context.Context, fornecedor model.Fornecedor) error {
	query := "INSERT INTO fornecedor (nome, estado_origem) VALUES ($1, $2)"
	_, err := repo.DB.ExecContext(ctx, query, fornecedor.Nome, fornecedor.EstadoOrigem)
	return err
}

// Read - busca um fornecedor por ID
func (repo *FornecedorRepo) Read(ctx context.Context, id int) (model.Fornecedor, error) {
	query := "SELECT id, nome, estado_origem FROM fornecedor WHERE id = $1"
	row := repo.DB.QueryRowContext(ctx, query, id)

	var fornecedor model.Fornecedor
	if err := row.Scan(&fornecedor.ID, &fornecedor.Nome, &fornecedor.EstadoOrigem); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fornecedor, nil // Não encontrou, retorna fornecedor vazio
		}
		return fornecedor, err
	}
	return fornecedor, nil
}

// Update - atualiza um fornecedor por ID
func (repo *FornecedorRepo) Update(ctx context.Context, fornecedor model.Fornecedor) error {
	query := "UPDATE fornecedor SET nome = $1, estado_origem = $2 WHERE id = $3"
	_, err := repo.DB.ExecContext(ctx, query, fornecedor.Nome, fornecedor.EstadoOrigem, fornecedor.ID)
	return err
}

// Delete - exclui um fornecedor por ID
func (repo *FornecedorRepo) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM fornecedor WHERE id = $1"
	_, err := repo.DB.ExecContext(ctx, query, id)
	return err
}
