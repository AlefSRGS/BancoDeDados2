package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/vinicius457/BancoDeDados2/internal/model"
)

// Interface IngredienteRepoInterface define os métodos para o CRUD de Ingrediente
type IngredienteRepoInterface interface {
	Create(ctx context.Context, ingrediente model.Ingrediente) error
	Read(ctx context.Context, id int) (model.Ingrediente, error)
	Update(ctx context.Context, ingrediente model.Ingrediente) error
	Delete(ctx context.Context, id int) error
}

// Estrutura IngredienteRepo que implementa IngredienteRepoInterface
type IngredienteRepo struct {
	DB *sql.DB
}

// Construtor para IngredienteRepo, retorna a interface IngredienteRepoInterface
func NewIngredienteRepo(db *sql.DB) IngredienteRepoInterface {
	return &IngredienteRepo{DB: db}
}

// Create - insere um novo ingrediente no banco de dados
func (repo *IngredienteRepo) Create(ctx context.Context, ingrediente model.Ingrediente) error {
	query := "INSERT INTO ingrediente (nome, data_fabricacao, data_validade, quantidade, observacao) VALUES ($1, $2, $3, $4, $5)"
	_, err := repo.DB.ExecContext(ctx, query, ingrediente.Nome, ingrediente.DataFabricacao, ingrediente.DataValidade, ingrediente.Quantidade, ingrediente.Observacao)
	return err
}

// Read - busca um ingrediente por ID
func (repo *IngredienteRepo) Read(ctx context.Context, id int) (model.Ingrediente, error) {
	query := "SELECT id, nome, data_fabricacao, data_validade, quantidade, observacao FROM ingrediente WHERE id = $1"
	row := repo.DB.QueryRowContext(ctx, query, id)

	var ingrediente model.Ingrediente
	if err := row.Scan(&ingrediente.ID, &ingrediente.Nome, &ingrediente.DataFabricacao, &ingrediente.DataValidade, &ingrediente.Quantidade, &ingrediente.Observacao); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ingrediente, nil // Não encontrou, retorna ingrediente vazio
		}
		return ingrediente, err
	}
	return ingrediente, nil
}

// Update - atualiza um ingrediente por ID
func (repo *IngredienteRepo) Update(ctx context.Context, ingrediente model.Ingrediente) error {
	query := "UPDATE ingrediente SET nome = $1, data_fabricacao = $2, data_validade = $3, quantidade = $4, observacao = $5 WHERE id = $6"
	_, err := repo.DB.ExecContext(ctx, query, ingrediente.Nome, ingrediente.DataFabricacao, ingrediente.DataValidade, ingrediente.Quantidade, ingrediente.Observacao, ingrediente.ID)
	return err
}

// Delete - exclui um ingrediente por ID
func (repo *IngredienteRepo) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM ingrediente WHERE id = $1"
	_, err := repo.DB.ExecContext(ctx, query, id)
	return err
}
