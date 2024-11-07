package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/vinicius457/BancoDeDados2/internal/model"
)

// Interface PratoRepoInterface define os métodos para o CRUD de Prato
type PratoRepoInterface interface {
	Create(ctx context.Context, prato model.Prato) error
	Read(ctx context.Context, id int) (model.Prato, error)
	Update(ctx context.Context, prato model.Prato) error
	Delete(ctx context.Context, id int) error
}

// Estrutura PratoRepo que implementa PratoRepoInterface
type PratoRepo struct {
	DB *sql.DB
}

// Construtor para PratoRepo, retorna a interface PratoRepoInterface
func NewPratoRepo(db *sql.DB) PratoRepoInterface {
	return &PratoRepo{DB: db}
}

// Create - insere um novo prato no banco de dados
func (repo *PratoRepo) Create(ctx context.Context, prato model.Prato) error {
	query := "INSERT INTO prato (nome, descricao, valor, disponibilidade) VALUES ($1, $2, $3, $4)"
	_, err := repo.DB.ExecContext(ctx, query, prato.Nome, prato.Descricao, prato.Valor, prato.Disponibilidade)
	return err
}

// Read - busca um prato por ID
func (repo *PratoRepo) Read(ctx context.Context, id int) (model.Prato, error) {
	query := "SELECT id, nome, descricao, valor, disponibilidade FROM prato WHERE id = $1"
	row := repo.DB.QueryRowContext(ctx, query, id)

	var prato model.Prato
	if err := row.Scan(&prato.ID, &prato.Nome, &prato.Descricao, &prato.Valor, &prato.Disponibilidade); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return prato, nil // Não encontrou, retorna prato vazio
		}
		return prato, err
	}
	return prato, nil
}

// Update - atualiza um prato por ID
func (repo *PratoRepo) Update(ctx context.Context, prato model.Prato) error {
	query := "UPDATE prato SET nome = $1, descricao = $2, valor = $3, disponibilidade = $4 WHERE id = $5"
	_, err := repo.DB.ExecContext(ctx, query, prato.Nome, prato.Descricao, prato.Valor, prato.Disponibilidade, prato.ID)
	return err
}

// Delete - exclui um prato por ID
func (repo *PratoRepo) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM prato WHERE id = $1"
	_, err := repo.DB.ExecContext(ctx, query, id)
	return err
}
